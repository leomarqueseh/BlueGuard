package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/httpclient"
	"github.com/leomarqueseh/BlueGuard/internal/plugins"
	"github.com/leomarqueseh/BlueGuard/internal/recon"
	"github.com/leomarqueseh/BlueGuard/internal/scanner"
	"github.com/leomarqueseh/BlueGuard/internal/worker"
)

func loadTargets(file string) ([]scanner.Target, error) {

	var targets []scanner.Target

	f, err := os.Open(file)
	if err != nil {
		return targets, err
	}
	defer f.Close()

	scannerFile := bufio.NewScanner(f)

	for scannerFile.Scan() {
		line := scannerFile.Text()
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

	target := flag.String("u", "", "Target URL")
	list := flag.String("l", "", "Targets list file")
	domain := flag.String("d", "", "Domain for subdomain discovery")
	workers := flag.Int("w", 10, "Number of workers")
	timeout := flag.Int("timeout", 10, "Timeout seconds")
	userAgent := flag.String("ua", "BlueGuard", "User-Agent")

	flag.Parse()

	if *target == "" && *list == "" && *domain == "" {
		fmt.Println("Use -u OR -l OR -d")
		return
	}

	// 🔥 ScanContext PROFISSIONAL
	scanCtx := &core.ScanContext{
		Timeout:   time.Duration(*timeout) * time.Second,
		UserAgent: *userAgent,
		Client:    httpclient.New(time.Duration(*timeout) * time.Second),
	}

	ctx, cancel := context.WithTimeout(context.Background(), scanCtx.Timeout)
	defer cancel()

	var targets []scanner.Target

	if *target != "" {
		targets = append(targets, scanner.Target{URL: *target})
	}

	if *list != "" {
		fileTargets, err := loadTargets(*list)
		if err != nil {
			fmt.Println(err)
			return
		}
		targets = append(targets, fileTargets...)
	}

	if *domain != "" {
		fmt.Println("[*] Discovering subdomains...")
		discovered := recon.Discover(*domain)

		for _, d := range discovered {
			targets = append(targets, scanner.Target{URL: d})
		}

		fmt.Println("[*] Found:", len(discovered))
	}

	reg := plugins.NewRegistry()

	pool := worker.NewPool(reg.All(), *workers, scanCtx)

	findings := pool.Run(ctx, targets)

	for _, f := range findings {

		fmt.Printf(
			"\n[%s] %s\nTarget: %s\n%s\n",
			f.Severity,
			f.Title,
			f.Target,
			f.Description,
		)
	}
}
