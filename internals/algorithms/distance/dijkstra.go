package distance

import (
	"edaa/internals/graph"
	"math"
)

func DijkstraForNode(g *graph.Graph, n *graph.Node) {
	var heap *graph.FibHeap = graph.NewFibHeap()
	
	for _, m := range g.NodesMap {
		if m == n {
			continue
		}
		m.Distance = math.MaxFloat64
		m.Previous = ""
		m.Visited = false
		heap.AddNode(m, m.Distance)
	}
	n.Distance = 0
	n.Previous = ""
	n.Visited = false
	heap.AddNode(n, n.Distance)
	
	for len(heap.NodesMap) != 0 {
		heap.PrintFibHeap()
		var u = g.NodesMap[heap.PopMin()]
		u.Visited = true
		
		for _, e := range u.Edges {
			var v = e.To
			if v.Visited {
				continue
			}
			
			if v.Distance > u.Distance + e.Weight {
				v.Distance = u.Distance + e.Weight
				v.Previous = u.Code
				heap.DecreaseKey(v.Code, v.Distance)
			}
		}
	}
}