package recon

func BuildGraph(a Asset) Graph {

	var nodes []Node
	var edges []Edge

	// 🔴 domínio raiz
	nodes = append(nodes, Node{
		ID:         a.Domain,
		Label:      a.Domain,
		Type:       "domain",
		Expandable: true,
	})

	// 🔵 subdomínios
	for _, sub := range a.Subdomains {

		nodes = append(nodes, Node{
			ID:         sub,
			Label:      sub,
			Type:       "subdomain",
			Expandable: true,
		})

		edges = append(edges, Edge{
			Source: a.Domain,
			Target: sub,
		})
	}

	return Graph{
		Nodes: nodes,
		Edges: edges,
	}
}
