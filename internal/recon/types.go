package recon

// 🔥 Estrutura principal
type Asset struct {
	Domain     string
	Subdomains []string
}

// 🔥 Grafo (mapa mental)
type Node struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}
