package gitexposed

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/analysis"
)

func Scan(
	baseURLs []string,
	timeout time.Duration,
) ([]Finding, error) {

	client := &http.Client{
		Timeout: timeout,
	}

	var findings []Finding

	for _, base := range baseURLs {
		base = strings.TrimRight(base, "/")

		for _, endpoint := range GitEndpoints {

			target := base + endpoint

			if analysis.Verbose {
				println("[DEBUG][GIT] Testando:", target)
			}

			req, err := http.NewRequest("GET", target, nil)
			if err != nil {
				continue
			}

			resp, err := client.Do(req)
			if err != nil || resp == nil {
				continue
			}

			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()

			if resp.StatusCode != 200 {
				continue
			}

			if isGitFile(endpoint, body) {
				findings = append(findings, Finding{
					BaseURL:  base,
					Endpoint: endpoint,
					Evidence: string(body[:min(len(body), 200)]),
					Severity: "High",
				})

				if analysis.Verbose {
					println("🚨 [GIT EXPOSED]", target)
				}
			}
		}
	}

	return findings, nil
}

func isGitFile(endpoint string, body []byte) bool {
	content := string(body)

	switch endpoint {
	case "/.git/config":
		return strings.Contains(content, "[core]") ||
			strings.Contains(content, "repositoryformatversion")

	case "/.git/HEAD":
		return strings.Contains(content, "ref:")

	case "/.git/description":
		return strings.Contains(content, "Unnamed repository")

	default:
		return false
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
