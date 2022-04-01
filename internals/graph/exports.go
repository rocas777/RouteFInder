package graph

import (
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

func newExportNode(node *Node) *exportNode {
	return &exportNode{Latitude: node.latitude, Longitude: node.longitude, IsStation: node.IsStation, Name: node.name, Zone: node.zone, Code: node.Code}
}
func newExportEdge(edge *Edge) *exportEdge {
	return &exportEdge{From: edge.From.Code, To: edge.To.Code, Weight: edge.Weight}
}

func (g *Graph) ExportNodes(filePath string) {
	outNodes := make([]*exportNode, len(g.Nodes))
	for i, node := range g.Nodes {
		outNodes[i] = newExportNode(node)
	}
	export(outNodes, filePath)
}
func (g *Graph) ExportEdges(filePath string) {
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
