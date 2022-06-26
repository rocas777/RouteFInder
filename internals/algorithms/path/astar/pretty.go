package astar

import (
	"edaa/internals/interfaces"
	"edaa/internals/types"
	"edaa/internals/utils"
	"fmt"
	"math"
)

func PreetyDisplay(path []interfaces.Edge, pathTime float64, explored int, startNode, endNode interfaces.Node) {
	fmt.Println("")
	fmt.Println("")
	roadTime := 0.0
	busTime := 0.0
	metroTime := 0.0
	for _, vertex := range path {
		fmt.Printf("%s -> %s  %s  %f\n", vertex.From().Id(), vertex.To().Id(), vertex.EdgeType(), vertex.Weight())
		switch vertex.EdgeType() {
		case types.Road:
			roadTime += vertex.Weight()
		case types.Bus:
			busTime += vertex.Weight()
		case types.Metro:
			metroTime += vertex.Weight()
		}
	}
	fmt.Println("")
	fmt.Println("")

	fmt.Printf("Start: %f,%f\n", startNode.Latitude(), startNode.Longitude())
	fmt.Printf("End: %f,%f\n", endNode.Latitude(), endNode.Longitude())
	fmt.Printf("Time: %d:%d\n", int(pathTime/60), int((pathTime/60-math.Floor(pathTime/60))*60))
	fmt.Printf("Distance: %f km\n", utils.GetDistanceBetweenNodes(startNode, endNode)/1000)

	fmt.Println("")
	fmt.Printf("Road Time: %d s\n", int(roadTime))
	fmt.Printf("Bus Time: %d s\n", int(busTime))
	fmt.Printf("Metro Time: %d s\n", int(metroTime))

	fmt.Println("Explored nodes:", explored)
}
