package recon

import (
	"bufio"
	"net"
	"os"
)

// Bruteforce tenta resolver subdomínios via wordlist
func Bruteforce(domain string, wordlist string) []string {

	var results []string

	file, err := os.Open(wordlist)
	if err != nil {
		return results
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		sub := scanner.Text()

		if sub == "" {
			continue
		}

		host := sub + "." + domain

		_, err := net.LookupHost(host)

		if err == nil {
			results = append(results, "https://"+host)
		}
	}

	return results
}
