package recon

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func RunAliveCheck() ([]string, error) {
	fmt.Println("🌐 [Recon] Verificando hosts vivos (httprobe)...")

	inputFile := "outputs/recon/subdomains.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, fmt.Errorf("não foi possível abrir %s", inputFile)
	}
	defer file.Close()

	cmd := exec.Command("httprobe")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// Envia subdomínios para o httprobe
	scanner := bufio.NewScanner(file)
	go func() {
		for scanner.Scan() {
			fmt.Fprintln(stdin, scanner.Text())
		}
		stdin.Close()
	}()

	var alive []string
	outScanner := bufio.NewScanner(stdout)
	for outScanner.Scan() {
		alive = append(alive, outScanner.Text())
	}

	cmd.Wait()

	if err := saveAlive(alive); err != nil {
		return nil, err
	}

	fmt.Printf("✅ [Recon] %d hosts vivos encontrados\n", len(alive))
	return alive, nil
}

func saveAlive(hosts []string) error {
	if err := os.MkdirAll("outputs/recon", 0755); err != nil {
		return err
	}

	file, err := os.Create("outputs/recon/alive.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, host := range hosts {
		file.WriteString(host + "\n")
	}

	return nil
}
