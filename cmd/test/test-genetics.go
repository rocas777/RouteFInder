package main

import (
	"edaa/internals/algorithms/path/genetics"
	"edaa/internals/exports/reuse"
	"edaa/internals/graph"
	"edaa/internals/interfaces"
	"math"
)

func main() {
	g := &graph.Graph{}
	graph.InitReuse(g)

	//clean(g)

	//walk_6055022105 walk_4758581322
	println(g.Nodes()[0].Id(), g.Nodes()[10000].Id())
	genetics.GeneticPath(g, g.Nodes()[0], g.Nodes()[1000])
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
		println("Removed Nodes", i)
		println(len(g.Nodes()))
	}
	reuse.ExportEdges(g, "data/reuse/edges.csv")
	reuse.ExportNodes(g, "data/reuse/nodes.csv")
}
