package recon

import (
	"os/exec"
	"strings"
)

// 🔥 estrutura
type Subfinder struct{}

// 🔥 construtor (OBRIGATÓRIO para funcionar no dashboard)
func NewSubfinder() Provider {
	return &Subfinder{}
}

func (s *Subfinder) Name() string {
	return "subfinder"
}

func (s *Subfinder) Run(domain string) ([]string, error) {

	cmd := exec.Command("subfinder", "-silent", "-d", domain)

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
