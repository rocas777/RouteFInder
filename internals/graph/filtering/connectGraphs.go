package filtering

import "edaa/internals/interfaces"

const walkingSpeed = 1.0

func ConnectGraphs(g interfaces.Graph, tree interfaces.NeighbourFinder) {
	println("Walkable Network Nodes:", len(g.WalkableNodes()))

	println("Metro Network Nodes:", len(g.MetroableNodes()))
	for _, node := range g.MetroableNodes() {
		closestNode, closestDistance := tree.GetClosest(node)
		//closestNode, closestDistance := g.GetClosestNode(node)

		node.AddDestination(closestNode, closestDistance/walkingSpeed)
		closestNode.AddDestination(node, closestDistance/walkingSpeed)
	}

	println("Bus Network Nodes:", len(g.BusableNodes()))
	for _, node := range g.BusableNodes() {
		closestNode, closestDistance := tree.GetClosest(node)
		//closestNode, closestDistance := g.GetClosestNode(node)

		node.AddDestination(closestNode, closestDistance/walkingSpeed)
		closestNode.AddDestination(node, closestDistance/walkingSpeed)
	}
}
