package filtering

import "edaa/internals/interfaces"

func ConnectGraphs(g interfaces.Graph, tree interfaces.NeighbourFinder) {
	println(len(g.WalkableNodes()))

	println(len(g.MetroableNodes()))
	for _, node := range g.MetroableNodes() {
		closestNode, closestDistance := tree.GetClosest(node)
		//closestNode, closestDistance := g.GetClosestNode(node)

		// todo change walking speed
		// m/s
		const walkingSpeed = 4.0
		node.AddDestination(closestNode, closestDistance/walkingSpeed)
		closestNode.AddDestination(node, closestDistance/walkingSpeed)
	}

	println(len(g.BusableNodes()))
	for _, node := range g.BusableNodes() {
		closestNode, closestDistance := tree.GetClosest(node)
		//closestNode, closestDistance := g.GetClosestNode(node)

		// todo change walking speed
		// m/s
		const walkingSpeed = 4.0
		node.AddDestination(closestNode, closestDistance/walkingSpeed)
		closestNode.AddDestination(node, closestDistance/walkingSpeed)
	}
}
