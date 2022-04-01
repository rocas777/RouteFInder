package main

import (
	"edaa/internals/algorithms/connectivity"
	"edaa/internals/dataStructures/kdtree"
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

	startK := time.Now()
	//for i := 0; i < 10; i++ {
	tree := kdtree.NewKDTree(&g)
	//}
	elapsedK := time.Since(startK)
	fmt.Printf("KD tree %s\n", elapsedK)

	var station *graph.Node
	for _, n := range g.BusableNodes {
		station = n
		break
	}
	startK = time.Now()

	closest, _ := tree.GetClosest(station)
	println("Node, CLosest:", station.Latitude, station.Longitude, closest.Lat(), closest.Lon(), utils.GetDistance(station.Latitude, station.Longitude, closest.Lat(), closest.Lon()))
	elapsedK = time.Since(startK)
	fmt.Printf("KD tree search %s\n", elapsedK)

	start := time.Now()
	g.ConnectGraphs(tree)
	elapsed := time.Since(start)
	utils.PrintMemUsage()
	fmt.Printf("Connecting graphs %s\n", elapsed)

	disconnectedComponents, number := connectivity.TarjanGetStronglyConnectedComponents(&g)
	printStronglyConnectedComponentsSizes(number, disconnectedComponents)

	asd := g.NodesMap["bus_CMPO4"]
	fmt.Printf("%v\n", asd)
	fmt.Printf("%v\n%v\n", g.NodesMap["bus_CMPO4"].IncomingEdges[0].From, g.NodesMap["bus_CMPO4"].Edges[0].To)
	fmt.Printf("%v\n%v\n", g.NodesMap["bus_CMPO4"].IncomingEdges[0].From.Code, g.NodesMap["bus_CMPO4"].Edges[0].To.Code)

	g.ExportOsm("exports/graph.osm")
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
