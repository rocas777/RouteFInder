package main

import (
	"edaa/internals/algorithms/connectivity/tarjan"
	"edaa/internals/dataStructures/kdtree"
	"edaa/internals/graph"
	"edaa/internals/graph/filtering"
	"edaa/internals/interfaces"
)

func main() {
	g := graph.Graph{}
	g.Init()
	cleanGraph(&g)

	tree := kdtree.NewKDTree(&g)
	filtering.ConnectGraphs(&g, tree)

	disconnectedComponents, number := tarjan.TarjanGetStronglyConnectedComponents(&g)
	printStronglyConnectedComponentsSizes(number, disconnectedComponents)

	filtering.Condensate(&g)
}

func cleanGraph(g *graph.Graph) { // filter isolated nodes
	filtering.FilterNodes(g)

	// calculate connectivity
	disconnectedComponents, _ := tarjan.TarjanGetStronglyConnectedComponents(g)

	// remove nodes that are in an isolated strongly connected component
	filtering.RemoveUnconnectedNodes(g, disconnectedComponents)
}

// prints the size of each StronglyConnectedComponentsSizes
func printStronglyConnectedComponentsSizes(number int64, disconnectedComponents [][]interfaces.SCC) {
	for _, components := range disconnectedComponents {
		for _, component := range components {
			println(component.Nodes()[0].Id(), len(component.Nodes()))
		}
	}
}
