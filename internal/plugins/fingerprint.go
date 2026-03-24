package plugins

import (
	"fmt"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/fingerprint"
	"github.com/leomarqueseh/BlueGuard/internal/i18n"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type Fingerprint struct{}

func (f *Fingerprint) Name() string {
	return "fingerprint"
}

func (f *Fingerprint) Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error) {

	var findings []scanner.Finding

	resp, err := ctx.Client.Get(target.URL, ctx.UserAgent)
	if err != nil {
		return findings, nil
	}

	result := fingerprint.Detect(resp.Headers, string(resp.Body))

	if len(result.Technologies) > 0 {

		findings = append(findings, scanner.Finding{
			Title:       i18n.T("fingerprint_title"),
			Description: fmt.Sprintf("%s: %v", i18n.T("detected"), result.Technologies),
			Severity:    "INFO",
			Target:      target.URL,
		})
	}

	return findings, nil
}
