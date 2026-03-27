package recon

// 🔍 Descoberta simples (mock inicial)
func Discover(domain string) Asset {

	// ⚠️ depois você pluga subfinder aqui
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
