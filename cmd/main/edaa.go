package main

import (
	"edaa/internals/algorithms/connectivity"
	"edaa/internals/graph"
)

func main() {
	g := graph.Graph{}
	g.Init()
	println(len(connectivity.TarjanGetStronglyConnectedComponents(&g)))
	println(g.GetCoordsBox())

	g.ExportNodes("exports/nodes.csv")
	g.ExportEdges("exports/edges.csv")

}
