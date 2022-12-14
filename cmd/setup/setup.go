package main

import (
	"edaa/internals/algorithms/connectivity/tarjan"
	"edaa/internals/dataStructures/kdtree"
	"edaa/internals/exports/reuse"
	"edaa/internals/graph"
	"edaa/internals/graph/filtering"
)

func main() {
	g := graph.Graph{}
	g.Init()

	//lat1, lat2, lon1, lon2 := g.GetCoordsBox()
	//filtering.Crop(&g, lat1+((lat2-lat1)/2), lon1+((lon2-lon1)/2), lat2, lon2)

	cleanGraph(&g)

	tree := kdtree.NewKDTree(&g)
	filtering.ConnectGraphs(&g, tree)

	disconnectedComponents, number := tarjan.TarjanGetStronglyConnectedComponents(&g)
	tarjan.PrintStronglyConnectedComponentsSizes(number, disconnectedComponents)

	//filtering.Condensate(&g)

	filtering.RemoveBrokenEdges(&g)

	reuse.ExportEdges(&g, "data/reuse/edges.csv")
	reuse.ExportNodes(&g, "data/reuse/nodes.csv")
}

func cleanGraph(g *graph.Graph) { // filter isolated nodes
	filtering.FilterNodes(g)

	// calculate connectivity
	disconnectedComponents, _ := tarjan.TarjanGetStronglyConnectedComponents(g)

	// remove nodes that are in an isolated strongly connected component
	filtering.RemoveUnconnectedNodes(g, disconnectedComponents)
}
