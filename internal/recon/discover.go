package recon

// 🔥 Discover central (orquestrador)
// Junta resultados de múltiplos providers (assetfinder, subfinder, etc)
func Discover(target string, providers []Provider) Asset {

	target = cleanDomain(target)

	asset := Asset{
		Domain: target,
	}

	seen := make(map[string]bool)

	// 🔥 executa todos providers
	for _, p := range providers {

		subs, err := p.Run(target)
		if err != nil {
			continue
		}

		for _, s := range subs {

			if s == "" {
				continue
			}

			if !seen[s] {
				seen[s] = true
				asset.Subdomains = append(asset.Subdomains, s)
			}
		}
	}

	return asset
}
