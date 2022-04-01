package graph

import (
	"edaa/internals/utils"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type tempBusNode struct {
	Code      string  `csv:"Code"`
	Name      string  `csv:"Name"`
	Zone      string  `csv:"Zone"`
	Latitude  float64 `csv:"Latitude"`
	Longitude float64 `csv:"Longitude"`
}

type tempMetroNode struct {
	Code      string  `csv:"stop_id"`
	Name      string  `csv:"stop_name"`
	Zone      string  `csv:"zone_id"`
	Latitude  float64 `csv:"stop_lat"`
	Longitude float64 `csv:"stop_lon"`
}
type tempStopTimeStruct struct {
	TripId        string `csv:"trip_id"`
	ArrivalTime   string `csv:"arrival_time"`
	DepartureTime string `csv:"departure_time"`
	StopId        string `csv:"stop_id"`
}

type Osm struct {
	XMLName xml.Name  `xml:"osm"`
	Nodes   []OsmNode `xml:"node"`
	Ways    []OsmWay  `xml:"way"`
}

type OsmNode struct {
	XMLName xml.Name `xml:"node"`
	ID      string   `xml:"id,attr"`
	Lat     float64  `xml:"lat,attr"`
	Lon     float64  `xml:"lon,attr"`
	Tags    []OsmTag `xml:"tag"`
}

type OsmTag struct {
	XMLName xml.Name `xml:"tag"`
	Key     string   `xml:"k,attr"`
	Value   string   `xml:"v,attr"`
}

type OsmWay struct {
	XMLName xml.Name `xml:"way"`
	Nodes   []OsmND  `xml:"nd"`
	Tags    []OsmTag `xml:"tag"`
}

type OsmND struct {
	XMLName xml.Name `xml:"nd"`
	Ref     string   `xml:"ref,attr"`
}

func (g *Graph) initBus() {

	// load csv data into temp struct
	helperMap := make(map[string]*Node)
	in, err := os.Open("data/bus/stops.csv")
	if err != nil {
		panic(err)
	}
	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			panic(err.Error())
		}
	}(in)
	var tempNodes []*tempBusNode
	if err := gocsv.UnmarshalFile(in, &tempNodes); err != nil {
		panic(err)
	}

	// pass temp structure data To final node
	// create map for faster initialization
	for _, tempNode := range tempNodes {
		node := NewStationNode(tempNode.Latitude, tempNode.Longitude, tempNode.Name, tempNode.Zone, "bus_"+tempNode.Code)
		g.AddNode(node)
		g.BusableNodes[node.Code] = node
		helperMap[tempNode.Code] = node
	}

	// read files From lines directory
	files, err := ioutil.ReadDir("data/bus/lines")
	if err != nil {
		log.Panic(err)
	}

	// load line information
	for _, file := range files {
		if !file.IsDir() {
			csvFile, err := os.Open("data/bus/lines/" + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			csvLines, err := csv.NewReader(csvFile).ReadAll()
			totalDistance := 0.0
			totalTime := 0.0
			var lastNode *Node = nil

			// calculate line distance
			for i, line := range csvLines {
				if i != 0 {
					currentNode := helperMap[line[0]]
					if lastNode != nil {
						totalDistance += utils.GetDistance(currentNode.latitude, currentNode.longitude, lastNode.latitude, lastNode.longitude)
					}
					lastNode = currentNode
				}
			}

			lastNode = nil
			// setup Edges
			for i, line := range csvLines {
				if i == 0 {
					totalTime = float64(utils.StringToInt(line[0])) * 3600
				} else {
					currentNode := helperMap[line[0]]
					if lastNode != nil {
						lastNode.AddDestination(currentNode, totalTime/totalDistance*utils.GetDistance(currentNode.latitude, currentNode.longitude, lastNode.latitude, lastNode.longitude))
						// currentNode.AddDestination(lastNode, totalTime/totalDistance*utils.GetDistance(currentNode.latitude, currentNode.longitude, lastNode.latitude, lastNode.longitude))
					}
					lastNode = currentNode
				}
			}
			csvFile.Close()
		}
	}
}

