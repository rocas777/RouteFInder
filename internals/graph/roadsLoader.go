package graph

import (
	"edaa/internals/interfaces"
	"edaa/internals/utils"
	"encoding/xml"
	"io/ioutil"
	"os"
)

const walkSpeed = 1.0

type osm struct {
	XMLName xml.Name  `xml:"osm"`
	Nodes   []osmNode `xml:"node"`
	Ways    []osmWay  `xml:"way"`
}

type osmNode struct {
	XMLName xml.Name `xml:"node"`
	ID      string   `xml:"id,attr"`
	Lat     float64  `xml:"lat,attr"`
	Lon     float64  `xml:"lon,attr"`
	Tags    []osmTag `xml:"tag"`
}

type osmTag struct {
	XMLName xml.Name `xml:"tag"`
	Key     string   `xml:"k,attr"`
	Value   string   `xml:"v,attr"`
}

type osmWay struct {
	XMLName xml.Name `xml:"way"`
	Nodes   []osmND  `xml:"nd"`
	Tags    []osmTag `xml:"tag"`
}

type osmND struct {
	XMLName xml.Name `xml:"nd"`
	Ref     string   `xml:"ref,attr"`
}

func InitRoads(g interfaces.Graph) {
	println("Setting Road Nodes")
	helperMap := make(map[string]interfaces.Node)
	in, err := os.Open("data/road/compressed.xml")
	if err != nil {
		panic(err)
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			panic(err.Error())
		}
	}(in)
	roadsData, _ := ioutil.ReadAll(in)
	var osmV osm
	if err := xml.Unmarshal(roadsData, &osmV); err != nil {
		panic(err)
	}

	for _, node := range osmV.Nodes {
		helperMap[node.ID] = NewNormalNode(node.Lat, node.Lon, "", "", "walk_"+node.ID)
	}
	for _, way := range osmV.Ways {
		isTwoWay := true
		for _, tag := range way.Tags {
			if tag.Key == "oneway" && tag.Value == "yes" {
				isTwoWay = false
				break
			}
		}
		var lastNode interfaces.Node
		for _, node := range way.Nodes {
			currentNode := helperMap[node.Ref]
			if lastNode == nil {
				lastNode = currentNode
			} else {
				dist := utils.GetDistance(lastNode.Latitude(), lastNode.Longitude(), currentNode.Latitude(), currentNode.Longitude())
				lastNode.AddDestination(currentNode, dist)
				if isTwoWay {
					currentNode.AddDestination(lastNode, dist)
				}
				lastNode = currentNode
			}
		}
	}
	for _, node := range helperMap {
		g.AddNode(node)
		g.WalkableNodes()[node.Id()] = node
	}
}
