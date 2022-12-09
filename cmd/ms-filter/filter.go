package main

import (
	"context"
	"edaa/internals/exports/reuse"
	"edaa/internals/graph"
	"encoding/xml"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
)

func main() {
	argsWithoutProg := os.Args[1:]

	maxLat, _ := strconv.ParseFloat(argsWithoutProg[0], 32)
	minLon, _ := strconv.ParseFloat(argsWithoutProg[1], 32)
	minLat, _ := strconv.ParseFloat(argsWithoutProg[2], 32)
	maxLon, _ := strconv.ParseFloat(argsWithoutProg[3], 32)
	inputFile := argsWithoutProg[4]
	outputFileNodes := argsWithoutProg[5]
	outputFileEdges := argsWithoutProg[6]

	err := compressFile(minLat, maxLat, maxLon, minLon, inputFile)

	runtime.GC()

	g := graph.Graph{}
	g.InitOnlyRoads()

	os.Remove("compressed.xml")

	reuse.ExportNodes(&g, outputFileNodes)
	reuse.ExportEdges(&g, outputFileEdges)

	if err != nil {
		return
	}
}

func compressFile(minLat float64, maxLat float64, maxLon float64, minLon float64, inputFile string) error {
	println("compressing")
	bounds := &osm.Bounds{
		MinLat: minLat,
		MaxLat: maxLat,
		MaxLon: maxLon,
		MinLon: minLon,
	}

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := osmpbf.New(context.Background(), file, runtime.GOMAXPROCS(-1))

	defer scanner.Close()

	var nodes osm.Nodes
	var ways osm.Ways
	var relations osm.Relations

	nodesMap := make(map[osm.NodeID]bool)
	nodesInWay := make(map[osm.NodeID]bool)

	//filter := "railway"
	filter := "highway"

	ignore := "footway"
	scanner.SkipRelations = true

	for scanner.Scan() {
		switch o := scanner.Object().(type) {
		case *osm.Node:
			if !(o.Lat >= bounds.MaxLat || o.Lat <= bounds.MinLat || o.Lon >= bounds.MaxLon || o.Lon <= bounds.MinLon) {
				nodes = append(nodes, o)
				nodesMap[o.ID] = true
			} else {
				nodesMap[o.ID] = false
			}
		case *osm.Way:
			tag := o.Tags.Find(filter)
			if tag != "" {
				ways = append(ways, o)
				tag = o.Tags.Find(ignore)
				if tag == "" {
					ways = append(ways, o)
				}
			}
		}
	}

	//filter ways by node
	var newWays osm.Ways
	for _, w := range ways {
		skip := false
		for _, n := range w.Nodes {
			if nodesMap[n.ID] == false {
				skip = true
			}
			nodesInWay[n.ID] = true
		}
		if !skip {
			newWays = append(newWays, w)
		}
	}
	ways = newWays

	//filter nodes by ways and tags
	var newNodes osm.Nodes
	for _, n := range nodes {
		//add if relate dto highway or in a way
		/*tag := n.Tags.Find(filter)
		if tag != "" {
			newNodes = append(newNodes, n)
		} else*/if _, ok := nodesInWay[n.ID]; ok {
			newNodes = append(newNodes, n)
		}
	}
	ways = newWays
	nodes = newNodes

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Preparing XML", len(nodes), len(ways))

	o := osm.OSM{
		Version:   0.6,
		Bounds:    bounds,
		Nodes:     nodes,
		Ways:      ways,
		Relations: relations,
	}

	data, _ := xml.MarshalIndent(o, " ", "	")

	fmt.Println("Xml prepared, writing to file")
	err = os.WriteFile("compressed.xml", data, 0644)
	return err
}
