package recon

// 🔥 BuildGraph transforma um Asset em um grafo (mapa mental)
// Isso será usado no dashboard (Cytoscape)
func BuildGraph(a Asset) Graph {

	var nodes []Node
	var edges []Edge

	// 🔴 Nó principal (domínio raiz)
	nodes = append(nodes, Node{
		ID:    a.Domain,
		Label: a.Domain,
		Type:  "domain",
	})

	// 🔵 Subdomínios
	for _, sub := range a.Subdomains {

		nodes = append(nodes, Node{
			ID:    sub,
			Label: sub,
			Type:  "subdomain",
		})

		// ligação domínio → subdomínio
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
