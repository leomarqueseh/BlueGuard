package openredirect

import (
	"net/http"
	"strings"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/analysis"
)

func ScanPassive(
	urls []string,
	timeout time.Duration,
	allowedDomain string,
) ([]Finding, error) {

	client := &http.Client{
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	var findings []Finding

	for _, u := range urls {

		if analysis.Verbose {
			analysis.Verbose = true
			println("[DEBUG][PASSIVE] Testando:", u)
		}

		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			continue
		}

		resp, err := client.Do(req)
		if err != nil || resp == nil {
			continue
		}
		resp.Body.Close()

		location := resp.Header.Get("Location")
		if location == "" {
			continue
		}

		if isExternalRedirect(location, allowedDomain) {
			findings = append(findings, Finding{
				URL:      u,
				Mode:     "passive",
				Evidence: location,
				Severity: "Medium",
			})

			if analysis.Verbose {
				println("🚨 [OPEN REDIRECT][PASSIVE]", u, "→", location)
			}
		}
	}

	return findings, nil
}

func isExternalRedirect(location, allowedDomain string) bool {
	loc := strings.ToLower(location)

	if strings.HasPrefix(loc, "/") {
		return false
	}

	if allowedDomain != "" && strings.Contains(loc, allowedDomain) {
		return false
	}

	return strings.HasPrefix(loc, "http") || strings.HasPrefix(loc, "//")
}
