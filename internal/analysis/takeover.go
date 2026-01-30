package analysis

import (
	"bufio"
	"context"
	"net"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/httpx"
)

type TakeoverFinding struct {
	Host     string
	Provider string
	Evidence string
}

// RunTakeover executa takeover passivo com DNS + HTTP controlados
func RunTakeover(filePath string, timeout time.Duration) ([]TakeoverFinding, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hosts := make(map[string]struct{})
	scanner := bufio.NewScanner(file)

	// 🔹 NORMALIZAÇÃO DE INPUT
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "http://") || strings.HasPrefix(line, "https://") {
			u, err := url.Parse(line)
			if err != nil || u.Hostname() == "" {
				if Verbose {
					logf("URL inválida ignorada: %s", line)
				}
				continue
			}
			hosts[u.Hostname()] = struct{}{}
		} else {
			hosts[line] = struct{}{}
		}
	}

	var findings []TakeoverFinding

	// 🔹 DNS resolver com timeout
	resolver := &net.Resolver{}

	for host := range hosts {

		if Verbose {
			logf("Analisando host: %s", host)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		cname, err := resolver.LookupCNAME(ctx, host)
		cancel()

		if err != nil || cname == "" {
			if Verbose {
				logf("CNAME não resolvido para %s", host)
			}
			continue
		}

		if Verbose {
			logf("CNAME encontrado: %s → %s", host, cname)
		}

		fp := DetectFingerprintByCNAME(cname)
		if fp == nil {
			if Verbose {
				logf("Nenhum fingerprint compatível para %s", cname)
			}
			continue
		}

		if Verbose {
			logf("Fingerprint candidato: %s (%s)", fp.Provider, host)
		}

		// 🔹 HTTP controlado (HTTPS → HTTP)
		resp, err := httpx.Fetch(host, timeout)
		if err != nil || resp == nil {
			if Verbose {
				logf("HTTP não respondeu para %s", host)
			}
			continue
		}

		// 🔹 MATCHERS AND + NEGATIVE
		if MatchFingerprint(fp, resp) {
			findings = append(findings, TakeoverFinding{
				Host:     host,
				Provider: fp.Provider,
				Evidence: "Fingerprint match (CNAME + HTTP)",
			})

			if Verbose {
				logf("🚨 POSSÍVEL TAKEOVER: %s (%s)", host, fp.Provider)
			}
		} else if Verbose {
			logf("HTTP respondeu, mas fingerprint não confirmou (%s)", host)
		}
	}

	return findings, nil
}
