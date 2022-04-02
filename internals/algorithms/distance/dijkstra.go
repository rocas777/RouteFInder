package distance

import (
	"edaa/internals/graph"
	"edaa/internals/interfaces"
	"math"
)

func DijkstraForNode(g *graph.Graph, n interfaces.Node) {
	var heap *graph.FibHeap = graph.NewFibHeap()
	
	for _, m := range g.NodesMap() {
		if m == n {
			continue
		}
		m.SetDistance(math.MaxFloat64)
		m.SetPrevious("")
		m.SetVisited(false)
		heap.AddNode(m, m.Distance())
	}
	n.SetDistance(0)
	n.SetPrevious("")
	n.SetVisited(false)
	heap.AddNode(n, n.Distance())
	
	for len(heap.NodesMap) != 0 {
		heap.PrintFibHeap()
		var u = g.NodesMap()[heap.PopMin()]
		u.SetVisited(true)
		
		for _, e := range u.OutEdges() {
			var v = e.To()
			if v.Visited() {
				continue
			}
			
			if v.Distance() > u.Distance() + e.Weight() {
				v.SetDistance(u.Distance() + e.Weight())
				v.SetPrevious(u.Id())
				heap.DecreaseKey(v.Id(), v.Distance())
			}
		}
	}
}