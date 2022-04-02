package visualization

import (
	"edaa/internals/graph"
	"edaa/internals/interfaces"
	"github.com/gocarina/gocsv"
	"os"
)

type exportNode struct {
	Latitude  float64
	Longitude float64
	IsStation bool
	Name      string
	Zone      string
	Code      string
}
type exportEdge struct {
	From   string
	To     string
	Weight float64
}

func newExportNode(node interfaces.Node) *exportNode {
	return &exportNode{Latitude: node.Latitude(), Longitude: node.Longitude(), IsStation: node.IsStation(), Name: node.Name(), Zone: node.Zone(), Code: node.Id()}
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
	var outEdges []*exportEdge
	for _, edge := range g.GetEdges() {
		outEdges = append(outEdges, newExportEdge(edge))
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
