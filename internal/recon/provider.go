package recon

// 🔥 interface padrão
type Provider interface {
	Run(domain string) ([]string, error)
	Name() string
}
