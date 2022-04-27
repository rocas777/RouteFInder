package utils

import (
	"edaa/internals/interfaces"
	"fmt"
	"github.com/umahmood/haversine"
	"runtime"
	"strconv"
)

//return distance in meters
func GetDistance(lat1, lon1, lat2, lon2 float64) float64 {
	spot1 := haversine.Coord{Lat: lat1, Lon: lon1}
	spot2 := haversine.Coord{Lat: lat2, Lon: lon2}
	_, km := haversine.Distance(spot1, spot2)
	return km * 1000
}

func GetDistanceBetweenNodes(n1, n2 interfaces.Node) float64 {
	spot1 := haversine.Coord{Lat: n1.Latitude(), Lon: n1.Longitude()}
	spot2 := haversine.Coord{Lat: n2.Latitude(), Lon: n2.Longitude()}
	_, km := haversine.Distance(spot1, spot2)
	return km * 1000
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func PrintMemUsage() {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func RemoveString(stringList []string, element string) []string {
	var index int = containsString(stringList, element)

	if index != -1 {
		stringList[index] = stringList[len(stringList)-1]
		return stringList[:len(stringList)-1]
	}

	return stringList
}

func containsString(stringList []string, element string) int {
	for i, v := range stringList {
		if v == element {
			return i
		}
	}
	return -1
}
