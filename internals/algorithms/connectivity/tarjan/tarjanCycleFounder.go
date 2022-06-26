package tarjan

import (
	"edaa/internals/interfaces"
	"fmt"
	"math"
)

//https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm

var index = 0
var counter int64

var S []interfaces.Node
var lowestLink map[string]int
var discoveredIndex map[string]int
var onStack map[string]bool

var unexplored map[string]bool

var components []interfaces.SCC
var disconnectedComponents [][]interfaces.SCC

// prints the size of each StronglyConnectedComponentsSizes
func PrintStronglyConnectedComponentsSizes(number int64, disconnectedComponents [][]interfaces.SCC) {
	fmt.Println()
	fmt.Println("Stongly connected components:", number)
	for _, c := range disconnectedComponents {
		for _, component := range c {
			fmt.Println("\tFirst node on the component:", component.Nodes()[0].Id())
			fmt.Println("\tComponent size:", len(component.Nodes()))
			fmt.Println("")
		}
	}
}

func TarjanGetStronglyConnectedComponents(g interfaces.Graph) ([][]interfaces.SCC, int64) {
	index = 0
	counter = 0

	S = make([]interfaces.Node, 0)
	lowestLink = make(map[string]int)
	discoveredIndex = make(map[string]int)
	onStack = make(map[string]bool)

	unexplored = make(map[string]bool)

	components = make([]interfaces.SCC, 0)
	disconnectedComponents = make([][]interfaces.SCC, 0)

	for _, n := range g.Nodes() {
		onStack[n.Id()] = false
	}
	for _, n := range g.Nodes() {
		discoveredIndex[n.Id()] = -1
	}
	for _, n := range g.Nodes() {
		unexplored[n.Id()] = true
	}
	for len(unexplored) > 0 {
		var nodeToExplore interfaces.Node
		for k := range unexplored {
			nodeToExplore, _ = g.GetNode(k)
			break
		}
		strongConnect(g, nodeToExplore)
		disconnectedComponents = append(disconnectedComponents, components)
		/*for _, component := range components {
			print(len(component), " ")
		}*/

		components = make([]interfaces.SCC, 0)
		counter++
	}
	return disconnectedComponents, counter
}

func strongConnect(g interfaces.Graph, node interfaces.Node) {
	delete(unexplored, node.Id())
	lowestLink[node.Id()] = index
	discoveredIndex[node.Id()] = index
	index++
	S = append(S, node)
	onStack[node.Id()] = true

	for _, e := range node.OutEdges() {
		if discoveredIndex[e.To().Id()] == -1 {
			strongConnect(g, e.To())
			lowestLink[node.Id()] = int(math.Min(float64(lowestLink[node.Id()]), float64(lowestLink[e.To().Id()])))
		} else if onStack[e.To().Id()] {
			lowestLink[node.Id()] = int(math.Min(float64(lowestLink[node.Id()]), float64(discoveredIndex[e.To().Id()])))
		}
	}
	if lowestLink[node.Id()] == discoveredIndex[node.Id()] {
		var w interfaces.Node
		var nodes []interfaces.Node
		hasStation := false
		for w == nil || w.Id() != node.Id() {
			w = S[len(S)-1]
			S = S[:len(S)-1]
			onStack[w.Id()] = false
			nodes = append(nodes, w)
			if w.IsStation() {
				hasStation = true
			}
		}
		components = append(components, NewSCC(nodes, hasStation))
	}
}
