package menuHelper

import (
	"edaa/internals/algorithms/connectivity/tarjan"
	"edaa/internals/dataStructures/kdtree"
	"edaa/internals/exports/reuse"
	"edaa/internals/graph"
	"edaa/internals/graph/filtering"
	"edaa/internals/interfaces"
)

func Setup() interfaces.Graph {
	println("")
	println("Starting Setup")

	g := graph.Graph{}
	println("")
	println("Initiating graph...")
	g.Init()

	//lat1, lat2, lon1, lon2 := g.GetCoordsBox()
	//filtering.Crop(&g, lat1+((lat2-lat1)/2), lon1+((lon2-lon1)/2), lat2, lon2)

	println("")
	println("Cleaning graph from isolated nodes...")
	cleanGraph(&g)

	println("")
	println("Connecting Stations to road Network...")
	tree := kdtree.NewKDTree(&g)
	filtering.ConnectGraphs(&g, tree)

	//disconnectedComponents, number := tarjan.TarjanGetStronglyConnectedComponents(&g)
	//tarjan.PrintStronglyConnectedComponentsSizes(number, disconnectedComponents)

	//filtering.Condensate(&g)

	println("")
	println("Removing Broken edges...")
	filtering.RemoveBrokenEdges(&g)

	println("")
	println("Exporting data to reuse in the future...")
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
