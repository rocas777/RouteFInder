package landmarks

import (
	"edaa/internals/algorithms/path/astar"
	"edaa/internals/interfaces"
	"edaa/internals/types"
	"math"

	fibHeap "github.com/starwander/GoFibonacciHeap"
)

type dijkstraNode struct {
	last            *dijkstraNode
	lastDistance    float64
	edgeType        types.EdgeType
	accumulatedDist float64
	realNode        interfaces.Node
}

type node struct {
	dist float64
	node *dijkstraNode
}

func (n *node) Tag() interface{} {
	return n.node
}

func (n *node) Key() float64 {
	return n.dist
}

type dijkstra struct {
	heap  *fibHeap.FibHeap
	graph interfaces.Graph
}

func (d *dijkstra) PreprocessLandmark(landmark int, source interfaces.Node) () {
	explored := make(map[string]interface{})
	dijkstraNodes := make(map[string]*dijkstraNode)
	
	// populate heap and map with all nodes with +infinity distance
	for _, n := range d.graph.Nodes() {
		if n.Id() != source.Id() {
			dN := &dijkstraNode{
				last:            nil,
				accumulatedDist: math.Inf(1),
				realNode:        n,
			}
			d.heap.Insert(dN, math.Inf(1))
			dijkstraNodes[n.Id()] = dN
		}
	}

	// insert source node with 0 distance
	dN := &dijkstraNode{
		last:            nil,
		accumulatedDist: 0,
		realNode:        source,
	}
	d.heap.Insert(dN, 0)
	dijkstraNodes[source.Id()] = dN

	// while there are nodes in the heap, extract minimum and update out nodes distances
	for d.heap.Num() > 0 {
		i, v := d.heap.ExtractMin()
		currentHeapNode := node{
			dist: v,
			node: i.(*dijkstraNode),
		}
		currentDijkstraNode := currentHeapNode.Tag().(*dijkstraNode)
		currentRealNode := currentDijkstraNode.realNode
		explored[currentRealNode.Id()] = ""
		
		// update node with distance from landmark
		currentRealNode.AddFromLandmark(landmark, currentDijkstraNode.accumulatedDist)
		
		// update node with distance to landmark
		as := astar.NewAstar(d.graph, func(from interfaces.Node, to interfaces.Node) float64 {return 0})
		_, weight, _ := as.Path(currentRealNode, source)
		currentRealNode.AddToLandmark(landmark, weight)
		
		// update out nodes distances in heap
		for _, edge := range currentRealNode.OutEdges() {
			destinationRealNode := edge.To()
			destinationDijkstraNode := dijkstraNodes[destinationRealNode.Id()]
			if _, ok := explored[destinationRealNode.Id()]; !ok {
				newDistance := edge.Weight() + currentDijkstraNode.accumulatedDist

				err := d.heap.DecreaseKey(destinationDijkstraNode, newDistance)
				if err == nil {
					destinationDijkstraNode.last = currentDijkstraNode
					destinationDijkstraNode.lastDistance = edge.Weight()
					destinationDijkstraNode.edgeType = edge.EdgeType()
					destinationDijkstraNode.accumulatedDist = edge.Weight() + currentDijkstraNode.accumulatedDist
				}
			}
		}
	}
	return
}

func NewDijkstra(graph interfaces.Graph) *dijkstra {
	return &dijkstra{graph: graph, heap: fibHeap.NewFibHeap()}
}