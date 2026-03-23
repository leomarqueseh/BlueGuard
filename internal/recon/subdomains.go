package recon

import (
	"fmt"
)

// Discover executa a descoberta de subdomínios para um domínio alvo
func Discover(domain string) []string {

	var results []string

	fmt.Println("[*] Starting subdomain discovery for:", domain)

	// 🔹 Bruteforce básico com wordlist
	subs := Bruteforce(domain, "subs.txt")

	for _, s := range subs {
		results = append(results, s)
	}

	fmt.Println("[*] Discovery finished")

	return unique(results)
}

// 🔥 Remove duplicados (importante para evitar scans repetidos)
func unique(input []string) []string {

	keys := make(map[string]bool)
	var list []string

	for _, entry := range input {

		if _, exists := keys[entry]; !exists {

			keys[entry] = true
			list = append(list, entry)

		}
	}

	return list
}
