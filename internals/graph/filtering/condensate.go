package filtering

import (
	"edaa/internals/graph"
	"edaa/internals/interfaces"
)

func Condensate(g *graph.Graph) {
	counter := 0
	var nodesToRemove []interfaces.Node
	for _, node := range g.Nodes() {
		if len(node.InEdges()) > 2 || len(node.OutEdges()) > 2 {
			continue
		}
		connectedNodes := make(map[string][]interfaces.Edge)
		for _, edge := range node.InEdges() {
			connectedNodes[edge.From().Id()] = append(connectedNodes[edge.From().Id()], edge)
		}
		for _, edge := range node.OutEdges() {
			connectedNodes[edge.To().Id()] = append(connectedNodes[edge.To().Id()], edge)
		}
		if len(connectedNodes) == 2 {
			nodesToRemove = append(nodesToRemove, node)
			counter++
			for _, inEdge := range node.InEdges() {
				for _, outEdge := range node.OutEdges() {
					if inEdge.From().Id() != outEdge.To().Id() {
						weight := inEdge.Weight() + outEdge.Weight()
						inEdge.From().AddDestination(outEdge.To(), weight)
					}
				}
			}
			for _, inEdge := range node.InEdges() {
				inEdge.From().RemoveConnections(node)
			}
			for _, outEdge := range node.OutEdges() {
				outEdge.To().RemoveConnections(node)
			}
		}
	}
	g.RemoveNodes(nodesToRemove)
	println("number of removed nodes", counter)
	println("new ammount of nodes", len(g.Nodes()))
}
