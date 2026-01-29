package analysis

import "fmt"

// Run executa toda a fase de análise do BlueGuard
func Run() error {
	fmt.Println("🧠 [Analysis] Iniciando análise")

	// 1️⃣ Endpoints passivos (GAU)
	urls, err := RunEndpointCollection([]string{})
	if err != nil {
		fmt.Println("⚠️ [Analysis] GAU falhou:", err)
	}

	// 2️⃣ Endpoints ativos (Katana)
	crawled, err := RunCrawler()
	if err != nil {
		fmt.Println("⚠️ [Analysis] Katana falhou:", err)
	}

	// Futuro: merge + dedupe
	_ = append(urls, crawled...)

	fmt.Println("✅ [Analysis] Análise finalizada")
	return nil
}
