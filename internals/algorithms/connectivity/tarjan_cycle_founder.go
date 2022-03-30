package connectivity

import (
	"edaa/internals/graph"
	"math"
)

//https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm

var index = 0
var S []*graph.Node
var lowestLink map[string]int
var discoveredIndex map[string]int
var onStack map[string]bool

var components [][]*graph.Node

func TarjanGetStronglyConnectedComponents(g *graph.Graph) [][]*graph.Node {
	lowestLink = make(map[string]int)
	discoveredIndex = make(map[string]int)
	onStack = make(map[string]bool)

	for _, n := range g.Nodes {
		onStack[n.Code] = false
	}
	for _, n := range g.Nodes {
		discoveredIndex[n.Code] = -1
	}
	strongConnect(g, g.Nodes[0])
	return components
}

func strongConnect(g *graph.Graph, node *graph.Node) {
	lowestLink[node.Code] = index
	discoveredIndex[node.Code] = index
	index++
	S = append(S, node)
	onStack[node.Code] = true

	for _, e := range node.Edges {
		if discoveredIndex[e.To.Code] == -1 {
			strongConnect(g, e.To)
			lowestLink[node.Code] = int(math.Min(float64(lowestLink[node.Code]), float64(lowestLink[e.To.Code])))
		} else if onStack[e.To.Code] {
			lowestLink[node.Code] = int(math.Min(float64(lowestLink[node.Code]), float64(discoveredIndex[e.To.Code])))
		}
	}
	if lowestLink[node.Code] == discoveredIndex[node.Code] {
		var w *graph.Node = nil
		components = append(components, []*graph.Node{})
		for w == nil || w.Code != node.Code {
			w = S[len(S)-1]
			S = S[:len(S)-1]
			onStack[w.Code] = false
			components[len(components)-1] = append(components[len(components)-1], w)
		}
	}
}
