package plugins

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
)

type GitDump struct{}

func (g *GitDump) Name() string {
	return "git_dump"
}

func (g *GitDump) Run(ctx *core.ScanContext, target scanner.Target) ([]scanner.Finding, error) {

	var findings []scanner.Finding

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	base := strings.TrimRight(target.URL, "/")
	savePath := fmt.Sprintf("outputs/git/%s/.git", sanitize(target.URL))

	os.MkdirAll(savePath, os.ModePerm)

	coreFiles := []string{
		"/.git/HEAD",
		"/.git/config",
		"/.git/index",
	}

	var downloaded int

	// 🔹 base files
	for _, f := range coreFiles {
		if downloadFile(client, base+f, savePath, f, ctx.UserAgent) {
			downloaded++
		}
	}

	// 🔥 HEAD → REF
	headPath := filepath.Join(savePath, "HEAD")
	ref := ""

	if data, err := os.ReadFile(headPath); err == nil {
		content := string(data)

		if strings.Contains(content, "ref:") {
			ref = strings.TrimSpace(strings.Split(content, "ref:")[1])

			refURL := base + "/.git/" + ref
			if downloadFile(client, refURL, savePath, "/.git/"+ref, ctx.UserAgent) {
				downloaded++
			}
		}
	}

	// 🔥 commit
	if ref != "" {

		refPath := filepath.Join(savePath, ref)

		if data, err := os.ReadFile(refPath); err == nil {

			hash := strings.TrimSpace(string(data))

			if len(hash) >= 40 {

				objPath := "/.git/objects/" + hash[:2] + "/" + hash[2:]

				if downloadFile(client, base+objPath, savePath, objPath, ctx.UserAgent) {
					downloaded++
				}
			}
		}
	}

	// 🔥 index heuristic
	indexPath := filepath.Join(savePath, "index")

	if file, err := os.Open(indexPath); err == nil {

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {

			line := scanner.Text()

			if len(line) >= 40 {

				hash := line[:40]

				objPath := "/.git/objects/" + hash[:2] + "/" + hash[2:]

				if downloadFile(client, base+objPath, savePath, objPath, ctx.UserAgent) {
					downloaded++
				}
			}
		}

		file.Close()
	}

	// 🔥 EXTRAÇÃO
	ExtractGit(target.URL)

	// 🔥 RECONSTRUÇÃO
	ReconstructGit(target.URL)

	if downloaded > 0 {

		findings = append(findings, scanner.Finding{
			Title:       "Git Repository Fully Reconstructed",
			Description: fmt.Sprintf("%d git objects downloaded, extracted and reconstructed", downloaded),
			Severity:    "CRITICAL",
			Target:      target.URL,
			Score:       10.0,
			Confirmed:   true,
		})
	}

	return findings, nil
}

// helper
func downloadFile(client *http.Client, url, basePath, gitPath, ua string) bool {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", ua)

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false
	}

	body, _ := io.ReadAll(resp.Body)

	if len(body) == 0 {
		return false
	}

	local := filepath.Join(basePath, strings.TrimPrefix(gitPath, "/.git/"))

	os.MkdirAll(filepath.Dir(local), os.ModePerm)
	os.WriteFile(local, body, 0644)

	return true
}
