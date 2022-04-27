package filtering

import "edaa/internals/interfaces"

func RemoveBrokenEdges(g interfaces.Graph) {
	for _, node := range g.Nodes() {
		outEdges := make([]interfaces.Edge, 0)
		for _, edge := range node.OutEdges() {
			if _, ok := g.NodesMap()[edge.To().Id()]; ok {
				outEdges = append(outEdges, edge)
			}
		}
		node.SetOutEdges(outEdges)
	}
}
