package astar

import (
	"edaa/internals/interfaces"
	"edaa/internals/types"
	fibHeap "github.com/starwander/GoFibonacciHeap"
	"math"
)

type astarNode struct {
	last            *astarNode
	lastDistance    float64
	edgeType        types.EdgeType
	accumulatedDist float64
	realNode        interfaces.Node
}

type node struct {
	dist float64
	node *astarNode
}

func (n *node) Tag() interface{} {
	return n.node
}

func (n *node) Key() float64 {
	return n.dist
}

type astar struct {
	heap  *fibHeap.FibHeap
	graph interfaces.Graph
	H     func(from interfaces.Node, to interfaces.Node) float64
}

func (a *astar) Path(source interfaces.Node, destination interfaces.Node) ([]interfaces.Edge, float64, int) {
	explored := make(map[string]interface{})
	astarNodes := make(map[string]*astarNode)

	aN := &astarNode{
		last:            nil,
		accumulatedDist: 0,
		realNode:        source,
	}
	a.heap.Insert(aN, 0)
	astarNodes[source.Id()] = aN

	for len(explored) < len(a.graph.Nodes()) {
		i, v := a.heap.ExtractMin()
		currentHeapNode := node{
			dist: v,
			node: i.(*astarNode),
		}
		currentAStarNode := currentHeapNode.Tag().(*astarNode)
		currentRealNode := currentAStarNode.realNode
		explored[currentRealNode.Id()] = ""
		if currentRealNode.Id() == destination.Id() {
			path, weight := a.fetchPath(currentHeapNode.Tag().(*astarNode))
			//fmt.Println(len(path))
			return path, weight, len(explored)
		}
		for _, edge := range currentRealNode.OutEdges() {
			destinationRealNode := edge.To()
			var destinationAStarNode *astarNode
			var ok bool
			if destinationAStarNode, ok = astarNodes[destinationRealNode.Id()]; !ok {
				astarNodes[destinationRealNode.Id()] = &astarNode{
					last:            nil,
					accumulatedDist: math.Inf(1),
					realNode:        destinationRealNode,
				}
			}
			destinationAStarNode = astarNodes[destinationRealNode.Id()]
			if _, ok := explored[destinationRealNode.Id()]; !ok {
				newDistance := edge.Weight() + currentAStarNode.accumulatedDist + a.H(destination, destinationRealNode)

				if destinationAStarNode.accumulatedDist > edge.Weight()+currentAStarNode.accumulatedDist {
					err := a.heap.Insert(destinationAStarNode, newDistance)
					if err == nil {
						destinationAStarNode.last = currentAStarNode
						destinationAStarNode.lastDistance = edge.Weight()
						destinationAStarNode.edgeType = edge.EdgeType()
						destinationAStarNode.accumulatedDist = edge.Weight() + currentAStarNode.accumulatedDist
					}
				}
			}
		}
	}
	return []interfaces.Edge{}, 0, len(explored)
}

func (a *astar) Dist(source interfaces.Node) []struct {
	node interfaces.Node
	dist float64
} {
	//TODO implement me
	panic("implement me")
}

func NewAstar(graph interfaces.Graph, H func(from interfaces.Node, to interfaces.Node) float64) *astar {
	return &astar{graph: graph, heap: fibHeap.NewFibHeap(), H: H}
}

func (a *astar) fetchPath(last *astarNode) ([]interfaces.Edge, float64) {
	tempPath := make([]interfaces.Edge, 0)
	exploring := last
	for exploring.last != nil {
		tE := astarEdge{
			from:     exploring.last.realNode,
			to:       exploring.realNode,
			edgeType: exploring.edgeType,
			weight:   exploring.lastDistance,
		}
		tempPath = append(tempPath, tE)
		exploring = exploring.last
	}

	outPath := make([]interfaces.Edge, len(tempPath))
	for i := len(tempPath) - 1; i >= 0; i-- {
		outPath[len(tempPath)-i-1] = tempPath[i]
	}
	return outPath, last.accumulatedDist
}
