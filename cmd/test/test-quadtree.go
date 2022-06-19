package main

import (
	kdtree "edaa/internals/dataStructures/quadtree"
	"edaa/internals/graph"
)

func main() {

	g := &graph.Graph{}
	graph.InitReuse(g)

	kdtree.NewQuadTree(g)

}
