package openredirect

import (
	"bufio"
	"os"
	"strings"
)

func LoadPayloads(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var payloads []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		payloads = append(payloads, line)
	}

	return payloads, nil
}
