package filtering

import (
	"edaa/internals/interfaces"
)

func FilterNodes(g interfaces.Graph) {
	nodesCopy := g.Nodes()
	g.SetNodes(make([]interfaces.Node, 0))
	g.SetNodesMap(make(map[string]interfaces.Node))
	for _, node := range nodesCopy {
		if node.Referenced() {
			g.AddNode(node)
		}
	}
}
