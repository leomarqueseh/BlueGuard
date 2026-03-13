package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/core"
	"github.com/leomarqueseh/BlueGuard/internal/recon"
)

func main() {

	target := flag.String("target", "", "Target URL or domain")
	timeout := flag.Int("timeout", 10, "Timeout in seconds")
	userAgent := flag.String("ua", "BlueGuard", "User-Agent")

	flag.Parse()

	if *target == "" {
		fmt.Println("Usage: blueguard -target example.com")
		return
	}

	ctx := &core.ScanContext{
		Target:    *target,
		Timeout:   time.Duration(*timeout) * time.Second,
		UserAgent: *userAgent,
	}

	results := recon.Run(ctx)

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
