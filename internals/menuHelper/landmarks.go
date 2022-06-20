package menuHelper

import (
	"edaa/internals/algorithms/path/landmarks"
	"edaa/internals/interfaces"
)

func Landmarks (g interfaces.Graph) (l *landmarks.Dijkstra) {
	d := landmarks.NewDijkstra(g)
	d.ProcessLandmarks()
	return d
}