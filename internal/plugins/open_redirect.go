package plugins

import (
	"strings"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/i18n"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type OpenRedirect struct{}

func (o *OpenRedirect) Name() string {
	return "open_redirect"
}

func (o *OpenRedirect) Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error) {

	var findings []scanner.Finding

	testURL := target.URL + "?redirect=https://evil.com"

	resp, err := ctx.Client.Get(testURL, ctx.UserAgent)
	if err != nil {
		return findings, nil
	}

	if strings.Contains(resp.URL, "evil.com") {

		findings = append(findings, scanner.Finding{
			Title:       i18n.T("open_redirect_title"),
			Description: i18n.T("open_redirect_desc"),
			Severity:    "MEDIUM",
			Target:      target.URL,
		})
	}

	return findings, nil
}
