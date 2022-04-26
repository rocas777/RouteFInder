package main

import (
	"edaa/internals/graph"
	"edaa/internals/utils"
	"edaa/internals/algorithms/distance"
	"fmt"
	"time"
)

func main() {
	g := graph.Graph{}
	g.Init()

	// calculate distance matrix floyd-warshall
	g.MetroableNodes()

	start := time.Now()
	distance.FloydWarshall(&g)
	elapsed := time.Since(start)
	utils.PrintMemUsage()
	fmt.Printf("Distance matrix %s\n", elapsed)

}

	
	
	
	
	
