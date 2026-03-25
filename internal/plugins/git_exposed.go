package plugins

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type GitExposed struct{}

func (g *GitExposed) Name() string {
	return "git_exposed"
}

func (g *GitExposed) Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error) {

	var findings []scanner.Finding

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	paths := []string{
		"/.git/HEAD",
		"/.git/config",
	}

	for _, path := range paths {

		url := strings.TrimRight(target.URL, "/") + path

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", ctx.UserAgent)

		resp, err := client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// precisa ser 200 OK
		if resp.StatusCode != 200 {
			continue
		}

		bodyBytes, _ := io.ReadAll(resp.Body)
		body := string(bodyBytes)

		// 🔥 VALIDAÇÃO REAL
		if strings.Contains(body, "ref: refs/heads") ||
			strings.Contains(body, "[core]") {

			findings = append(findings, scanner.Finding{
				Title:       "Git Repository Exposed (CONFIRMED)",
				Description: fmt.Sprintf("Sensitive git file exposed: %s", path),
				Severity:    "HIGH",
				Target:      url,
				Score:       9.5,
				Confirmed:   true,
			})

			return findings, nil
		}
	}

	return findings, nil
}
