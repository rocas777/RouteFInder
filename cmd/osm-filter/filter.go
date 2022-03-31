package main

import (
	"context"
	"encoding/xml"
	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
	"os"
	"runtime"
)

func main() {

	bounds := &osm.Bounds{
		MinLat: 41.08061,
		MaxLat: 41.25590,
		MaxLon: -8.457085,
		MinLon: -8.713925,
	}

	file, err := os.Open("data/road/portugal-latest.osm.pbf")
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
		tag := n.Tags.Find(filter)
		if tag != "" {
			newNodes = append(newNodes, n)
		} else if _, ok := nodesInWay[n.ID]; ok {
			newNodes = append(newNodes, n)
		}
	}
	ways = newWays
	nodes = newNodes

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	println("Preparing XML", len(nodes), len(ways))

	o := osm.OSM{
		Version:   0.6,
		Bounds:    bounds,
		Nodes:     nodes,
		Ways:      ways,
		Relations: relations,
	}

	data, _ := xml.MarshalIndent(o, " ", "	")

	println("Xml prepared, writing to file")
	err = os.WriteFile("data/road/compressed.xml", data, 0644)
	if err != nil {
		return
	}
}
