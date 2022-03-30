package main

import (
	"edaa/internals/algorithms/connectivity"
	"edaa/internals/graph"
)

func main() {
	g := graph.Graph{}
	g.Init()
	println(len(connectivity.TarjanGetStronglyConnectedComponents(&g)))
}
