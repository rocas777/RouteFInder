package graph

import (
	"edaa/internals/utils"
	"encoding/csv"
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
	helperMap := loadMetroStops(g)
	loadMetroEdges(helperMap)
}

func loadMetroEdges(helperMap map[string]*Node) {
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
}

func loadMetroStops(g *Graph) map[string]*Node {
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
		helperMap[tempNode.Code] = node
	}
	return helperMap
}
