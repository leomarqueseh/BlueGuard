package plugins

import (
	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type HeaderExposure struct{}

func (h *HeaderExposure) Name() string {
	return "header_exposure"
}

func (h *HeaderExposure) Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error) {

	var findings []scanner.Finding

	resp, err := ctx.Client.Get(target.URL, ctx.UserAgent)
	if err != nil {
		return findings, nil
	}

	if server := resp.Headers.Get("Server"); server != "" {

		findings = append(findings, scanner.Finding{
			Title:       "Server Header Exposed",
			Description: server,
			Severity:    "LOW",
			Target:      target.URL,
		})
	}

	return findings, nil
}
