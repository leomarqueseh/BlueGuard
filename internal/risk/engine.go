package risk

import (
	"strings"

	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (r *Engine) Analyze(findings []scanner.Finding) []scanner.Finding {

	var enhanced []scanner.Finding

	hasApache := false
	hasCloudflare := false

	for _, f := range findings {

		if f.Title == "Technology Fingerprint" {

			desc := strings.ToLower(f.Description)

			if strings.Contains(desc, "apache") {
				hasApache = true
			}

			if strings.Contains(desc, "cloudflare") {
				hasCloudflare = true
			}
		}
	}

	for _, f := range findings {

		newSeverity := f.Severity

		if f.Title == "Possible Open Redirect" && hasApache {
			newSeverity = "HIGH"
		}

		if f.Title == "Server Header Exposed" && hasCloudflare {
			newSeverity = "INFO"
		}

		f.Severity = newSeverity

		// 🔥 SCORE
		f.Score = CalculateScore(f)

		enhanced = append(enhanced, f)
	}

	return enhanced
}
