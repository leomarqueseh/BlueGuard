package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/httpclient"
	"github.com/leomarqueseh/BlueGuard/internal/plugins"
	"github.com/leomarqueseh/BlueGuard/internal/recon"
	"github.com/leomarqueseh/BlueGuard/internal/risk"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
	"github.com/leomarqueseh/BlueGuard/internal/worker"
)

//
// 🔥 Armazena resultados do scan (memória)
//
var results []scanner.Finding

//
// 🚀 Inicializa servidor web
//
func Start() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/scan", scanHandler)
	http.HandleFunc("/results", resultsHandler)
	http.HandleFunc("/recon", reconHandler)
	http.HandleFunc("/expand", expandHandler)

	fmt.Println("[+] Dashboard running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

//
// 🏠 HOME PAGE (Interface principal)
//
func homeHandler(w http.ResponseWriter, r *http.Request) {

	html := `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>BlueGuard</title>

<style>
body {
	margin:0;
	height:100vh;
	display:flex;
	align-items:center;
	justify-content:center;
	background:#000;
	font-family:Arial;
	color:#fff;
}

.container {
	text-align:center;
	width:100%;
	max-width:700px;
}

.logo {
	font-size:42px;
	font-weight:bold;
	margin-bottom:10px;
}

.icon {
	font-size:40px;
	margin-bottom:10px;
}

.subtitle {
	color:#94a3b8;
	margin-bottom:40px;
}

.input-group {
	display:flex;
	background:#111;
	border-radius:50px;
	padding:5px;
	box-shadow:0 0 20px rgba(255,255,255,0.1);
}

.input-group input {
	flex:1;
	padding:18px;
	border:none;
	outline:none;
	background:transparent;
	color:#fff;
	font-size:16px;
}

.input-group button {
	background:#000;
	border:1px solid #333;
	padding:0 25px;
	border-radius:50px;
	color:#fff;
	cursor:pointer;
}

/* plugins */
.plugins {
	margin-top:30px;
	text-align:left;
	background:#000;
	padding:20px;
	border-radius:12px;
	border:1px solid #1e293b;
}

.plugin-item {
	display:flex;
	gap:10px;
	margin-bottom:10px;
}

.footer {
	margin-top:40px;
	color:#64748b;
}
</style>
</head>

<body>

<div class="container">

<div class="icon">🎩👓</div>
<div class="logo">BlueGuard</div>
<div class="subtitle">Offensive Security Platform</div>

<form action="/scan">

<div class="input-group">
	<input name="target" placeholder="https://example.com" required>
	<button>Scan</button>
</div>

<div class="plugins">
<h3>Plugins</h3>

<label class="plugin-item"><input type="checkbox" name="plugins" value="open_redirect" checked>Open Redirect</label>
<label class="plugin-item"><input type="checkbox" name="plugins" value="header_exposure" checked>Header Exposure</label>
<label class="plugin-item"><input type="checkbox" name="plugins" value="tech_fingerprint" checked>Tech Fingerprint</label>
<label class="plugin-item"><input type="checkbox" name="plugins" value="git_exposed">Git Exposed</label>
<label class="plugin-item"><input type="checkbox" name="plugins" value="git_dump">Git Dump</label>

</div>

</form>

<div class="footer">stealth • recon • exploit</div>

</div>
</body>
</html>
`

	w.Write([]byte(html))
}

//
// 🔍 Executa scan
//
func scanHandler(w http.ResponseWriter, r *http.Request) {

	target := r.URL.Query().Get("target")

	ctx := &core.ScanContext{
		Timeout:   10 * time.Second,
		UserAgent: "BlueGuard",
		Client:    httpclient.New(10 * time.Second),
	}

	reg := plugins.NewRegistry()
	pool := worker.NewPool(reg.All(), 5, ctx)

	findings := pool.Run(r.Context(), []scanner.Target{
		{URL: target},
	})

	// 🔥 aplica análise de risco
	findings = risk.New().Analyze(findings)

	results = findings

	http.Redirect(w, r, "/results?target="+target, http.StatusSeeOther)
}

//
// 📊 RESULTADOS + MAPA
//
func resultsHandler(w http.ResponseWriter, r *http.Request) {

	target := r.URL.Query().Get("target")

	var rows string

	// 🔥 mesmo sem vulnerabilidade, mostra relatório
	if len(results) == 0 {
		rows = `<tr><td colspan="4">No vulnerabilities found. Target appears secure.</td></tr>`
	} else {
		for _, f := range results {
			rows += fmt.Sprintf(`
<tr>
<td>%s</td>
<td>%s</td>
<td>%.1f</td>
<td>%s</td>
</tr>
`, f.Title, f.Severity, f.Score, f.Target)
		}
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<script src="https://unpkg.com/cytoscape/dist/cytoscape.min.js"></script>

<style>
body { background:#000; color:#fff; font-family:Arial; }
.container { padding:20px; }
#cy { height:600px; }
</style>

</head>
<body>

<div class="container">

<h2>Results</h2>

<table>
<tr><th>Title</th><th>Severity</th><th>Score</th><th>Target</th></tr>
%s
</table>

<h3>Recon Map</h3>
<div id="cy"></div>

</div>

<script>
async function loadGraph(){
	const res = await fetch("/recon?target=%s");
	const data = await res.json();

	let cy = cytoscape({
		container: document.getElementById('cy'),

		elements: [
			...data.nodes.map(n => ({
				data: { id: n.id, label: n.label, type: n.type }
			})),
			...data.edges.map(e => ({
				data: { source: e.source, target: e.target }
			}))
		],

		layout: { name: 'cose' }
	});

	// 🔥 expansão inteligente
	cy.on('tap', 'node', async function(evt){
		const node = evt.target;

		const res = await fetch("/expand?node=" + node.id() + "&type=" + node.data("type"));
		const data = await res.json();

		data.nodes.forEach(n => {
			if(!cy.getElementById(n.id).length){
				cy.add({ data: n });
			}
		});

		data.edges.forEach(e => cy.add({ data: e }));

		cy.layout({ name: 'cose' }).run();
	});
}

loadGraph();
</script>

</body>
</html>
`, rows, target)

	w.Write([]byte(html))
}

//
// 🌐 RECON (corrigido com providers)
//
func reconHandler(w http.ResponseWriter, r *http.Request) {

	target := r.URL.Query().Get("target")

	// 🔥 providers vazios (evita erro)
	var providers []recon.Provider

	asset := recon.Discover(target, providers)

	graph := recon.BuildGraph(asset)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graph)
}

//
// 🔥 EXPANSÃO DINÂMICA DO GRAFO
//
func expandHandler(w http.ResponseWriter, r *http.Request) {

	node := r.URL.Query().Get("node")
	typ := r.URL.Query().Get("type")

	var nodes []recon.Node
	var edges []recon.Edge

	switch typ {

	case "domain", "subdomain":

		ips := recon.ResolveIP(node)

		for _, ip := range ips {

			id := node + "_ip_" + ip

			nodes = append(nodes, recon.Node{
				ID:    id,
				Label: ip,
				Type:  "ip",
			})

			edges = append(edges, recon.Edge{
				Source: node,
				Target: id,
			})
		}

	case "ip":

		ports := recon.ScanPorts(node)

		for _, p := range ports {

			id := node + "_port_" + p.Port

			nodes = append(nodes, recon.Node{
				ID:    id,
				Label: p.Port + "/" + p.Service,
				Type:  "port",
			})

			edges = append(edges, recon.Edge{
				Source: node,
				Target: id,
			})
		}
	}

	json.NewEncoder(w).Encode(recon.Graph{
		Nodes: nodes,
		Edges: edges,
	})
}
