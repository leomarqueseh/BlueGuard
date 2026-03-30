package recon

// 🔥 Asset representa o alvo analisado
type Asset struct {
	Domain     string   // domínio principal
	Subdomains []string // subdomínios encontrados
	IPs        []string // IPs (opcional)
}

// 🔥 Node do grafo (Cytoscape)
type Node struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	Type       string `json:"type"`
	Parent     string `json:"parent,omitempty"`
	Expandable bool   `json:"expandable"`
}

// 🔥 Edge (ligação entre nós)
type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

// 🔥 Graph completo enviado para o frontend
type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}
