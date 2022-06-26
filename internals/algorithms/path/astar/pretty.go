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

	p,pathTime := costCostNoPenalty(path)

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

	fmt.Printf("Price: %f\n", p)

	fmt.Println("Explored nodes:", explored)
}
var prices map[int]float64 = map[int]float64{
	2 : 1.25,
	3 : 1.6,
	4 : 2.0,
	5 : 2.4,
	6 : 2.85,
	7 : 3.25,
	8 : 3.65,
	9 : 4.05,
}
func costCostNoPenalty(path []interfaces.Edge) (float64,float64){
	t := 0.0
	cost := 0.0
	lastBus := true
	zones := make(map[string]interface{})
	for _, v := range path {
		if v.To().IsStation() {
			zones[v.To().Zone()] = "";
		}
		if lastBus && v.EdgeType() == types.Road{
			cost += 2.0
		}
		if v.EdgeType() != types.Road{
			lastBus = true;
			t += v.Weight()
		}else{
			lastBus = false
			t += v.Weight()
		}
	}
	/*fmt.Println(cost)
	fmt.Println(prices[len(zones)])
	fmt.Println(walk)
	fmt.Println(transport)
	fmt.Println()*/
	return math.Min(cost,prices[len(zones)]),t
}
