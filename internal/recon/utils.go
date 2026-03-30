package recon

import "strings"

func cleanDomain(input string) string {

	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "http://", "")
	input = strings.ReplaceAll(input, "https://", "")
	input = strings.TrimSuffix(input, "/")

	return input
}
