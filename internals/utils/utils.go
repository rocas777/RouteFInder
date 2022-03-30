package utils

import (
	"github.com/umahmood/haversine"
	"strconv"
)

//return distance in meters
func GetDistance(lat1, lon1, lat2, lon2 float64) float64 {
	spot1 := haversine.Coord{Lat: lat1, Lon: lon1}
	spot2 := haversine.Coord{Lat: lat2, Lon: lon2}
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
