package main

import (
	"edaa/internals/algorithms/path/astar"
	"edaa/internals/graph"
	"edaa/internals/interfaces"
	"edaa/internals/utils"
	"fmt"
)

func main() {
	utils.PrintMemUsage()

	g := graph.NewGraph()
	buildTestGraph(g)

	as := astar.NewAstar(g, func(from interfaces.Node, to interfaces.Node) float64 {
		return utils.GetDistance(from.Latitude(), from.Longitude(), to.Latitude(), to.Longitude()) / (20)
	})

	path, time, explored := as.Path(g.NodesMap()["node0"], g.NodesMap()["node15"])

	for _, vertex := range path {
		fmt.Printf("%s -> %s  %s  %f\n", vertex.From().Id(), vertex.To().Id(), vertex.EdgeType(), vertex.Weight())
	}

	fmt.Println("time:", time)
	fmt.Println("explored nodes:", explored)
}

func buildTestGraph(g *graph.Graph) {
	node0 := graph.NewNormalNode(0, 0, "node0,0", "zone0", "node0")
	node1 := graph.NewNormalNode(0, 1, "node0,1", "zone0", "node1")
	node2 := graph.NewNormalNode(0, 2, "node0,2", "zone0", "node2")
	node3 := graph.NewNormalNode(0, 3, "node0,3", "zone0", "node3")
	node4 := graph.NewNormalNode(1, 0, "node1,0", "zone0", "node4")
	node5 := graph.NewNormalNode(1, 1, "node1,1", "zone0", "node5")
	node6 := graph.NewNormalNode(1, 2, "node1,2", "zone0", "node6")
	node7 := graph.NewNormalNode(1, 3, "node1,3", "zone0", "node7")
	node8 := graph.NewNormalNode(2, 0, "node2,0", "zone0", "node8")
	node9 := graph.NewNormalNode(2, 1, "node2,1", "zone0", "node9")
	node10 := graph.NewNormalNode(2, 2, "node2,2", "zone0", "node10")
	node11 := graph.NewNormalNode(2, 3, "node2,3", "zone0", "node11")
	node12 := graph.NewNormalNode(3, 0, "node3,0", "zone0", "node12")
	node13 := graph.NewNormalNode(3, 1, "node3,1", "zone0", "node13")
	node14 := graph.NewNormalNode(3, 2, "node3,2", "zone0", "node14")
	node15 := graph.NewNormalNode(3, 3, "node3,3", "zone0", "node15")

	node0.AddDestination(node1, 1)
	node0.AddDestination(node4, 1)
	node1.AddDestination(node0, 1)
	node1.AddDestination(node2, 1)
	node1.AddDestination(node5, 1)
	node2.AddDestination(node1, 1)
	node2.AddDestination(node3, 1)
	node2.AddDestination(node6, 0.5)
	node3.AddDestination(node2, 1)
	node3.AddDestination(node7, 3)
	node4.AddDestination(node0, 1)
	node4.AddDestination(node5, 2)
	node4.AddDestination(node8, 1)
	node5.AddDestination(node1, 1)
	node5.AddDestination(node4, 2)
	node5.AddDestination(node6, 1)
	node5.AddDestination(node9, 2)
	node6.AddDestination(node2, 0.5)
	node6.AddDestination(node5, 1)
	node6.AddDestination(node7, 2)
	node6.AddDestination(node10, 3)
	node7.AddDestination(node3, 3)
	node7.AddDestination(node6, 2)
	node7.AddDestination(node11, 2)
	node8.AddDestination(node4, 1)
	node8.AddDestination(node9, 1)
	node8.AddDestination(node12, 1)
	node9.AddDestination(node5, 2)
	node9.AddDestination(node8, 1)
	node9.AddDestination(node10, 2)
	node9.AddDestination(node13, 1)
	node10.AddDestination(node6, 3)
	node10.AddDestination(node9, 2)
	node10.AddDestination(node11, 1)
	node10.AddDestination(node14, 0.5)
	node11.AddDestination(node7, 2)
	node11.AddDestination(node10, 1)
	node11.AddDestination(node15, 3)
	node12.AddDestination(node8, 1)
	node12.AddDestination(node13, 2)
	node13.AddDestination(node9, 1)
	node13.AddDestination(node12, 2)
	node13.AddDestination(node14, 1)
	node14.AddDestination(node10, 0.5)
	node14.AddDestination(node13, 1)
	node14.AddDestination(node15, 2)
	node15.AddDestination(node11, 3)
	node15.AddDestination(node14, 2)

	g.AddNode(node0)
	g.AddNode(node1)
	g.AddNode(node2)
	g.AddNode(node3)
	g.AddNode(node4)
	g.AddNode(node5)
	g.AddNode(node6)
	g.AddNode(node7)
	g.AddNode(node8)
	g.AddNode(node9)
	g.AddNode(node10)
	g.AddNode(node11)
	g.AddNode(node12)
	g.AddNode(node13)
	g.AddNode(node14)
	g.AddNode(node15)
}
