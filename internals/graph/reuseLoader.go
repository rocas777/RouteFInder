package graph

import (
	"edaa/internals/interfaces"
	"errors"
	"github.com/gocarina/gocsv"
	"os"
	"strings"
)

type tempReuseNode struct {
	Latitude   float64 `csv:"Latitude"`
	Longitude  float64 `csv:"Longitude"`
	Name       string  `csv:"Name"`
	Zone       string  `csv:"Zone"`
	ID         string  `csv:"Code"`
	Referenced bool    `csv:"Referenced"`
	IsStation  bool    `csv:"IsStation"`
}

type tempReuseEdge struct {
	From   string  `csv:"From"`
	To     string  `csv:"To"`
	Weight float64 `csv:"Weight"`
}

func InitReuse(g interfaces.Graph) {
	os.Remove("data/reuse/path_edges.csv")
	g.SetNodesMap(make(map[string]interfaces.Node))
	g.SetWalkableNodes(make(map[string]interfaces.Node))
	g.SetBusableNodes(make(map[string]interfaces.Node))
	g.SetMetroableNodes(make(map[string]interfaces.Node))

	if _, err := os.Stat("data/reuse/nodes.csv"); errors.Is(err, os.ErrNotExist) {
		println("You should run setup first")
		os.Exit(1)
	}
	if _, err := os.Stat("data/reuse/edges.csv"); errors.Is(err, os.ErrNotExist) {
		println("You should run setup first")
		os.Exit(1)
	}
	loadNodes(g)
	loadEdges(g)
	println("loaded", len(g.Nodes()), "nodes")
}

func loadNodes(g interfaces.Graph) {
	// load csv data into temp struct
	in, err := os.Open("data/reuse/nodes.csv")
	if err != nil {
		panic(err)
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			panic(err.Error())
		}
	}(in)
	var tempNodes []*tempReuseNode
	if err := gocsv.UnmarshalFile(in, &tempNodes); err != nil {
		panic(err)
	}
	for _, node := range tempNodes {
		var n interfaces.Node
		if node.IsStation {
			n = NewStationNode(node.Latitude, node.Longitude, node.Name, node.Zone, node.ID)
		} else {
			n = NewNormalNode(node.Latitude, node.Longitude, node.Name, node.Zone, node.ID)
		}
		g.AddNode(n)
		switch {
		case strings.Contains(node.ID, "walk"):
			g.WalkableNodes()[node.ID] = n
		case strings.Contains(node.ID, "bus"):
			g.BusableNodes()[node.ID] = n
		case strings.Contains(node.ID, "metro"):
			g.MetroableNodes()[node.ID] = n
		}
	}
}

func loadEdges(g interfaces.Graph) {
	in, err := os.Open("data/reuse/edges.csv")
	if err != nil {
		panic(err)
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			panic(err.Error())
		}
	}(in)
	var tempEdges []*tempReuseEdge
	if err := gocsv.UnmarshalFile(in, &tempEdges); err != nil {
		panic(err)
	}
	for _, edge := range tempEdges {
		to, _ := g.GetNode(edge.To)
		from, _ := g.GetNode(edge.From)
		if to != nil && from != nil {
			from.AddDestination(to, edge.Weight)
		}
	}
}
