package main

import (
	kdtree "edaa/internals/dataStructures/quadtree"
	"edaa/internals/graph"
	tile_server "edaa/internals/visualization/tile-server"
)

func main() {

	g := &graph.Graph{}
	graph.InitReuse(g)

	quadtree := kdtree.NewQuadTree(g)
	tile_server.TileServer(quadtree.Root,nil)
}
