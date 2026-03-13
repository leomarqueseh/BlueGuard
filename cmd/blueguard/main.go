package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/recon"
)

func main() {

	target := flag.String("u", "", "Target URL or domain")
	timeout := flag.Int("timeout", 10, "Timeout in seconds")
	userAgent := flag.String("ua", "BlueGuard", "User-Agent")
	jsonOutput := flag.Bool("json", false, "Output JSON")

	flag.Parse()

	if *target == "" {
		fmt.Println("BlueGuard Security Scanner")
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("  blueguard -u example.com")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  -u        Target URL or domain")
		fmt.Println("  -timeout  Timeout in seconds")
		fmt.Println("  -ua       Custom User-Agent")
		fmt.Println("  -json     Output JSON format")
		return
	}

	ctx := &core.ScanContext{
		Target:    *target,
		Timeout:   time.Duration(*timeout) * time.Second,
		UserAgent: *userAgent,
	}

	results := recon.Run(ctx)

	// JSON output (placeholder para futura implementação)
	if *jsonOutput {
		fmt.Println("[JSON output coming soon]")
		return
	}

	// Output padrão
	for _, f := range results {

		fmt.Printf(
			"\n[%s] %s\nTarget: %s\nConfidence: %d%%\n%s\n",
			f.Severity,
			f.Title,
			f.Target,
			f.Confidence,
			f.Description,
		)

	}
}
