package menuHelper

import (
	"edaa/internals/algorithms/path/astar"
	"edaa/internals/algorithms/path/genetics"
	"edaa/internals/algorithms/path/landmarks"
	"edaa/internals/dataStructures/kdtree"
	"edaa/internals/graph"
	"edaa/internals/interfaces"
	"edaa/internals/utils"
	"fmt"
)

func PathFinder(g interfaces.Graph, tree *kdtree.KDTree) {
	as := astar.NewAstar(g, func(from interfaces.Node, to interfaces.Node) float64 {
		return utils.GetDistanceBetweenNodes(from, to) / (20)
	})

	startNode := g.Nodes()[7999]
	endNode := g.Nodes()[157000]
	// var slat float64
	// var slon float64
	// var elat float64
	// var elon float64
	// println("Starting Point(Leave empty for random)")
	// fmt.Scanf("%f,%f", &slat, &slon)

	// println("Ending Point(Leave empty for random)")
	// fmt.Scanf("%f,%f", &elat, &elon)

	// if slat != 0 || slon != 0 {
	// 	startNode, _ = tree.GetClosest(graph.NewNormalNode(slat, slon, "", "", ""))
	// }
	// if elat != 0 || elon != 0 {
	// 	endNode, _ = tree.GetClosest(graph.NewNormalNode(elat, elon, "", "", ""))
	// }

	//startNode := g.NodesMap()["metro_27"]
	//endNode := g.NodesMap()["metro_76"]

	path, pathTime, explored := as.Path(startNode, endNode)

	astar.PreetyDisplay(path, pathTime, explored, startNode, endNode)
	astar.ExportEdges(path)
}

func PathFinderGenetics(g interfaces.Graph, tree *kdtree.KDTree) {
	startNode := g.Nodes()[7999]
	endNode := g.Nodes()[157000]
	var slat float64
	var slon float64
	var elat float64
	var elon float64
	println("Starting Point(Leave empty for random)")
	fmt.Scanf("%f,%f", &slat, &slon)

	println("Ending Point(Leave empty for random)")
	fmt.Scanf("%f,%f", &elat, &elon)

	if slat != 0 || slon != 0 {
		startNode, _ = tree.GetClosest(graph.NewNormalNode(slat, slon, "", "", ""))
	}
	if elat != 0 || elon != 0 {
		endNode, _ = tree.GetClosest(graph.NewNormalNode(elat, elon, "", "", ""))
	}

	//startNode := g.NodesMap()["metro_27"]
	//endNode := g.NodesMap()["metro_76"]

	path, pathTime, explored := genetics.GeneticPath(g, startNode, endNode)

	astar.PreetyDisplay(path, pathTime, explored, startNode, endNode)
	astar.ExportEdgesGenetics(path)
}

func PathFinderLandmarks(g interfaces.Graph, tree *kdtree.KDTree, d *landmarks.Dijkstra) {
	startNode := g.Nodes()[7999]
	endNode := g.Nodes()[157000]
	// var slat float64
	// var slon float64
	// var elat float64
	// var elon float64
	// println("Starting Point(Leave empty for random)")
	// fmt.Scanf("%f,%f", &slat, &slon)

	// println("Ending Point(Leave empty for random)")
	// fmt.Scanf("%f,%f", &elat, &elon)

	// if slat != 0 || slon != 0 {
	// 	startNode, _ = tree.GetClosest(graph.NewNormalNode(slat, slon, "", "", ""))
	// }
	// if elat != 0 || elon != 0 {
	// 	endNode, _ = tree.GetClosest(graph.NewNormalNode(elat, elon, "", "", ""))
	// }

	activeLandmarks := d.SelectActiveLandmarks(startNode, endNode)

	as := astar.NewAstar(g, func(from interfaces.Node, to interfaces.Node) float64 {
		return landmarks.Heuristic(from, to, activeLandmarks)
	})

	path, pathTime, explored := as.Path(startNode, endNode)

	astar.PreetyDisplay(path, pathTime, explored, startNode, endNode)
	astar.ExportEdges(path)
}
