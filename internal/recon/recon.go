package recon

import (
	"fmt"

	"github.com/leomarqueseh/BlueGuard/internal/analysis"
)

func Run(target string) error {
	fmt.Println("🔍 [Recon] Iniciando reconhecimento")
	fmt.Println("🌐 [Recon] Alvo:", target)

	// 1️⃣ Enumeração de subdomínios
	_, err := RunSubdomainEnum(target)
	if err != nil {
    		return err
	}

	// 2️⃣ Verificação de hosts vivos
	alive, err := RunAliveCheck()
	if err != nil {
		return err
	}

	// 3️⃣ 🔥 AQUI É O PONTO EXATO DO ANALYSIS 🔥
	_, err = analysis.RunEndpointCollection(alive)
	if err != nil {
		return err
	}

	fmt.Println("✅ [Recon] Recon finalizado com sucesso")
	return nil
}
