package distance

import (
	"edaa/internals/graph"
	"edaa/internals/interfaces"
	"math"
)

// https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
 
func min(vars ...float64) float64 {
    min := vars[0]

    for _, i := range vars {
        if min > i {
            min = i
        }
    }

    return min
}
 
func FloydWarshall(g *graph.Graph) [][]float64{
	
	m := make(map[interfaces.Node]int)
	var n = len(g.Nodes())

    dist := make([][]float64,n)
    for i := range dist {
        di := make([]float64, n)
        for j := range di {
            di[j] = math.Inf(1)
        }
        di[i] = 0
        dist[i] = di
    }

	for i, node := range g.Nodes() {
		m[node] = i
	}	

	for i, gv := range g.Nodes() {
		for _, value := range gv.OutEdges() {
			dist[i][m[value.To()]] = value.Weight()
		}
	}	

    for k, _ := range g.Nodes() {
        for i, _ := range g.Nodes() {
			for j, _ := range g.Nodes() {
				dist[i][j] = min(dist[i][j],dist[i][k]+dist[k][j])
        	}
    	}
	}
	return dist
}
