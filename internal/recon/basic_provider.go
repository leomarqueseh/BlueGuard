package recon

// 🔥 Provider básico (fallback interno)
// Garante que sempre teremos algo no grafo
type BasicProvider struct{}

func (p *BasicProvider) Name() string {
	return "basic"
}

func (p *BasicProvider) Run(domain string) ([]string, error) {

	// 🔥 fallback simples (evita tela vazia)
	return []string{
		"www." + domain,
		"api." + domain,
		"dev." + domain,
	}, nil
}
