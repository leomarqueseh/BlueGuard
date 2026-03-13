package plugins

import (
	"context"
	"net/http"

	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type OpenRedirect struct{}

func (p *OpenRedirect) Name() string {
	return "open_redirect"
}

func (p *OpenRedirect) Run(ctx context.Context, target scanner.Target) ([]scanner.Finding, error) {

	testURL := target.URL + "?redirect=https://evil.com"

	resp, err := http.Get(testURL)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 {

		finding := scanner.Finding{
			Title:       "Possible Open Redirect",
			Description: "Endpoint may allow open redirect",
			Severity:    "MEDIUM",
			Target:      target.URL,
			Evidence:    testURL,
		}

		return []scanner.Finding{finding}, nil
	}

	return nil, nil
}
