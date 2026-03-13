package plugins

import (
	"context"
	"net/http"

	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type HeaderExposure struct{}

func (p *HeaderExposure) Name() string {
	return "header_exposure"
}

func (p *HeaderExposure) Run(ctx context.Context, target scanner.Target) ([]scanner.Finding, error) {

	resp, err := http.Get(target.URL)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	server := resp.Header.Get("Server")

	if server != "" {

		finding := scanner.Finding{
			Title:       "Server Header Exposed",
			Description: "Server header leaks backend technology",
			Severity:    "LOW",
			Target:      target.URL,
			Evidence:    server,
		}

		return []scanner.Finding{finding}, nil
	}

	return nil, nil
}
