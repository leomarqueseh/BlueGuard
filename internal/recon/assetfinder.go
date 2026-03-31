package recon

import (
	"os/exec"
	"strings"
)

// 🔥 estrutura
type Assetfinder struct{}

// 🔥 construtor (OBRIGATÓRIO)
func NewAssetfinder() Provider {
	return &Assetfinder{}
}

func (a *Assetfinder) Name() string {
	return "assetfinder"
}

func (a *Assetfinder) Run(domain string) ([]string, error) {

	cmd := exec.Command("assetfinder", "--subs-only", domain)

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(out), "\n")

	var subs []string

	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l != "" {
			subs = append(subs, l)
		}
	}

	return subs, nil
}
