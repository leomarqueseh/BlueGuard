package analysis

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func RunEndpointCollection(aliveHosts []string) ([]string, error) {
	fmt.Println("📡 [Analysis] Coletando endpoints (gau)...")

	gauCmd := exec.Command("gau", "--silent")
	gauStdin, _ := gauCmd.StdinPipe()
	gauStdout, _ := gauCmd.StdoutPipe()

	if err := gauCmd.Start(); err != nil {
		return nil, err
	}

	go func() {
		seen := make(map[string]bool)
		for _, host := range aliveHosts {
			if !seen[host] {
				fmt.Fprintln(gauStdin, host)
				seen[host] = true
			}
		}
		gauStdin.Close()
	}()

	var urls []string
	scanner := bufio.NewScanner(gauStdout)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	gauCmd.Wait()

	saveFile("outputs/analysis/urls_raw.txt", urls)
	fmt.Printf("✅ [Analysis] %d URLs coletadas\n", len(urls))

	return urls, nil
}

// util simples (caso não exista ainda)
func saveFile(path string, data []string) {
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()

	for _, line := range data {
		fmt.Fprintln(f, line)
	}
}
