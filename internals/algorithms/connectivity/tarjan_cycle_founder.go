package connectivity

import (
	"edaa/internals/graph"
	"math"
)

//https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm

var index = 0
var counter int64

var S []*graph.Node
var lowestLink map[string]int
var discoveredIndex map[string]int
var onStack map[string]bool

var unexplored map[string]bool

var components []*graph.SCC
var disconnectedComponents [][]*graph.SCC

func TarjanGetStronglyConnectedComponents(g *graph.Graph) ([][]*graph.SCC, int64) {
	index = 0
	counter = 0

	S = make([]*graph.Node, 0)
	lowestLink = make(map[string]int)
	discoveredIndex = make(map[string]int)
	onStack = make(map[string]bool)

	unexplored = make(map[string]bool)

	components = make([]*graph.SCC, 0)
	disconnectedComponents = make([][]*graph.SCC, 0)

	for _, n := range g.Nodes {
		onStack[n.Code] = false
	}
	for _, n := range g.Nodes {
		discoveredIndex[n.Code] = -1
	}
	for _, n := range g.Nodes {
		unexplored[n.Code] = true
	}
	for len(unexplored) > 0 {
		var nodeToExplore *graph.Node
		for k := range unexplored {
			nodeToExplore, _ = g.GetNode(k)
			break
		}
		strongConnect(g, nodeToExplore)
		disconnectedComponents = append(disconnectedComponents, components)
		/*for _, component := range components {
			print(len(component), " ")
		}*/

		components = make([]*graph.SCC, 0)
		counter++
	}
	return disconnectedComponents, counter
}

func strongConnect(g *graph.Graph, node *graph.Node) {
	delete(unexplored, node.Code)
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
		var w *graph.Node
		var nodes []*graph.Node
		hasStation := false
		for w == nil || w.Code != node.Code {
			w = S[len(S)-1]
			S = S[:len(S)-1]
			onStack[w.Code] = false
			nodes = append(nodes, w)
			if w.IsStation {
				hasStation = true
			}
		}
		components = append(components, graph.NewSCC(nodes, hasStation))
	}
}
