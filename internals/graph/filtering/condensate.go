package filtering

import (
	"edaa/internals/graph"
	"edaa/internals/interfaces"
)

func Condensate(g *graph.Graph) {
	for _, node := range g.Nodes() {
		if len(node.InEdges()) > 2 || len(node.OutEdges()) > 2 {
			continue
		}
		connectedNodes := make(map[string][]interfaces.Edge)
		for _, edge := range node.InEdges() {
			connectedNodes[edge.To().Id()] = append(connectedNodes[edge.To().Id()], edge)
		}
		for _, edge := range node.OutEdges() {
			connectedNodes[edge.To().Id()] = append(connectedNodes[edge.To().Id()], edge)
		}
	}
}
