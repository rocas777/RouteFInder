package graph

import (
	"encoding/xml"
	"github.com/gocarina/gocsv"
	"hash/fnv"
	"os"
	"strconv"
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
	return &exportNode{Latitude: node.Latitude, Longitude: node.Longitude, IsStation: node.IsStation, Name: node.name, Zone: node.zone, Code: node.Code}
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

type OsmE struct {
	XMLName xml.Name   `xml:"osm"`
	Nodes   []OsmNodeE `xml:"node"`
	Ways    []OsmWayE  `xml:"way"`
	Version string     `xml:"version,attr"`
}

type OsmNodeE struct {
	XMLName xml.Name  `xml:"node"`
	ID      string    `xml:"id,attr"`
	Lat     float64   `xml:"lat,attr"`
	Lon     float64   `xml:"lon,attr"`
	Tags    []OsmTagE `xml:"tag"`
	Version string    `xml:"version,attr"`
}

type OsmTagE struct {
	XMLName xml.Name `xml:"tag"`
	Key     string   `xml:"k,attr"`
	Value   string   `xml:"v,attr"`
}

type OsmWayE struct {
	XMLName xml.Name  `xml:"way"`
	Nodes   []OsmNDE  `xml:"nd"`
	Tags    []OsmTagE `xml:"tag"`
}

type OsmNDE struct {
	XMLName xml.Name `xml:"nd"`
	Ref     string   `xml:"ref,attr"`
}

func hash(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}

func (g *Graph) ExportOsm(filePath string) {
	var outNodes []OsmNodeE
	for _, node := range g.Nodes {
		id := node.Code
		id = strconv.FormatUint(uint64(hash(id)), 10)
		outNodes = append(outNodes, OsmNodeE{
			ID:      id,
			Lat:     node.Latitude,
			Lon:     node.Longitude,
			Tags:    nil,
			Version: "1",
		})
	}
	out := OsmE{
		Nodes:   outNodes,
		Ways:    nil,
		Version: "0.6",
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err.Error())
	}
	data, _ := xml.MarshalIndent(out, "", "    ")
	if err != nil {
		panic(err.Error())
	}
	file.Write(data)
	defer file.Close()
}
