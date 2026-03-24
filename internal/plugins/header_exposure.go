package plugins

import (
	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/i18n"
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
			Title:       i18n.T("header_exposed_title"),
			Description: server,
			Severity:    "LOW",
			Target:      target.URL,
		})
	}

	return findings, nil
}
