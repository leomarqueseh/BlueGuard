package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
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
// 📂 Carrega targets a partir de arquivo
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
		line := strings.TrimSpace(sc.Text())

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

	// =========================================================
	// 🔥 FLAGS CLI
	// =========================================================
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

	// 🔥 CONTROLE DE PLUGINS
	pluginsFlag := flag.String("plugins", "", "Plugins to include (comma separated)")
	excludeFlag := flag.String("exclude", "", "Plugins to exclude (comma separated)")

	flag.Parse()

	// =========================================================
	// 🌐 DASHBOARD MODE
	// =========================================================
	if *web {
		dashboard.Start()
		return
	}

	// =========================================================
	// 🌍 IDIOMA
	// =========================================================
	i18n.Lang = *lang

	// =========================================================
	// 🚨 VALIDAÇÃO
	// =========================================================
	if *target == "" && *list == "" && *domain == "" {
		fmt.Println("Use -u OR -l OR -d")
		return
	}

	// =========================================================
	// 🔥 CONTEXTO DO SCAN
	// =========================================================
	scanCtx := &core.ScanContext{
		Timeout:   time.Duration(*timeout) * time.Second,
		UserAgent: *userAgent,
		Client:    httpclient.New(time.Duration(*timeout) * time.Second),
	}

	ctx, cancel := context.WithTimeout(context.Background(), scanCtx.Timeout)
	defer cancel()

	var targets []scanner.Target

	// =========================================================
	// 🎯 TARGET ÚNICO
	// =========================================================
	if *target != "" {
		targets = append(targets, scanner.Target{URL: *target})
	}

	// =========================================================
	// 📄 LISTA DE TARGETS
	// =========================================================
	if *list != "" {
		fileTargets, err := loadTargets(*list)
		if err != nil {
			fmt.Println(err)
			return
		}
		targets = append(targets, fileTargets...)
	}

	// =========================================================
	// 🌐 RECON (SUBDOMÍNIOS)
	// =========================================================
	if *domain != "" {

		fmt.Println("[*] Running recon...")

		// ⚠️ Novo padrão com providers (nil usa default)
		asset := recon.Discover(*domain, nil)

		// adiciona domínio principal
		targets = append(targets, scanner.Target{
			URL: asset.Domain,
		})

		// adiciona subdomínios encontrados
		for _, sub := range asset.Subdomains {
			targets = append(targets, scanner.Target{
				URL: sub,
			})
		}

		fmt.Println("[*] Found subdomains:", len(asset.Subdomains))
	}

	// =========================================================
	// 🔌 SISTEMA DE PLUGINS
	// =========================================================
	reg := plugins.NewRegistry()
	allPlugins := reg.All()

	var pluginsList []plugins.Plugin

	// 🔹 WHITELIST (-plugins)
	if *pluginsFlag != "" {

		selected := strings.Split(*pluginsFlag, ",")

		for _, p := range allPlugins {
			for _, name := range selected {
				if p.Name() == strings.TrimSpace(name) {
					pluginsList = append(pluginsList, p)
				}
			}
		}

	} else {
		// default → todos plugins
		pluginsList = allPlugins
	}

	// 🔥 BLACKLIST (-exclude)
	if *excludeFlag != "" {

		excluded := strings.Split(*excludeFlag, ",")

		var filtered []plugins.Plugin

		for _, p := range pluginsList {

			skip := false

			for _, ex := range excluded {
				if p.Name() == strings.TrimSpace(ex) {
					skip = true
					break
				}
			}

			if !skip {
				filtered = append(filtered, p)
			}
		}

		pluginsList = filtered
	}

	fmt.Println("[*] Plugins loaded:", len(pluginsList))

	// =========================================================
	// ⚡ WORKER POOL
	// =========================================================
	pool := worker.NewPool(pluginsList, *workers, scanCtx)

	// =========================================================
	// 🚀 EXECUÇÃO DO SCAN
	// =========================================================
	findings := pool.Run(ctx, targets)

	// =========================================================
	// 🧠 RISK ENGINE (score + severity)
	// =========================================================
	riskEngine := risk.New()
	findings = riskEngine.Analyze(findings)

	// =========================================================
	// 📊 RELATÓRIO INTELIGENTE (NOVO)
	// =========================================================
	var pluginNames []string
	for _, p := range pluginsList {
		pluginNames = append(pluginNames, p.Name())
	}

	mainTarget := ""
	if len(targets) > 0 {
		mainTarget = targets[0].URL
	}

	summary := report.GenerateSummary(mainTarget, findings, pluginNames)

	// =========================================================
	// 📄 JSON OUTPUT
	// =========================================================
	if *jsonOutput {
		out, err := json.MarshalIndent(findings, "", "  ")
		if err != nil {
			fmt.Println("Error generating JSON:", err)
			return
		}
		fmt.Println(string(out))
		return
	}

	// =========================================================
	// 📄 HTML OUTPUT
	// =========================================================
	if *htmlOutput != "" {
		err := report.GenerateHTML(findings, *htmlOutput)
		if err != nil {
			fmt.Println("Error generating HTML:", err)
			return
		}

		fmt.Println("[+] HTML report saved:", *htmlOutput)
		return
	}

	// =========================================================
	// 🖥️ SAÍDA CLI (PROFISSIONAL)
	// =========================================================

	// 🔥 imprime resumo SEMPRE (mesmo sem vulnerabilidade)
	fmt.Println(summary.String())

	// 🔹 detalhes individuais (se houver)
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
