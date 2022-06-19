package landmarks

import (
	"edaa/internals/interfaces"
	"edaa/internals/types"
	"fmt"
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

func (d* dijkstra) ProcessLandmarks () {
	landmarks := initLandmarks(d.graph.NodesMap())
	for i, n := range landmarks {
		d.PreprocessLandmark(i, n)
	}
	fmt.Println(d.graph.Nodes()[0].GetFromLandmarks())
}

func initLandmarks(nodesMap map[string]interfaces.Node) ([]interfaces.Node) {
	landmarks := []interfaces.Node{
		nodesMap["walk_1227750946"],
		nodesMap["walk_1319947243"],
		nodesMap["walk_1419373987"],
		nodesMap["walk_1417923344"],
		nodesMap["walk_1438734394"],
		nodesMap["walk_1440898396"],
		nodesMap["walk_4438272432"],
		nodesMap["walk_7554539118"],
		nodesMap["walk_1562394871"],
		nodesMap["walk_1390095467"],
		nodesMap["walk_4580255471"],
		nodesMap["walk_1164520592"],
	}
	
	return landmarks
}

func (d *dijkstra) PreprocessLandmark(landmark int, source interfaces.Node) () {
	fmt.Println("preprocessing landmark", landmark)
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