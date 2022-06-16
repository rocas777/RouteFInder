package astar

import (
	"edaa/internals/interfaces"
	"os"

	"github.com/gocarina/gocsv"
)

type exportEdge struct {
	From   string
	To     string
	Weight float64
}

func newExportEdge(edge interfaces.Edge) *exportEdge {
	return &exportEdge{From: edge.From().Id(), To: edge.To().Id(), Weight: edge.Weight()}
}

func ExportEdges(edges []interfaces.Edge) {
	var outEdges = make([]*exportEdge, len(edges))
	for i, edge := range edges {
		outEdges[i] = newExportEdge(edge)
	}
	export(outEdges, "data/reuse/path_edges.csv")
}

func ExportEdgesGenetics(edges []interfaces.Edge) {
	var outEdges = make([]*exportEdge, len(edges))
	for i, edge := range edges {
		outEdges[i] = newExportEdge(edge)
	}
	export(outEdges, "data/reuse/path_edges_genetics.csv")
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
