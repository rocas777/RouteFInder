package menuHelper

import (
	"edaa/internals/algorithms/connectivity/tarjan"
	"edaa/internals/dataStructures/kdtree"
	"edaa/internals/exports/reuse"
	"edaa/internals/graph"
	"edaa/internals/graph/filtering"
	"edaa/internals/interfaces"
	"fmt"
)

func Setup() interfaces.Graph {
	fmt.Println("")
	fmt.Println("Starting Setup")

	g := graph.Graph{}
	fmt.Println("")
	fmt.Println("Initiating graph...")
	g.Init()

	//lat1, lat2, lon1, lon2 := g.GetCoordsBox()
	//filtering.Crop(&g, lat1+((lat2-lat1)/2), lon1+((lon2-lon1)/2), lat2, lon2)

	fmt.Println("")
	fmt.Println("Cleaning graph from isolated nodes...")
	cleanGraph(&g)

	fmt.Println("")
	fmt.Println("Connecting Stations to road Network...")
	tree := kdtree.NewKDTree(&g)
	filtering.ConnectGraphs(&g, tree)

	//disconnectedComponents, number := tarjan.TarjanGetStronglyConnectedComponents(&g)
	//tarjan.PrintStronglyConnectedComponentsSizes(number, disconnectedComponents)

	//filtering.Condensate(&g)

	fmt.Println("")
	fmt.Println("Removing Broken edges...")
	filtering.RemoveBrokenEdges(&g)

	fmt.Println("")
	fmt.Println("Exporting data to reuse in the future...")
	reuse.ExportEdges(&g, "data/reuse/edges.csv")
	reuse.ExportNodes(&g, "data/reuse/nodes.csv")

	return &g
}

func cleanGraph(g *graph.Graph) { // filter isolated nodes
	filtering.FilterNodes(g)

	// calculate connectivity
	disconnectedComponents, _ := tarjan.TarjanGetStronglyConnectedComponents(g)

	// remove nodes that are in an isolated strongly connected component
	filtering.RemoveUnconnectedNodes(g, disconnectedComponents)
}
