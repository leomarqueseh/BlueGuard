package recon

import (
	"os/exec"
	"strings"
)

// 🔥 Resultado de porta
type Port struct {
	Port    string
	Service string
	Version string
}

// 🔍 Scan de portas via Nmap
func ScanPorts(target string) []Port {

	var results []Port

	// ⚡ comando nmap (rápido + version detection)
	cmd := exec.Command("nmap", "-sV", "-T4", "-Pn", target)

	out, err := cmd.Output()
	if err != nil {
		return results
	}

	lines := strings.Split(string(out), "\n")

	for _, line := range lines {

		// exemplo:
		// 80/tcp open http Apache 2.4.41
		if strings.Contains(line, "/tcp") && strings.Contains(line, "open") {

			fields := strings.Fields(line)

			if len(fields) < 3 {
				continue
			}

			port := strings.Split(fields[0], "/")[0]
			service := fields[2]

			version := ""
			if len(fields) > 3 {
				version = strings.Join(fields[3:], " ")
			}

			results = append(results, Port{
				Port:    port,
				Service: service,
				Version: version,
			})
		}
	}

	return results
}
