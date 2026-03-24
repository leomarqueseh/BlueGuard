package report

import (
	"fmt"
	"os"
	"strings"

	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

func GenerateHTML(findings []scanner.Finding, output string) error {

	var html strings.Builder

	html.WriteString(`
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>BlueGuard Report</title>
<style>
body {
	font-family: Arial;
	background: #0d1117;
	color: #c9d1d9;
}
.container {
	width: 80%;
	margin: auto;
}
.card {
	border-radius: 8px;
	padding: 15px;
	margin: 10px 0;
}
.high { background: #ff4d4f; }
.medium { background: #faad14; }
.low { background: #1890ff; }
.info { background: #52c41a; }
h1 { text-align: center; }
</style>
</head>
<body>
<div class="container">
<h1>BlueGuard Security Report</h1>
`)

	for _, f := range findings {

		class := strings.ToLower(f.Severity)

		html.WriteString(fmt.Sprintf(`
<div class="card %s">
	<h2>%s</h2>
	<p><strong>Target:</strong> %s</p>
	<p><strong>Severity:</strong> %s</p>
	<p><strong>Score:</strong> %.1f</p>
	<p>%s</p>
</div>
`, class, f.Title, f.Target, f.Severity, f.Score, f.Description))
	}

	html.WriteString(`
</div>
</body>
</html>
`)

	return os.WriteFile(output, []byte(html.String()), 0644)
}
