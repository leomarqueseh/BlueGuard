package plugins

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type OpenRedirect struct{}

func (o *OpenRedirect) Name() string {
	return "open_redirect"
}

func (o *OpenRedirect) Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error) {

	var findings []scanner.Finding

	parsed, err := url.Parse(target.URL)
	if err != nil {
		return findings, nil
	}

	query := parsed.Query()

	params := []string{"url", "redirect", "next", "dest", "return", "redir"}

	// 🔥 client REAL com redirect control
	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// não seguir redirect automaticamente
			return http.ErrUseLastResponse
		},
	}

	for _, p := range params {

		query.Set(p, "https://evil.com")
		parsed.RawQuery = query.Encode()

		testURL := parsed.String()

		req, _ := http.NewRequest("GET", testURL, nil)
		req.Header.Set("User-Agent", ctx.UserAgent)

		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// 🔥 VALIDAÇÃO REAL
		location := resp.Header.Get("Location")

		if strings.Contains(location, "evil.com") {
			findings = append(findings, scanner.Finding{
				Title:       "Open Redirect (CONFIRMED)",
				Description: fmt.Sprintf("Confirmed via parameter '%s' redirecting to external domain", p),
				Severity:    "HIGH",
				Target:      testURL,
				Score:       9.0,
				Confirmed:   true,
			})
			return findings, nil
		}

		if resp.StatusCode >= 300 && resp.StatusCode < 400 {
			findings = append(findings, scanner.Finding{
				Title:       "Possible Open Redirect",
				Description: fmt.Sprintf("Parameter '%s' may allow redirection (not confirmed)", p),
				Severity:    "MEDIUM",
				Target:      testURL,
				Score:       6.5,
				Confirmed:   false,
			})
		}
	}

	return findings, nil
}
