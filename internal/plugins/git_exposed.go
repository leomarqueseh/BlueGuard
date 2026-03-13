package plugins

import (
	"context"
	"net/http"

	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type GitExposed struct{}

func (p *GitExposed) Name() string {
	return "git_exposed"
}

func (p *GitExposed) Run(ctx context.Context, target scanner.Target) ([]scanner.Finding, error) {

	url := target.URL + "/.git/HEAD"

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		finding := scanner.Finding{
			Title:       "Git Repository Exposed",
			Description: "The .git repository is accessible",
			Severity:    "HIGH",
			Target:      target.URL,
			Evidence:    url,
		}

		return []scanner.Finding{finding}, nil
	}

	return nil, nil
}
