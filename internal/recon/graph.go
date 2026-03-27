package recon

import (
	"net"
)

func BuildGraph(a Asset) Graph {

	var nodes []Node
	var edges []Edge

	// 🔴 domínio principal
	nodes = append(nodes, Node{
		ID:    a.Domain,
		Label: a.Domain,
		Type:  "domain",
	})

	for _, sub := range a.Subdomains {

		// 🔵 subdomínio
		nodes = append(nodes, Node{
			ID:    sub,
			Label: sub,
			Type:  "subdomain",
		})

		edges = append(edges, Edge{
			Source: a.Domain,
			Target: sub,
		})

		// 🔥 resolver IP
		ips, _ := net.LookupHost(sub)

		for _, ip := range ips {

			ipID := sub + "_ip_" + ip

			nodes = append(nodes, Node{
				ID:    ipID,
				Label: ip,
				Type:  "ip",
			})

			edges = append(edges, Edge{
				Source: sub,
				Target: ipID,
			})

			// 🔥 NMAP
			ports := ScanPorts(ip)

			for _, p := range ports {

				portID := ipID + "_port_" + p.Port

				nodes = append(nodes, Node{
					ID:    portID,
					Label: p.Port + "/" + p.Service,
					Type:  "port",
				})

				edges = append(edges, Edge{
					Source: ipID,
					Target: portID,
				})

				// 🔥 versão (service details)
				if p.Version != "" {

					verID := portID + "_ver"

					nodes = append(nodes, Node{
						ID:    verID,
						Label: p.Version,
						Type:  "version",
					})

					edges = append(edges, Edge{
						Source: portID,
						Target: verID,
					})
				}
			}
		}
	}

	return Graph{
		Nodes: nodes,
		Edges: edges,
	}
}
