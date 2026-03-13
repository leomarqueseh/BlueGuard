package analysis

func (g *AssetGraph) AddPath(host, path string) {
	h, exists := g.Hosts[host]
	if !exists {
		h = &HostNode{
			Host:  host,
			Paths: make(map[string]*PathNode),
		}
		g.Hosts[host] = h
	}

	if _, exists := h.Paths[path]; !exists {
		h.Paths[path] = &PathNode{
			Path: path,
		}
	}
}
