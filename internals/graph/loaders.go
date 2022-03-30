package graph

import (
	"edaa/internals/utils"
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"log"
	"os"
)

func (g *Graph) initBus() {
	type tempBusNode struct {
		Code      string  `csv:"Code"`
		Name      string  `csv:"Name"`
		Zone      string  `csv:"Zone"`
		Latitude  float64 `csv:"Latitude"`
		Longitude float64 `csv:"Longitude"`
	}

	//load csv data into temp struct
	helperMap := make(map[string]*Node)
	in, err := os.Open("data/bus/stops.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()
	var tempNodes []*tempBusNode
	if err := gocsv.UnmarshalFile(in, &tempNodes); err != nil {
		panic(err)
	}

	//pass temp structure data To final node
	//create map for faster initialization
	for _, tempNode := range tempNodes {
		node := NewStationNode(tempNode.Latitude, tempNode.Longitude, tempNode.Name, tempNode.Zone, tempNode.Code)
		g.AddNode(node)
		helperMap[tempNode.Code] = node
	}

	//read files From lines directory
	files, err := ioutil.ReadDir("data/bus/lines")
	if err != nil {
		log.Fatal(err)
	}

	//load line information
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

			//calculate line distance
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
			//setup Edges
			for i, line := range csvLines {
				if i == 0 {
					totalTime = float64(utils.StringToInt(line[0]))
				} else {
					currentNode := helperMap[line[0]]
					if lastNode != nil {
						lastNode.AddDestination(currentNode, totalTime/totalDistance*utils.GetDistance(currentNode.latitude, currentNode.longitude, lastNode.latitude, lastNode.longitude))
						//currentNode.AddDestination(lastNode, totalTime/totalDistance*utils.GetDistance(currentNode.latitude, currentNode.longitude, lastNode.latitude, lastNode.longitude))
					}
					lastNode = currentNode
				}
			}
			csvFile.Close()
		}
	}
}
