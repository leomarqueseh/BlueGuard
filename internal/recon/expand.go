package recon

import (
	"net"
	"strings"
)

// 🌐 RESOLVE IP
func ResolveIP(domain string) []string {

	var ips []string

	resolved, err := net.LookupIP(domain)
	if err != nil {
		return ips
	}

	for _, ip := range resolved {
		ips = append(ips, ip.String())
	}

	return ips
}

// 🧠 VERSÃO DO SERVIÇO (placeholder)
func GetServiceVersion(node string) string {

	if strings.Contains(node, "80") {
		return "HTTP Service"
	}

	if strings.Contains(node, "443") {
		return "HTTPS Service"
	}

	if strings.Contains(node, "22") {
		return "SSH Service"
	}

	return "Unknown version"
}
