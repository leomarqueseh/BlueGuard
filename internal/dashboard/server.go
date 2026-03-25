package dashboard

import (
	"fmt"
	"net/http"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/httpclient"
	"github.com/leomarqueseh/BlueGuard/internal/plugins"
	"github.com/leomarqueseh/BlueGuard/internal/risk"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
	"github.com/leomarqueseh/BlueGuard/internal/worker"
)

var results []scanner.Finding

func Start() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/scan", scanHandler)
	http.HandleFunc("/results", resultsHandler)

	fmt.Println("[+] Dashboard running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// HOME
func homeHandler(w http.ResponseWriter, r *http.Request) {

	html := `
<!DOCTYPE html>
<html>
<head>
<title>BlueGuard</title>
<style>
body {
	background:#0d1117;
	color:#c9d1d9;
	font-family:Arial;
	text-align:center;
}
input {
	padding:12px;
	width:320px;
	border-radius:6px;
	border:none;
}
button {
	padding:12px;
	background:#238636;
	color:white;
	border:none;
	border-radius:6px;
	cursor:pointer;
}
</style>
</head>
<body>

<h1>BlueGuard Scanner</h1>

<form action="/scan">
<input name="target" placeholder="https://example.com"/>
<button>Scan</button>
</form>

<br>
<a href="/results">Ver resultados</a>

</body>
</html>
`
	w.Write([]byte(html))
}

// SCAN
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

	http.Redirect(w, r, "/results", http.StatusSeeOther)
}

// RESULTS (NESSUS STYLE)
func resultsHandler(w http.ResponseWriter, r *http.Request) {

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

		color := "#3fb950"

		switch f.Severity {
		case "HIGH":
			color = "#f85149"
		case "MEDIUM":
			color = "#d29922"
		case "LOW":
			color = "#58a6ff"
		}

		rows += fmt.Sprintf(`
<tr>
<td>%s</td>
<td><span style="color:%s;font-weight:bold;">%s</span></td>
<td>%.1f</td>
<td><a href="%s" target="_blank">%s</a></td>
</tr>
`, f.Title, color, f.Severity, f.Score, f.Target, f.Target)
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>BlueGuard Dashboard</title>

<style>
body {
	margin:0;
	font-family:Arial;
	background:#0d1117;
	color:#c9d1d9;
	display:flex;
}

/* SIDEBAR */
.sidebar {
	width:220px;
	background:#161b22;
	height:100vh;
	padding:20px;
}
.sidebar h2 {
	color:#58a6ff;
}
.sidebar a {
	display:block;
	color:#c9d1d9;
	margin:10px 0;
	text-decoration:none;
}

/* MAIN */
.main {
	flex:1;
	padding:30px;
}

/* TITLE */
h1 {
	margin-bottom:20px;
}

/* CARDS */
.cards {
	display:flex;
	gap:20px;
	margin-bottom:30px;
}
.card {
	flex:1;
	padding:20px;
	border-radius:10px;
	text-align:center;
	font-weight:bold;
	font-size:18px;
}
.high { background:#f85149; }
.medium { background:#d29922; }
.low { background:#58a6ff; }
.info { background:#3fb950; }

/* TABLE */
table {
	width:100%%;
	border-collapse:collapse;
	background:#161b22;
	border-radius:10px;
	overflow:hidden;
}
th {
	background:#21262d;
	padding:12px;
	text-align:left;
}
td {
	padding:12px;
	border-bottom:1px solid #30363d;
}
tr:hover {
	background:#21262d;
}

a {
	color:#58a6ff;
	text-decoration:none;
}
</style>

</head>
<body>

<div class="sidebar">
	<h2>BlueGuard</h2>
	<a href="/">🏠 Home</a>
	<a href="/results">📊 Results</a>
</div>

<div class="main">

<h1>Security Dashboard</h1>

<div class="cards">
	<div class="card high">HIGH<br>%d</div>
	<div class="card medium">MEDIUM<br>%d</div>
	<div class="card low">LOW<br>%d</div>
	<div class="card info">INFO<br>%d</div>
</div>

<table>
<tr>
<th>Vulnerability</th>
<th>Severity</th>
<th>Score</th>
<th>Target</th>
</tr>
%s
</table>

</div>

</body>
</html>
`, high, medium, low, info, rows)

	w.Write([]byte(html))
}