func (g *Graph) initMetro() {

	//load nodes
	helperMap := func(g *Graph) map[string]*Node {
		helperMap := make(map[string]*Node)
		in, err := os.Open("data/metro/stops.txt")
		if err != nil {
			panic(err)
		}
		defer func(in *os.File) {
			err := in.Close()
			if err != nil {
				panic(err.Error())
			}
		}(in)
		var tempNodes []*tempMetroNode
		if err := gocsv.UnmarshalFile(in, &tempNodes); err != nil {
			panic(err)
		}

		for _, tempNode := range tempNodes {
			node := NewStationNode(tempNode.Latitude, tempNode.Longitude, tempNode.Name, tempNode.Zone, "metro_"+tempNode.Code)
			g.AddNode(node)
			g.MetroableNodes[node.Code] = node
			helperMap[tempNode.Code] = node
		}
		return helperMap
	}(g)

	//load edges
	func(helperMap map[string]*Node) {
		lines := make(map[string][]*tempStopTimeStruct)
		in, err := os.Open("data/metro/stop_times.txt")
		if err != nil {
			panic(err)
		}
		defer func(in *os.File) {
			err := in.Close()
			if err != nil {
				panic(err.Error())
			}
		}(in)
		var tempStopTimes []*tempStopTimeStruct
		if err := gocsv.UnmarshalFile(in, &tempStopTimes); err != nil {
			panic(err)
		}
		for _, stopTime := range tempStopTimes {
			if _, ok := lines[stopTime.TripId]; ok {
				lines[stopTime.TripId] = append(lines[stopTime.TripId], stopTime)
			} else {
				lines[stopTime.TripId] = make([]*tempStopTimeStruct, 1)
				lines[stopTime.TripId][0] = stopTime
			}
		}

		var lastNode *Node
		var lastStopTime *tempStopTimeStruct
		for _, line := range lines {
			lastNode = nil
			lastStopTime = nil
			for _, stopTime := range line {
				currentNode := helperMap[stopTime.StopId]
				currentStopTime := stopTime
				if lastNode != nil && lastStopTime != nil {
					lastTime, _ := time.Parse("15:04:05", lastStopTime.DepartureTime)
					currentTime, _ := time.Parse("15:04:05", currentStopTime.DepartureTime)

					seconds := currentTime.Sub(lastTime)
					lastNode.AddDestination(currentNode, seconds.Seconds())
				}
				lastNode = currentNode
				lastStopTime = currentStopTime
			}
		}
	}(helperMap)
}

func (g *Graph) initRoads() {
	helperMap := make(map[string]*Node)
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
	var osm Osm
	if err := xml.Unmarshal(roadsData, &osm); err != nil {
		panic(err)
	}

	for _, node := range osm.Nodes {
		isStation := false
		for _, tag := range node.Tags {
			if (tag.Key == "highway" && tag.Value == "bus_stop") || (tag.Key == "station" && tag.Value == "subway") {
				isStation = true
				break
			}
		}
		if isStation {
			//helperMap[node.ID] = NewStationNode(node.Lat, node.Lon, "", "", node.ID)
			helperMap[node.ID] = NewNormalNode(node.Lat, node.Lon, "", "", "walk_"+node.ID)
		} else {
			helperMap[node.ID] = NewNormalNode(node.Lat, node.Lon, "", "", "walk_"+node.ID)
		}
	}

	for _, way := range osm.Ways {
		isTwoWay := true
		for _, tag := range way.Tags {
			if tag.Key == "oneway" && tag.Value == "yes" {
				isTwoWay = false
				break
			}
		}
		var lastNode *Node
		for _, node := range way.Nodes {
			if lastNode == nil {
				lastNode = helperMap[node.Ref]
			} else {
				currentNode := helperMap[node.Ref]
				dist := utils.GetDistance(lastNode.latitude, lastNode.longitude, currentNode.latitude, currentNode.longitude)
				lastNode.AddDestination(currentNode, dist)
				if isTwoWay {
					currentNode.AddDestination(lastNode, dist)
				}
			}
		}
	}
	for _, node := range helperMap {
		g.AddNode(node)
		g.WalkableNodes[node.Code] = node
	}
}
