package analysis

// AssetGraph representa a superfície de ataque descoberta
type AssetGraph struct {
	Hosts map[string]*HostNode
}

// HostNode representa um host (domínio / IP)
type HostNode struct {
	Host  string
	Paths map[string]*PathNode
}

// PathNode representa um endpoint
type PathNode struct {
	Path        string
	Redirectable bool
	Params      []string
}

