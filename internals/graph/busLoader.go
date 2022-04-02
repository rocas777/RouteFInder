package graph

import (
	"edaa/internals/interfaces"
	"edaa/internals/utils"
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"log"
	"os"
)

type tempBusNode struct {
	Code      string  `csv:"Code"`
	Name      string  `csv:"Name"`
	Zone      string  `csv:"Zone"`
	Latitude  float64 `csv:"Latitude"`
	Longitude float64 `csv:"Longitude"`
}

func InitBus(g interfaces.Graph) {
	// load csv data into temp struct
	helperMap := make(map[string]interfaces.Node)
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

	// pass temp structure data to final node
	// create map for faster initialization
	for _, tempNode := range tempNodes {
		node := NewStationNode(tempNode.Latitude, tempNode.Longitude, tempNode.Name, tempNode.Zone, "bus_"+tempNode.Code)
		g.BusableNodes()[node.Id()] = node
		g.AddNode(node)
		helperMap[tempNode.Code] = node
	}

	// read files from lines directory
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
			csvLines, _ := csv.NewReader(csvFile).ReadAll()
			totalDistance := 0.0
			totalTime := 0.0
			var lastNode interfaces.Node

			// calculate line distance
			for i, line := range csvLines {
				currentNode := helperMap[line[0]]
				if i != 0 {
					if lastNode != nil {
						totalDistance += utils.GetDistance(currentNode.Latitude(), currentNode.Longitude(), lastNode.Latitude(), lastNode.Longitude())
					}
					lastNode = currentNode
				}
			}

			lastNode = nil
			// setup outEdges
			for i, line := range csvLines {
				if i == 0 {
					totalTime = float64(utils.StringToInt(line[0])) * 60
				} else {
					currentNode := helperMap[line[0]]

					if lastNode != nil {
						lastNode.AddDestination(currentNode, totalTime/totalDistance*utils.GetDistance(currentNode.Latitude(), currentNode.Longitude(), lastNode.Latitude(), lastNode.Longitude()))
						// currentNode.AddDestination(lastNode, totalTime/totalDistance*utils.GetDistance(currentNode.latitude, currentNode.longitude, lastNode.latitude, lastNode.longitude))
					}
					lastNode = currentNode
				}
			}
		}
	}
}
