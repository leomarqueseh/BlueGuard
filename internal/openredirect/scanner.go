package openredirect

import (
	"net/http"
	"strings"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/analysis"
)

func ScanStandard(
	bases []string,
	payloadFile string,
	timeout time.Duration,
	allowedDomain string,
) ([]Finding, error) {

	payloads, err := LoadPayloads(payloadFile)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	var findings []Finding

	for _, base := range bases {
		for _, payload := range payloads {

			testURL := strings.Replace(base, "FUZZ", payload, 1)

			if analysis.Verbose {
				println("[DEBUG][STANDARD] Testando:", testURL)
			}

			req, err := http.NewRequest("GET", testURL, nil)
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
					URL:      testURL,
					Mode:     "standard",
					Payload:  payload,
					Evidence: location,
					Severity: "Medium",
				})

				if analysis.Verbose {
					println("🚨 [OPEN REDIRECT][STANDARD]", testURL, "→", location)
				}
			}
		}
	}

	return findings, nil
}
