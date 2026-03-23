package plugins

import (
	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type GitExposed struct{}

func (g *GitExposed) Name() string {
	return "git_exposed"
}

func (g *GitExposed) Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error) {

	var findings []scanner.Finding

	url := target.URL + "/.git/config"

	resp, err := ctx.Client.Get(url, ctx.UserAgent)
	if err != nil {
		return findings, nil
	}

	if resp.StatusCode == 200 {

		findings = append(findings, scanner.Finding{
			Title:       "Git Repository Exposed",
			Description: ".git/config accessible",
			Severity:    "HIGH",
			Target:      target.URL,
		})
	}

	return findings, nil
}
