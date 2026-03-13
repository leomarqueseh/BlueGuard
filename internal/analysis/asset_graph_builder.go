package analysis

func NewAssetGraph() *AssetGraph {
	return &AssetGraph{
		Hosts: make(map[string]*HostNode),
	}
}
