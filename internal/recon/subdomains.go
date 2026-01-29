package recon

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func RunSubdomainEnum(target string) ([]string, error) {
	fmt.Println("🔎 [Recon] Enumerando subdomínios...")

	subdomains := make(map[string]bool)

	tools := [][]string{
		{"subfinder", "-silent", "-d", target},
		{"assetfinder", "--subs-only", target},
	}

	for _, tool := range tools {
		cmd := exec.Command(tool[0], tool[1:]...)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return nil, err
		}

		if err := cmd.Start(); err != nil {
			return nil, err
		}

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			sub := scanner.Text()
			subdomains[sub] = true
		}

		cmd.Wait()
	}

	var result []string
	for sub := range subdomains {
		result = append(result, sub)
	}

	if err := saveSubdomains(result); err != nil {
		return nil, err
	}

	fmt.Printf("✅ [Recon] %d subdomínios encontrados\n", len(result))
	return result, nil
}

func saveSubdomains(subs []string) error {
	if err := os.MkdirAll("outputs/recon", 0755); err != nil {
		return err
	}

	file, err := os.Create("outputs/recon/subdomains.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, sub := range subs {
		file.WriteString(sub + "\n")
	}

	return nil
}
