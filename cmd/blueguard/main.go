package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/httpclient"
	"github.com/leomarqueseh/BlueGuard/internal/i18n"
	"github.com/leomarqueseh/BlueGuard/internal/plugins"
	"github.com/leomarqueseh/BlueGuard/internal/recon"
	"github.com/leomarqueseh/BlueGuard/internal/risk"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
	"github.com/leomarqueseh/BlueGuard/internal/worker"
)

// 🔹 Carregar targets de arquivo
func loadTargets(file string) ([]scanner.Target, error) {

	var targets []scanner.Target

	f, err := os.Open(file)
	if err != nil {
		return targets, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			continue
		}

		targets = append(targets, scanner.Target{
			URL: line,
		})
	}

	return targets, nil
}

func main() {

	// 🔹 Flags
	target := flag.String("u", "", "Target URL")
	list := flag.String("l", "", "Targets list file")
	domain := flag.String("d", "", "Domain for discovery")
	workers := flag.Int("w", 10, "Workers")
	timeout := flag.Int("timeout", 10, "Timeout seconds")
	userAgent := flag.String("ua", "BlueGuard", "User-Agent")
	lang := flag.String("lang", "en", "Language (en|pt-BR)")
	jsonOutput := flag.Bool("json", false, "Output JSON")

	flag.Parse()

	// 🔥 idioma
	i18n.Lang = *lang

	// 🔴 validação
	if *target == "" && *list == "" && *domain == "" {
		fmt.Println("Use -u OR -l OR -d")
		return
	}

	// 🔥 contexto do scan
	scanCtx := &core.ScanContext{
		Timeout:   time.Duration(*timeout) * time.Second,
		UserAgent: *userAgent,
		Client:    httpclient.New(time.Duration(*timeout) * time.Second),
	}

	ctx, cancel := context.WithTimeout(context.Background(), scanCtx.Timeout)
	defer cancel()

	var targets []scanner.Target

	// 🔹 target único
	if *target != "" {
		targets = append(targets, scanner.Target{URL: *target})
	}

	// 🔹 lista de targets
	if *list != "" {
		fileTargets, err := loadTargets(*list)
		if err != nil {
			fmt.Println(err)
			return
		}
		targets = append(targets, fileTargets...)
	}

	// 🔹 descoberta de subdomínios
	if *domain != "" {

		fmt.Println("[*] Discovering subdomains...")

		discovered := recon.Discover(*domain)

		for _, d := range discovered {
			targets = append(targets, scanner.Target{URL: d})
		}

		fmt.Println("[*] Found:", len(discovered))
	}

	// 🔥 plugins
	reg := plugins.NewRegistry()
	fmt.Println("[*] Plugins loaded:", len(reg.All()))

	// 🔥 worker pool
	pool := worker.NewPool(reg.All(), *workers, scanCtx)

	// 🔥 execução do scan
	findings := pool.Run(ctx, targets)

	// 🔥 risk engine + score
	riskEngine := risk.New()
	findings = riskEngine.Analyze(findings)

	// 🔹 JSON output (PROFISSIONAL)
	if *jsonOutput {
		out, err := json.MarshalIndent(findings, "", "  ")
		if err != nil {
			fmt.Println("Error generating JSON:", err)
			return
		}
		fmt.Println(string(out))
		return
	}

	// 🔹 output padrão
	for _, f := range findings {

		fmt.Printf(
			"\n[%s] %s\nTarget: %s\nScore: %.1f\n%s\n",
			i18n.Severity(f.Severity),
			f.Title,
			f.Target,
			f.Score,
			f.Description,
		)
	}
}
