package report

import (
	"fmt"

	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type Summary struct {
	Target          string
	TotalFindings   int
	High            int
	Medium          int
	Low             int
	Info            int
	Recommendations []string
}

// 🔥 Gera resumo inteligente
func GenerateSummary(target string, findings []scanner.Finding, plugins []string) Summary {

	var s Summary
	s.Target = target
	s.TotalFindings = len(findings)

	// contar severidade
	for _, f := range findings {
		switch f.Severity {
		case "HIGH":
			s.High++
		case "MEDIUM":
			s.Medium++
		case "LOW":
			s.Low++
		default:
			s.Info++
		}
	}

	// 🔥 recomendações baseadas no scan
	if s.TotalFindings == 0 {

		s.Recommendations = append(s.Recommendations,
			"No vulnerabilities found, but this does NOT mean the system is secure.",
			"Run authenticated scans for deeper analysis.",
			"Enable continuous monitoring.",
			"Test with aggressive profile.",
		)

	} else {

		s.Recommendations = append(s.Recommendations,
			"Fix HIGH vulnerabilities immediately.",
			"Review exposed headers and configurations.",
			"Implement WAF protection.",
		)
	}

	return s
}

// 🔥 versão texto CLI
func (s Summary) String() string {

	return fmt.Sprintf(`
==============================
🛡️ BLUEGUARD REPORT
==============================

Target: %s

Findings:
HIGH: %d
MEDIUM: %d
LOW: %d
INFO: %d

Total: %d

Recommendations:
- %s
`, s.Target, s.High, s.Medium, s.Low, s.Info, s.TotalFindings, formatList(s.Recommendations))
}

func formatList(items []string) string {
	out := ""
	for _, i := range items {
		out += "\n- " + i
	}
	return out
}
