package plugins

import (
	"fmt"
	"strings"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

//
// 🔹 TechFingerprint
// Detecta tecnologias via headers + body
//
type TechFingerprint struct{}

//
// 🔹 Nome do plugin (usado no CLI/dashboard)
//
func (t *TechFingerprint) Name() string {
	return "techfingerprint"
}

//
// 🔹 Execução do plugin
//
func (t *TechFingerprint) Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error) {

	var findings []scanner.Finding

	// 🔹 requisição HTTP
	resp, err := ctx.Client.Get(target.URL, ctx.UserAgent)
	if err != nil {
		return findings, nil
	}

	var techs []string

	// =============================
	// 🔹 HEADERS (map[string][]string)
	// =============================

	if resp.Headers != nil {

		// 🔹 Server
		if serverList, ok := resp.Headers["Server"]; ok && len(serverList) > 0 {
			techs = append(techs, serverList[0]) // 👈 pega o primeiro valor
		}

		// 🔹 X-Powered-By
		if poweredList, ok := resp.Headers["X-Powered-By"]; ok && len(poweredList) > 0 {
			techs = append(techs, poweredList[0])
		}
	}

	// =============================
	// 🔹 BODY
	// =============================

	body := strings.ToLower(string(resp.Body))

	if strings.Contains(body, "cloudflare") {
		techs = append(techs, "cloudflare")
	}

	if strings.Contains(body, "nginx") {
		techs = append(techs, "nginx")
	}

	if strings.Contains(body, "apache") {
		techs = append(techs, "apache")
	}

	// =============================
	// 🔹 RESULTADO
	// =============================

	if len(techs) > 0 {

		findings = append(findings, scanner.Finding{
			Title:       "Technology Fingerprint",
			Description: fmt.Sprintf("Detected: %v", techs),
			Severity:    "INFO",
			Target:      target.URL,
			Score:       1.0,
			Confirmed:   true,
		})
	}

	return findings, nil
}
