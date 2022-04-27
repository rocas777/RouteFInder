package graph

import (
	"edaa/internals/interfaces"
	"edaa/internals/utils"
	"github.com/gocarina/gocsv"
	"os"
	"time"
)

type tempMetroNode struct {
	Code      string  `csv:"stop_id"`
	Name      string  `csv:"stop_name"`
	Zone      string  `csv:"zone_id"`
	Latitude  float64 `csv:"stop_lat"`
	Longitude float64 `csv:"stop_lon"`
}
type tempStopTimeStruct struct {
	TripID        string `csv:"trip_id"`
	ArrivalTime   string `csv:"arrival_time"`
	DepartureTime string `csv:"departure_time"`
	StopID        string `csv:"stop_id"`
}

func loadMetroNodes(g interfaces.Graph) map[string]interfaces.Node {
	helperMap := make(map[string]interfaces.Node)
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
		g.MetroableNodes()[node.Id()] = node
		helperMap[tempNode.Code] = node
	}
	return helperMap
}

func loadMetroEdges(helperMap map[string]interfaces.Node) {
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
		if _, ok := lines[stopTime.TripID]; ok {
			lines[stopTime.TripID] = append(lines[stopTime.TripID], stopTime)
		} else {
			lines[stopTime.TripID] = make([]*tempStopTimeStruct, 1)
			lines[stopTime.TripID][0] = stopTime
		}
	}

	var lastNode interfaces.Node
	var lastStopTime *tempStopTimeStruct
	fastestSpeed := 0.0
	for _, line := range lines {
		lastNode = nil
		lastStopTime = nil
		for _, stopTime := range line {
			currentNode := helperMap[stopTime.StopID]
			currentStopTime := stopTime
			if lastNode != nil && lastStopTime != nil {
				lastTime, err1 := time.Parse("15:04:05", lastStopTime.DepartureTime)
				currentTime, err2 := time.Parse("15:04:05", currentStopTime.DepartureTime)

				if err1 != nil || err2 != nil {
					continue
				}

				dist := utils.GetDistanceBetweenNodes(lastNode, currentNode)

				seconds := currentTime.Sub(lastTime)

				if dist/seconds.Seconds() > fastestSpeed && seconds.Seconds() > 0.0 {
					fastestSpeed = dist / seconds.Seconds()
				}

				lastNode.AddDestination(currentNode, seconds.Seconds())
			}
			lastNode = currentNode
			lastStopTime = currentStopTime
		}
	}
	//panic(fastestSpeed)
}

func InitMetro(g interfaces.Graph) {
	// load nodes
	helperMap := loadMetroNodes(g)

	// load outEdges
	loadMetroEdges(helperMap)
}
