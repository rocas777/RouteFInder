package path

import (
	"edaa/internals/graph"
	"edaa/internals/interfaces"
	//"math"
)


func a_star(g *graph.Graph, start, finish interfaces.Node) {

	var fCost float64
	var gCost float64
	var hCost float64
	var dist = make(map[interfaces.Node]float64)
	var fCosts = make(map[interfaces.Node]float64)
	var parent = make(map[interfaces.Node]float64)
	var edges []interfaces.Edge
	var currentNode, to interfaces.Node
	

	queue := make(PriorityQueue, len(g.Nodes()))
	queue.Push(start)
	dist[start] = 0

	// while queue is not empty
	for queue.Len() > 0 {

		// getting the current edge and vertex
		currentNode = queue.Pop()
		queue.Pop() // dequeue
		edges = currentNode.OutEdges()

		// terminating search when reaching goal
		if currentNode == finish { break }

		// iterating through all the vertices that we can reach from current vertex
		for i := 0; i < len(edges); i++ {

			// getting the next connected vertex
			to = edges[i].To()
			// calculating gCost
			gCost = edges[i].Weight()

			// relaxing the edge if needed
			if gCost < dist[to] {
				// updating distance if shorter path is found
				dist[to] = gCost
				// setting heuristic in case of A*
				//hCost = heuristic(to, finish)
				// calculating fCost (hCost is zero in case of dijkstra)
				fCost = gCost // + hCost
				// enqueue
				queue.Push(to)
				fCosts[to] = fCost
				// memorizing the parent vertex for building path
				parent[to] = currentNode
			}
		}
	}


}