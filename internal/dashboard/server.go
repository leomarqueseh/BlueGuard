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

// 🔥 armazenamento temporário
var results []scanner.Finding

// 🚀 servidor
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
// 🏠 HOME (INALTERADO)
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
	background: radial-gradient(circle at top,#0f172a,#020617);
	font-family:Arial;
	color:#fff;
}

.container {
	text-align:center;
	width:100%;
	max-width:700px;
}

.logo {
	font-size:48px;
	font-weight:bold;
	margin-bottom:10px;
}

.subtitle {
	color:#94a3b8;
	margin-bottom:40px;
}

/* INPUT */
.input-group {
	display:flex;
	background:#111827;
	border-radius:50px;
	padding:5px;
	box-shadow:0 0 20px rgba(34,197,94,0.2);
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
	background:#22c55e;
	border:none;
	padding:0 25px;
	border-radius:50px;
	color:#fff;
	font-weight:bold;
	cursor:pointer;
}

/* 🔥 PLUGINS */
.plugins {
	margin-top:30px;
	text-align:left;
	background:#020617;
	padding:20px;
	border-radius:12px;
	border:1px solid #1e293b;
}

.plugins h3 {
	margin-bottom:15px;
	color:#94a3b8;
}

.plugin-item {
	display:flex;
	align-items:center;
	gap:10px;
	margin-bottom:12px;
	cursor:pointer;
}

.plugin-item input {
	width:18px;
	height:18px;
	cursor:pointer;
}

.plugin-item span {
	color:#cbd5f5;
}

/* FOOTER */
.footer {
	margin-top:40px;
	color:#64748b;
}
</style>
</head>

<body>

<div class="container">

<div class="logo">🛡️ BlueGuard</div>
<div class="subtitle">Next-gen vulnerability scanner</div>

<form action="/scan">

<div class="input-group">
	<input name="target" placeholder="https://example.com" required>
	<button>Scan</button>
</div>

<!-- 🔥 PLUGINS (ADICIONADO SEM QUEBRAR DESIGN) -->
<div class="plugins">

<h3>Plugins</h3>

<label class="plugin-item">
	<input type="checkbox" name="plugins" value="open_redirect" checked>
	<span>Open Redirect</span>
</label>

<label class="plugin-item">
	<input type="checkbox" name="plugins" value="header_exposure" checked>
	<span>Header Exposure</span>
</label>

<label class="plugin-item">
	<input type="checkbox" name="plugins" value="tech_fingerprint" checked>
	<span>Tech Fingerprint</span>
</label>

<label class="plugin-item">
	<input type="checkbox" name="plugins" value="git_exposed">
	<span>Git Exposed</span>
</label>

<label class="plugin-item">
	<input type="checkbox" name="plugins" value="git_dump">
	<span>Git Dump</span>
</label>

</div>

</form>

<div class="footer">
Secure • Fast • Modular
</div>

</div>

</body>
</html>
`

	w.Write([]byte(html))
}

//
// 🔍 SCAN
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

	targets := []scanner.Target{
		{URL: target},
	}

	findings := pool.Run(r.Context(), targets)
	findings = risk.New().Analyze(findings)

	results = findings

	http.Redirect(w, r, "/results?target="+target, http.StatusSeeOther)
}

//
// 📊 RESULTS + MAP
//
func resultsHandler(w http.ResponseWriter, r *http.Request) {

	target := r.URL.Query().Get("target")

	var high, medium, low, info int

	for _, f := range results {
		switch f.Severity {
		case "HIGH":
			high++
		case "MEDIUM":
			medium++
		case "LOW":
			low++
		default:
			info++
		}
	}

	var rows string

	for _, f := range results {
		rows += fmt.Sprintf(`
