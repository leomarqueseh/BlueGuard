package analysis

import (
	"bufio"
	"fmt"
	"os/exec"
)

func RunCrawler() ([]string, error) {
	fmt.Println("🕷️ [Analysis] Crawling ativo (katana)...")

	cmd := exec.Command(
		"katana",
		"-list", "outputs/recon/alive.txt",
		"-silent",
	)

	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	var urls []string
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	cmd.Wait()

	saveFile("outputs/analysis/urls_crawled.txt", urls)
	fmt.Printf("✅ [Analysis] %d URLs descobertas via crawling\n", len(urls))

	return urls, nil
}
