package main

import (
	"edaa/internals/algorithms/path/genetics"
	kdtree2 "edaa/internals/dataStructures/kdtree"
	"edaa/internals/exports/reuse"
	"edaa/internals/graph"
	"edaa/internals/interfaces"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	g := &graph.Graph{}
	graph.InitReuse(g)

	//clean(g)

	//walk_6055022105 walk_4758581322

	rand.Seed(time.Now().UnixNano())

	var cycles uint64 = 0
	var bad uint64 = 0
	c := make(chan int)
	kdtree := kdtree2.NewKDTree(g)

	maxGoroutines := 10
	guard := make(chan struct{}, maxGoroutines)

	go func() {
		for ; cycles < 100; cycles++ {
			guard <- struct{}{} // would block if guard channel is already filled
			go func(n uint64) {
				genticIteration(g, c, kdtree)
				<-guard
			}(cycles)
		}
	}()
	for i := 0; i < 100; i++ {
		select {
		case r := <-c:
			fmt.Println(i,r)
			if r == 0 {
				bad++
			}
		}
	}
	fmt.Println("Cycles:", cycles)
	fmt.Println("Bad:", bad)
}

func genticIteration(g *graph.Graph, c chan int, kdtree *kdtree2.KDTree) {
	p1 := rand.Intn(len(g.Nodes()))
	p2 := rand.Intn(len(g.Nodes()))
	_, _, ok := genetics.GeneticPath(g, g.Nodes()[p1], g.Nodes()[p2], kdtree)
	c <- ok
}

func clean(g *graph.Graph) {
	i := math.Inf(1)
	for i != 0 {
		i = 0
		var nodesToRemove []interfaces.Node
		for _, node := range g.Nodes() {
			if len(node.OutEdges()) == 1 && len(node.InEdges()) == 1 {
				i++
				weight := node.OutEdges()[0].Weight() + node.InEdges()[0].Weight()
				node.InEdges()[0].From().RemoveConnections(node)
				node.OutEdges()[0].To().RemoveConnections(node)
				node.InEdges()[0].From().AddDestination(node.OutEdges()[0].To(), weight)
				nodesToRemove = append(nodesToRemove, node)
			}
		}
		g.RemoveNodes(nodesToRemove)
		fmt.Println("Removed Nodes", i)
		fmt.Println(len(g.Nodes()))
	}
	reuse.ExportEdges(g, "data/reuse/edges.csv")
	reuse.ExportNodes(g, "data/reuse/nodes.csv")
}
