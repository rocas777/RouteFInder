package reuse

import (
	"edaa/internals/graph"
	"edaa/internals/interfaces"
	"github.com/gocarina/gocsv"
	"os"
)

type exportNode struct {
	Latitude   float64
	Longitude  float64
	Name       string
	Zone       string
	Code       string
	Referenced bool
	IsStation  bool
}
type exportEdge struct {
	From   string
	To     string
	Weight float64
}

func newExportNode(node interfaces.Node) *exportNode {
	return &exportNode{Referenced: node.Referenced(), Latitude: node.Latitude(), Longitude: node.Longitude(), IsStation: node.IsStation(), Name: node.Name(), Zone: node.Zone(), Code: node.Id()}
}
func newExportEdge(edge interfaces.Edge) *exportEdge {
	return &exportEdge{From: edge.From().Id(), To: edge.To().Id(), Weight: edge.Weight()}
}

func ExportNodes(g *graph.Graph, filePath string) {
	outNodes := make([]*exportNode, len(g.Nodes()))
	for i, node := range g.Nodes() {
		outNodes[i] = newExportNode(node)
	}
	export(outNodes, filePath)
}
func ExportEdges(g *graph.Graph, filePath string) {
	edges := g.GetEdges()
	var outEdges = make([]*exportEdge, len(edges))
	for i, edge := range edges {
		outEdges[i] = newExportEdge(edge)
	}
	export(outEdges, filePath)
}

func export(export interface{}, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err.Error())
	}
	err = gocsv.MarshalFile(export, file)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
}
