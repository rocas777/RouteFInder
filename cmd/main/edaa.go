package main

import (
	"edaa/internals/algorithms/connectivity"
	"edaa/internals/graph"
	"edaa/internals/utils"
	"fmt"
	"sort"
	"time"
)

func main() {
	utils.PrintMemUsage()

	g := graph.Graph{}
	g.Init()
	cleanGraph(&g)

	start := time.Now()
	g.ConnectGraphs()
	elapsed := time.Since(start)
	utils.PrintMemUsage()
	fmt.Printf("Connecting graphs %s\n", elapsed)

	disconnectedComponents, number := connectivity.TarjanGetStronglyConnectedComponents(&g)
	printStronglyConnectedComponentsSizes(number, disconnectedComponents)

	g.ExportNodes("exports/nodes.csv")
	g.ExportEdges("exports/edges.csv")
}

func cleanGraph(g *graph.Graph) { // filter isolated nodes
	g.FilterNodes()
	println("Filtered nodes:", len(g.Nodes))
	utils.PrintMemUsage()

	// calculate connectivity
	start := time.Now()
	disconnectedComponents, number := connectivity.TarjanGetStronglyConnectedComponents(g)
	elapsed := time.Since(start)
	utils.PrintMemUsage()
	fmt.Printf("Connectivity %s\n", elapsed)
	printStronglyConnectedComponentsSizes(number, disconnectedComponents)

	// remove nodes that are in an isolated strongly connected component
	g.RemoveUnconnectedNodes(disconnectedComponents)
	println("Removed Unconnected nodes:", len(g.Nodes))

	// recalculate connectivity
	start = time.Now()
	disconnectedComponents, number = connectivity.TarjanGetStronglyConnectedComponents(g)
	elapsed = time.Since(start)
	utils.PrintMemUsage()
	fmt.Printf("Connectivity %s\n", elapsed)
	printStronglyConnectedComponentsSizes(number, disconnectedComponents)

	utils.PrintMemUsage()

	fmt.Printf("Showing components %s\n", elapsed)
}

// prints the size of each StronglyConnectedComponentsSizes
func printStronglyConnectedComponentsSizes(number int64, disconnectedComponents [][]*graph.SCC) {
	var componentSizes []int
	println("Number of disconnected Componnets:", number)
	for _, components := range disconnectedComponents {
		for _, component := range components {
			componentSizes = append(componentSizes, len(component.Nodes))
			if len(component.Nodes) == 1 {
				println(component.Nodes[0].Code)
			}
		}
	}
	sort.Ints(componentSizes)
	for _, componentSize := range componentSizes {
		print(componentSize, " ")
	}
	println()
}
