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
	"github.com/leomarqueseh/BlueGuard/internal/dashboard"
	"github.com/leomarqueseh/BlueGuard/internal/httpclient"
	"github.com/leomarqueseh/BlueGuard/internal/i18n"
	"github.com/leomarqueseh/BlueGuard/internal/plugins"
	"github.com/leomarqueseh/BlueGuard/internal/recon"
	"github.com/leomarqueseh/BlueGuard/internal/report"
	"github.com/leomarqueseh/BlueGuard/internal/risk"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
	"github.com/leomarqueseh/BlueGuard/internal/worker"
)

//
// 🔹 Carregar lista de targets de arquivo
//
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

	// 🔥 FLAGS CLI
	target := flag.String("u", "", "Target URL")
	list := flag.String("l", "", "Targets list file")
	domain := flag.String("d", "", "Domain for recon (subdomains)")
	workers := flag.Int("w", 10, "Workers")
	timeout := flag.Int("timeout", 10, "Timeout seconds")
	userAgent := flag.String("ua", "BlueGuard", "User-Agent")
	lang := flag.String("lang", "en", "Language (en|pt-BR)")
	jsonOutput := flag.Bool("json", false, "Output JSON")
	htmlOutput := flag.String("html", "", "Output HTML file")
	web := flag.Bool("web", false, "Start web dashboard")

	flag.Parse()

	// 🌐 DASHBOARD MODE
	if *web {
		dashboard.Start()
		return
	}

	// 🌍 idioma
	i18n.Lang = *lang

	// 🚨 validação
	if *target == "" && *list == "" && *domain == "" {
		fmt.Println("Use -u OR -l OR -d")
		return
	}

	// 🔥 contexto de execução
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

	// 🔥 RECON (subdomínios)
	if *domain != "" {

		fmt.Println("[*] Running recon...")

		asset := recon.Discover(*domain)

		// 🔹 domínio principal
		targets = append(targets, scanner.Target{
			URL: asset.Domain,
		})

		// 🔹 subdomínios
		for _, sub := range asset.Subdomains {
			targets = append(targets, scanner.Target{
				URL: sub,
			})
		}

		fmt.Println("[*] Found subdomains:", len(asset.Subdomains))
	}

	// 🔥 plugins
	reg := plugins.NewRegistry()
	pluginsList := reg.All()

	fmt.Println("[*] Plugins loaded:", len(pluginsList))

	// 🔥 worker pool
	pool := worker.NewPool(pluginsList, *workers, scanCtx)

	// 🔥 executar scan
	findings := pool.Run(ctx, targets)

	// 🔥 análise de risco (score/severity)
	riskEngine := risk.New()
	findings = riskEngine.Analyze(findings)

	// 🔹 saída JSON
	if *jsonOutput {
		out, err := json.MarshalIndent(findings, "", "  ")
		if err != nil {
			fmt.Println("Error generating JSON:", err)
			return
		}
		fmt.Println(string(out))
		return
	}

	// 🔹 saída HTML
	if *htmlOutput != "" {
		err := report.GenerateHTML(findings, *htmlOutput)
		if err != nil {
			fmt.Println("Error generating HTML:", err)
			return
		}

		fmt.Println("[+] HTML report saved:", *htmlOutput)
		return
	}

	// 🔹 saída padrão CLI
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
