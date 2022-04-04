package distance

import (
	"edaa/internals/graph"
	"fmt"
	"math"
)

// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm

func FloydWarshall(g *graph.Graph) {
	var heap *graph.FibHeap = graph.NewFibHeap()
	fmt.Printf(len(g.Nodes()))
	dist := make([][]float64, len(g.Nodes()))
	for i := range dist {
		di := make([]float64, len(g.Nodes()))
		for j := range di {
			di[j] = math.Inf(1)
		}
		di[i] = 0
		dist[i] = di
	}
	for u, graphs := range g.NodesMap() {
		for _, v := range graphs {
			dist[u][v.To()] = v.Distance()
		}
	}
	for k, dk := range dist {
		for _, di := range dist {
			for j, dij := range di {
				if d := di[k] + dk[j]; dij > d {
					di[j] = d
				}
			}
		}
	}
	return dist[0]
}
