package menuHelper

import (
	"edaa/internals/algorithms/connectivity/tarjan"
	"edaa/internals/interfaces"
)

func Connectivity(g interfaces.Graph) {
	disconnectedComponents, number := tarjan.TarjanGetStronglyConnectedComponents(g)
	tarjan.PrintStronglyConnectedComponentsSizes(number, disconnectedComponents)
}
