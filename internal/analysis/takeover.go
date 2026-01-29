package analysis

import (
	"bufio"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type TakeoverFinding struct {
	Host     string
	Provider string
	Evidence string
}

func RunTakeover(filePath string, timeout time.Duration) ([]TakeoverFinding, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hosts := make(map[string]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "http") {
			u, err := url.Parse(line)
			if err == nil && u.Hostname() != "" {
				hosts[u.Hostname()] = true
			}
		} else {
			hosts[line] = true
		}
	}

	client := http.Client{Timeout: timeout}
	var results []TakeoverFinding

	for host := range hosts {
		cname, err := net.LookupCNAME(host)
		if err != nil {
			continue
		}

		provider := detectProvider(cname)
		if provider == "" {
			continue
		}

		resp, err := client.Get("http://" + host)
		if err != nil {
			continue
		}

		buf := make([]byte, 2048)
		resp.Body.Read(buf)
		resp.Body.Close()

		if fingerprint(provider, string(buf)) {
			results = append(results, TakeoverFinding{
				Host:     host,
				Provider: provider,
				Evidence: "Unclaimed service response",
			})
		}
	}

	return results, nil
}

func detectProvider(cname string) string {
	switch {
	case strings.Contains(cname, "amazonaws.com"):
		return "AWS S3"
	case strings.Contains(cname, "github.io"):
		return "GitHub Pages"
	case strings.Contains(cname, "herokudns.com"):
		return "Heroku"
	case strings.Contains(cname, "azurewebsites.net"):
		return "Azure"
	}
	return ""
}

func fingerprint(provider, body string) bool {
	switch provider {
	case "AWS S3":
		return strings.Contains(body, "NoSuchBucket")
	case "GitHub Pages":
		return strings.Contains(body, "There isn't a GitHub Pages site here")
	case "Heroku":
		return strings.Contains(body, "No such app")
	case "Azure":
		return strings.Contains(body, "404 Web Site not found")
	}
	return false
}
