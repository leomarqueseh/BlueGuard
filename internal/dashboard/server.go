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

// 🚀 Start server
func Start() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/scan", scanHandler)
	http.HandleFunc("/results", resultsHandler)

	fmt.Println("[+] Dashboard running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// 🔹 HOME
func homeHandler(w http.ResponseWriter, r *http.Request) {

	html := `
<!DOCTYPE html>
<html>
<head>
<title>BlueGuard Dashboard</title>
<style>
body {
	font-family: Arial;
	background:#0d1117;
	color:#c9d1d9;
	text-align:center;
}
input {
	padding:10px;
	width:300px;
	border-radius:5px;
	border:none;
}
button {
	padding:10px;
	border:none;
	background:#58a6ff;
	color:white;
	border-radius:5px;
	cursor:pointer;
}
</style>
</head>
<body>

<h1>BlueGuard Scanner</h1>

<form action="/scan">
<input name="target" placeholder="https://example.com"/>
<button type="submit">Scan</button>
</form>

<br>
<a href="/results">View Results</a>

</body>
</html>
`
	w.Write([]byte(html))
}

// 🔹 SCAN
func scanHandler(w http.ResponseWriter, r *http.Request) {

	target := r.URL.Query().Get("target")

	if target == "" {
		w.Write([]byte("Missing target"))
		return
	}

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

	riskEngine := risk.New()
	findings = riskEngine.Analyze(findings)

	// 🔥 evita null
	if findings == nil {
		findings = []scanner.Finding{}
	}

	results = findings

	http.Redirect(w, r, "/results", http.StatusSeeOther)
}

// 🔹 RESULTS (UI PROFISSIONAL)
func resultsHandler(w http.ResponseWriter, r *http.Request) {

	if results == nil {
		results = []scanner.Finding{}
	}

	var cards string

	// 🔥 contadores
	var high, medium, low, info int

	for _, f := range results {

		color := "#52c41a"
		badge := "INFO"

		switch f.Severity {
		case "HIGH":
			color = "#ff4d4f"
			badge = "HIGH"
			high++
		case "MEDIUM":
			color = "#faad14"
			badge = "MEDIUM"
			medium++
		case "LOW":
			color = "#1890ff"
			badge = "LOW"
			low++
		default:
			info++
		}

		cards += fmt.Sprintf(`
		<div class="card" style="border-left: 5px solid %s;">
			<div class="badge" style="background:%s;">%s</div>
			<h2>%s</h2>
			<p><strong>Target:</strong> %s</p>
			<p><strong>Score:</strong> %.1f</p>
			<p>%s</p>
		</div>
		`, color, color, badge, f.Title, f.Target, f.Score, f.Description)
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<title>BlueGuard Results</title>
<style>
body {
	font-family: Arial;
	background: #0d1117;
	color: #c9d1d9;
}
.container {
	width: 80%%;
	margin: auto;
}
.header {
	text-align: center;
	margin-bottom: 20px;
}
.summary span {
	margin: 0 10px;
	font-weight: bold;
}
.card {
	background: #161b22;
	padding: 20px;
	margin: 15px 0;
	border-radius: 10px;
	box-shadow: 0 0 10px rgba(0,0,0,0.5);
	position: relative;
}
.badge {
	position: absolute;
	top: 15px;
	right: 15px;
	padding: 5px 10px;
	border-radius: 5px;
	color: white;
	font-size: 12px;
}
h1 {
	text-align: center;
}
a {
	color: #58a6ff;
}
</style>
</head>
<body>

<div class="container">

<div class="header">
	<h1>Scan Results</h1>
	<div class="summary">
		<span style="color:#ff4d4f;">HIGH: %d</span>
		<span style="color:#faad14;">MEDIUM: %d</span>
		<span style="color:#1890ff;">LOW: %d</span>
		<span style="color:#52c41a;">INFO: %d</span>
	</div>
</div>

%s

<br>
<a href="/">← Back</a>

</div>

</body>
</html>
`, high, medium, low, info, cards)

	w.Write([]byte(html))
}
