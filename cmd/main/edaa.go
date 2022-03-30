package main

import (
	"edaa/internals/algorithms/connectivity"
	"edaa/internals/graph"
)

func main() {
	g := graph.Graph{}
	g.Init()
	components := connectivity.TarjanGetStronglyConnectedComponents(&g)
	for _, component := range components {
		for _, node := range component {
			println(node.Code)
		}
		println()
		println()
		println()
	}
}
