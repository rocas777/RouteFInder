package landmarks

import (
	"edaa/internals/interfaces"
	"edaa/internals/types"
	"fmt"
	"math"
	"sort"

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

type Dijkstra struct {
	heap  *fibHeap.FibHeap
	graph interfaces.Graph
	lats  [12]float64
	lons  [12]float64
}

func (d *Dijkstra) ProcessLandmarks() {
	landmarks := initLandmarks(d.graph.NodesMap())
	for i, n := range landmarks {
		d.lats[i] = n.Latitude()
		d.lons[i] = n.Longitude()
		d.PreprocessLandmark(i, n)
	}
}

func initLandmarks(nodesMap map[string]interfaces.Node) []interfaces.Node {
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

func (d *Dijkstra) PreprocessLandmark(landmark int, source interfaces.Node) {
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

func NewDijkstra(graph interfaces.Graph) *Dijkstra {
	return &Dijkstra{graph: graph, heap: fibHeap.NewFibHeap()}
}

func (d *Dijkstra) SelectActiveLandmarks(from interfaces.Node, to interfaces.Node) [3]int {
	fromL := from.GetFromLandmarks()
	toL := to.GetFromLandmarks()
	
	var active [3]int
	
	potentials := []struct {
		Index int
		Potential float64
	}{
		{0, math.Abs(fromL[0] - toL[0])},
		{1, math.Abs(fromL[1] - toL[1])},
		{2, math.Abs(fromL[2] - toL[2])},
		{3, math.Abs(fromL[3] - toL[3])},
		{4, math.Abs(fromL[4] - toL[4])},
		{5, math.Abs(fromL[5] - toL[5])},
		{6, math.Abs(fromL[6] - toL[6])},
		{7, math.Abs(fromL[7] - toL[7])},
		{8, math.Abs(fromL[8] - toL[8])},
		{9, math.Abs(fromL[9] - toL[9])},
		{10, math.Abs(fromL[10] - toL[10])},
		{11, math.Abs(fromL[11] - toL[11])},
	}
	
	fmt.Println("before sort:", potentials);
	
	sort.SliceStable(potentials, func(i, j int) bool {
		return potentials[i].Potential < potentials[j].Potential;
	});
	fmt.Println("after sort:", potentials);
	
	active[0] = potentials[11].Index;
	active[1] = potentials[10].Index;
	active[2] = potentials[9].Index;
	
	fmt.Println(active);
	return active
}

func Heuristic(from interfaces.Node, to interfaces.Node, activeLandmarks [3]int) float64 {
	fromL := from.GetFromLandmarks()
	toL := to.GetFromLandmarks()

	var max float64

	for _, i := range activeLandmarks {
		potential := math.Abs(fromL[i] - toL[i])
		if potential > max {
			max = potential
		}
	}

	return max
}