<tr>
<td>%s</td>
<td class="sev %s">%s</td>
<td>%.1f</td>
<td>%s</td>
</tr>
`, f.Title, f.Severity, f.Severity, f.Score, f.Target)
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<script src="https://unpkg.com/cytoscape/dist/cytoscape.min.js"></script>

<style>
body {
	margin:0;
	font-family:Arial;
	background:#020617;
	color:#e2e8f0;
}

/* HEADER */
.header {
	padding:20px;
	font-size:22px;
	font-weight:bold;
	border-bottom:1px solid #1e293b;
}

/* CONTAINER */
.container {
	padding:20px;
}

/* CARDS */
.cards {
	display:grid;
	grid-template-columns: repeat(4,1fr);
	gap:15px;
	margin-bottom:25px;
}

.card {
	padding:20px;
	border-radius:12px;
	text-align:center;
	font-weight:bold;
	font-size:18px;
}

.high { background:#dc2626; }
.medium { background:#d97706; }
.low { background:#2563eb; }
.info { background:#16a34a; }

/* TABLE */
.table-box {
	background:#020617;
	border:1px solid #1e293b;
	border-radius:12px;
	padding:15px;
	margin-bottom:30px;
}

table {
	width:100%%;
	border-collapse:collapse;
}

th {
	text-align:left;
	padding:12px;
	color:#94a3b8;
	border-bottom:1px solid #1e293b;
}

td {
	padding:12px;
	border-bottom:1px solid #0f172a;
}

tr:hover {
	background:#0f172a;
}

/* severity color */
.sev.HIGH { color:#ef4444; }
.sev.MEDIUM { color:#f59e0b; }
.sev.LOW { color:#3b82f6; }
.sev.INFO { color:#22c55e; }

/* GRAPH */
.graph-box {
	background:#020617;
	border:1px solid #1e293b;
	border-radius:12px;
	padding:15px;
}

#cy {
	height:600px;
	border-radius:10px;
	background:#020617;
}
</style>
</head>

<body>

<div class="header">🛡️ BlueGuard Dashboard</div>

<div class="container">

<!-- CARDS -->
<div class="cards">
	<div class="card high">HIGH<br>%d</div>
	<div class="card medium">MEDIUM<br>%d</div>
	<div class="card low">LOW<br>%d</div>
	<div class="card info">INFO<br>%d</div>
</div>

<!-- TABLE -->
<div class="table-box">
<h3>Vulnerabilities</h3>
<table>
<tr>
<th>Title</th>
<th>Severity</th>
<th>Score</th>
<th>Target</th>
</tr>
%s
</table>
</div>

<!-- GRAPH -->
<div class="graph-box">
<h3>Recon Map</h3>
<div id="cy"></div>
</div>

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

		style: [

			{
				selector: 'node',
				style: {
					'label': 'data(label)',
					'color': '#fff',
					'text-valign': 'top',
					'text-halign': 'center',
					'font-size': 12,
					'background-color': '#334155',
					'width': 35,
					'height': 35
				}
			},

			/* ROOT */
			{
				selector: 'node[type="domain"]',
				style: {
					'background-color': '#2563eb',
					'width': 50,
					'height': 50,
					'font-size': 14
				}
			},

			/* SUBDOMAIN */
			{
				selector: 'node[type="subdomain"]',
				style: {
					'background-color': '#64748b'
				}
			},

			/* IP */
			{
				selector: 'node[type="ip"]',
				style: {
					'background-color': '#22c55e'
				}
			},

			/* PORT */
			{
				selector: 'node[type="port"]',
				style: {
					'background-color': '#f59e0b',
					'shape': 'rectangle',
					'width': 60
				}
			},

			{
				selector: 'edge',
				style: {
					'line-color': '#475569',
					'width': 2
				}
			}
		],

		layout: {
			name: 'cose',
			padding: 30
		}
	});

	/* 🔥 CLICK EXPANSION */
	cy.on('tap', 'node', async function(evt){
		const node = evt.target;
		const id = node.id();

		const res = await fetch("/expand?node=" + id);
		const data = await res.json();

		data.nodes.forEach(n => {
			if(!cy.getElementById(n.id).length){
				cy.add({ data: n });
			}
		});

		data.edges.forEach(e => {
			cy.add({ data: e });
		});

		cy.layout({ name: 'cose' }).run();
	});
}

loadGraph();
</script>

</body>
</html>
`, high, medium, low, info, rows, target)

	w.Write([]byte(html))
}


//
// 🌐 RECON
//
func reconHandler(w http.ResponseWriter, r *http.Request) {

	target := r.URL.Query().Get("target")

	asset := recon.Discover(target)
	graph := recon.BuildGraph(asset)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graph)
}

//
// 🔥 EXPANSÃO
//
func expandHandler(w http.ResponseWriter, r *http.Request) {

	node := r.URL.Query().Get("node")
	typ := r.URL.Query().Get("type")

	var nodes []recon.Node
	var edges []recon.Edge

	if typ == "subdomain" {

		ips := recon.ResolveIP(node)

		for _, ip := range ips {

			id := node + "_ip_" + ip

			nodes = append(nodes, recon.Node{
				ID:         id,
				Label:      ip,
				Type:       "ip",
				Expandable: true,
			})

			edges = append(edges, recon.Edge{
				Source: node,
				Target: id,
			})
		}
	}

	if typ == "ip" {

		ports := recon.ScanPorts(node)

		for _, p := range ports {

			id := node + "_port_" + p.Port

			nodes = append(nodes, recon.Node{
				ID:         id,
				Label:      p.Port + "/" + p.Service,
				Type:       "port",
				Expandable: true,
			})

			edges = append(edges, recon.Edge{
				Source: node,
				Target: id,
			})
		}
	}

	if typ == "port" {

		version := recon.GetServiceVersion(node)

		id := node + "_ver"

		nodes = append(nodes, recon.Node{
			ID:    id,
			Label: version,
			Type:  "version",
		})

		edges = append(edges, recon.Edge{
			Source: node,
			Target: id,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recon.Graph{
		Nodes: nodes,
		Edges: edges,
	})
}
