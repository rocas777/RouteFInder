package main

import (
	"edaa/internals/algorithms/connectivity/tarjan"
	"edaa/internals/graph"
)

func main() {
	g := graph.Graph{}
	graph.InitReuse(&g)

	disconnectedComponents, number := tarjan.TarjanGetStronglyConnectedComponents(&g)
	tarjan.PrintStronglyConnectedComponentsSizes(number, disconnectedComponents)
}
