package menuHelper

import (
	"edaa/internals/algorithms/path/landmarks"
	"edaa/internals/interfaces"
)

func Landmarks (g interfaces.Graph) {
	d := landmarks.NewDijkstra(g)
	d.ProcessLandmarks()
}