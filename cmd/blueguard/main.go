package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/leomarqueseh/BlueGuard/internal/analysis"
	"github.com/leomarqueseh/BlueGuard/internal/recon"
)

func banner() {
	fmt.Println(`
██████╗ ██╗     ██╗   ██╗███████╗ ██████╗ ██╗   ██╗ █████╗ ██████╗ ██████╗
██╔══██╗██║     ██║   ██║██╔════╝██╔════╝ ██║   ██║██╔══██╗██╔══██╗██╔══██╗
██████╔╝██║     ██║   ██║█████╗  ██║  ███╗██║   ██║███████║██████╔╝██║  ██║
██╔══██╗██║     ██║   ██║██╔══╝  ██║   ██║██║   ██║██╔══██║██╔══██╗██║  ██║
██████╔╝███████╗╚██████╔╝███████╗╚██████╔╝╚██████╔╝██║  ██║██║  ██║██████╔╝
╚═════╝ ╚══════╝ ╚═════╝ ╚══════╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝
BlueGuard - Security Recon & Takeover Scanner
`)
}

func main() {
	var (
		target   string
		listFile string

		passiveLite bool
		passive     bool
		active      bool
		full        bool
		stealth     bool
		verbose     bool

		rate    int
		delay   time.Duration
		timeout time.Duration
	)

	flag.StringVar(&target, "t", "", "Target único (ex: example.com)")
	flag.StringVar(&listFile, "list", "", "Arquivo .txt com domínios ou URLs")

	flag.BoolVar(&passiveLite, "passive-lite", false, "Apenas takeover (rápido e seguro)")
	flag.BoolVar(&passive, "passive", false, "Recon passivo + takeover")
	flag.BoolVar(&active, "active", false, "Scan ativo (em evolução)")
	flag.BoolVar(&full, "full", false, "Scan completo (planejado)")
	flag.BoolVar(&stealth, "stealth", false, "Modo silencioso")
	flag.BoolVar(&verbose, "verbose", false, "Exibe logs detalhados")

	flag.IntVar(&rate, "rate", 0, "Requests por segundo")
	flag.DurationVar(&delay, "delay", 0, "Delay entre requests")
	flag.DurationVar(&timeout, "timeout", 8*time.Second, "Timeout HTTP")

	flag.Parse()

	// conecta verbose ao core
	analysis.Verbose = verbose

	banner()

	// ===============================
	// Validação básica
	// ===============================
	if target == "" && listFile == "" {
		fmt.Println("❌ Use -t ou -list")
		return
	}

	// ===============================
	// Carrega fingerprints
	// ===============================
	if err := analysis.LoadFingerprints("fingerprints"); err != nil {
		fmt.Println("❌ Erro ao carregar fingerprints:", err)
		return
	}

	// ===============================
	// Modos
	// ===============================
	switch {
	case passiveLite:
		fmt.Println("🟢 Modo: PASSIVE-LITE (takeover only)")
		runPassiveLite(target, listFile, timeout)

	case passive:
		fmt.Println("🟡 Modo: PASSIVE (recon + takeover)")
		runPassive(target)

	case active:
		fmt.Println("🟠 Modo: ACTIVE")
		fmt.Println("⚠️ Scan ativo ainda não implementado")

	case full:
		fmt.Println("🔴 Modo: FULL")
		fmt.Println("⚠️ Scan completo ainda não implementado")

	default:
		fmt.Println("ℹ️ Nenhum modo selecionado")
	}
}

// =======================================================
// PASSIVE-LITE
// =======================================================
func runPassiveLite(target, listFile string, timeout time.Duration) {
	var input string

	if listFile != "" {
		input = listFile
		fmt.Println("📄 Usando lista:", listFile)
	} else {
		_ = os.MkdirAll("outputs", 0755)
		input = "outputs/passive_lite_urls.txt"
		_ = os.WriteFile(input, []byte("http://"+target+"\n"), 0644)
		fmt.Println("📄 Usando target único:", target)
	}

	findings, err := analysis.RunTakeover(input, timeout)
	if err != nil {
		fmt.Println("❌ Erro:", err)
		return
	}

	if len(findings) == 0 {
		fmt.Println("✅ Nenhum takeover encontrado")
		return
	}

	fmt.Println("🔥 POSSÍVEIS TAKEOVERS:")
	for _, f := range findings {
		fmt.Printf(" - %s (%s)\n", f.Host, f.Provider)
	}
}

// =======================================================
// PASSIVE
// =======================================================
func runPassive(target string) {
	fmt.Println("🔍 Iniciando reconhecimento passivo")

	if err := recon.Run(target); err != nil {
		fmt.Println("❌ Erro no recon:", err)
		return
	}

	fmt.Println("✅ Reconhecimento passivo finalizado")
}
