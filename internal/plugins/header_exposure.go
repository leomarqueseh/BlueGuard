package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/yourusername/blueguard/internal/scanner"
)

type HeaderExposure struct{}

func (p *HeaderExposure) ID() string {
	return "BG-HTTP-001"
}

func (p *HeaderExposure) Name() string {
	return "Sensitive Header Exposure"
}

func (p *HeaderExposure) Description() string {
	return "Detects exposure of sensitive HTTP headers."
}

func (p *HeaderExposure) Run(ctx context.Context, target scanner.Target) ([]scanner.Finding, error) {

	client := scanner.NewHTTPClient()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target.URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var findings []scanner.Finding

	if server := resp.Header.Get("Server"); server != "" {
		findings = append(findings, scanner.Finding{
			PluginID:    p.ID(),
			Title:       "Server Header Exposed",
			Description: "The server header reveals backend technology.",
			Severity:    "Low",
			Evidence:    server,
			Timestamp:   time.Now(),
		})
	}

	return findings, nil
}
