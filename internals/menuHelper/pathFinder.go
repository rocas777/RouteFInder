package menuHelper

import (
	"edaa/internals/algorithms/path/astar"
	"edaa/internals/interfaces"
	"edaa/internals/utils"
)

func PathFinder(g interfaces.Graph) {
	as := astar.NewAstar(g, func(from interfaces.Node, to interfaces.Node) float64 {
		return 0
		return utils.GetDistanceBetweenNodes(from, to) / (33 / 3.6)
	})

	startNode := g.Nodes()[7999]
	endNode := g.Nodes()[157000]
	//startNode := g.NodesMap()["metro_27"]
	//endNode := g.NodesMap()["metro_76"]

	path, pathTime, explored := as.Path(startNode, endNode)

	astar.PreetyDisplay(path, pathTime, explored, startNode, endNode)
	astar.ExportEdges(path)
}
