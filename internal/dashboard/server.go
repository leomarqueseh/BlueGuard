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

// 🔥 armazenamento temporário (futuro: banco)
var results []scanner.Finding

// 🚀 inicia servidor
func Start() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/scan", scanHandler)
	http.HandleFunc("/results", resultsHandler)
	http.HandleFunc("/recon", reconHandler)

	fmt.Println("[+] Dashboard running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

//
// 🏠 HOME (ESTILO SaaS MODERNO)
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

.plugins {
	margin-top:40px;
	text-align:left;
}

.plugin-item {
	display:flex;
	align-items:center;
	gap:10px;
	margin-bottom:10px;
}

.plugin-item input {
	appearance:none;
	width:18px;
	height:18px;
	border:2px solid #22c55e;
	border-radius:4px;
	cursor:pointer;
}

.plugin-item input:checked {
	background:#22c55e;
}

.plugin-item input:checked::after {
	content:"✓";
	position:absolute;
	color:#000;
	font-size:12px;
	top:-2px;
	left:3px;
}

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

<div class="plugins">

<h3>Plugins</h3>

<div class="plugin-item">
<input type="checkbox" name="plugins" value="open_redirect" checked>
<span>Open Redirect</span>
</div>

<div class="plugin-item">
<input type="checkbox" name="plugins" value="header_exposure" checked>
<span>Header Exposure</span>
</div>

<div class="plugin-item">
<input type="checkbox" name="plugins" value="tech_fingerprint" checked>
<span>Tech Fingerprint</span>
</div>

<div class="plugin-item">
<input type="checkbox" name="plugins" value="git_exposed">
<span>Git Exposed</span>
</div>

<div class="plugin-item">
<input type="checkbox" name="plugins" value="git_dump">
<span>Git Dump</span>
</div>

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
// 📊 RESULTS + RECON MAP
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
<td>%s</td>
<td>%.1f</td>
<td>%s</td>
</tr>
`, f.Title, f.Severity, f.Score, f.Target)
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<script src="https://unpkg.com/cytoscape/dist/cytoscape.min.js"></script>

<style>
body { background:#0d1117; color:#fff; font-family:Arial; }
.container { padding:20px; }
.cards { display:flex; gap:10px; }
.card { flex:1; padding:20px; border-radius:10px; text-align:center; }
.high { background:#f85149; }
.medium { background:#d29922; }
.low { background:#58a6ff; }
.info { background:#3fb950; }
#cy { height:500px; margin-top:20px; }
table { width:100%%; margin-top:20px; }
td,th { padding:10px; border-bottom:1px solid #333; }
</style>

</head>

<body>

<div class="container">

<h1>Dashboard</h1>

<div class="cards">
<div class="card high">HIGH %d</div>
<div class="card medium">MEDIUM %d</div>
<div class="card low">LOW %d</div>
<div class="card info">INFO %d</div>
</div>

<h2>Vulnerabilities</h2>

<table>
<tr>
<th>Title</th>
<th>Severity</th>
<th>Score</th>
<th>Target</th>
</tr>
%s
</table>

<h2>Recon Map</h2>
<div id="cy"></div>

</div>

<script>
async function loadRoot(){
	const res = await fetch("/recon?target=%s");
	const data = await res.json();

	let cy = cytoscape({
		container: document.getElementById('cy'),
		elements: data.nodes.map(n=>({data:{id:n.id,label:n.label,type:n.type}}))
			.concat(data.edges.map(e=>({data:{source:e.source,target:e.target}}))),
		style: [{selector:'node',style:{'label':'data(label)','color':'#fff'}}],
		layout: { name: 'cose' }
	});
}
loadRoot();
</script>

</body>
</html>
`, high, medium, low, info, rows, target)

	w.Write([]byte(html))
}

//
// 🌐 RECON API
//
func reconHandler(w http.ResponseWriter, r *http.Request) {

	target := r.URL.Query().Get("target")

	asset := recon.Discover(target)
	graph := recon.BuildGraph(asset)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graph)
}
