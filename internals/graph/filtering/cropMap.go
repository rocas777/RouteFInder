package filtering

import (
	"edaa/internals/interfaces"
)

func Crop(g interfaces.Graph, lat1, lon1, lat2, lon2 float64) {
	nodesToRemove := make([]interfaces.Node, 0)
	for _, n := range g.Nodes() {
		//println(lat1, lon1, lat2, lon2)
		if n.Latitude() > lat1 || n.Latitude() < lat2 || n.Longitude() > lon1 || n.Longitude() < lon2 {
			nodesToRemove = append(nodesToRemove, n)
		}
	}
	println("bef size", len(g.Nodes()))
	g.RemoveNodes(nodesToRemove)
	println("after size", len(g.Nodes()))
	for _, n := range g.Nodes() {
		outEdges := make([]interfaces.Edge, 0)
		for _, e := range n.OutEdges() {
			if _, error := g.GetNode(e.To().Id()); error == nil {
				outEdges = append(outEdges, e)
			}
		}

		inEdges := make([]interfaces.Edge, 0)
		for _, e := range n.InEdges() {
			if _, error := g.GetNode(e.To().Id()); error == nil {
				inEdges = append(inEdges, e)
			}
		}
		n.SetInEdges(inEdges)
		n.SetOutEdges(outEdges)
	}
}
