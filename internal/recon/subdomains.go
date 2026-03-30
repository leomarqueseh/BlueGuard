package recon

// 🔥 Discover subdomínios (simples por enquanto)
func Discover(domain string) Asset {

	// 🔥 limpar domínio (evita bug tipo http://api.example.com)
	domain = cleanDomain(domain)

	subs := []string{
		"www." + domain,
		"api." + domain,
		"dev." + domain,
	}

	return Asset{
		Domain:     domain,
		Subdomains: subs,
	}
}
