package filtering

import "edaa/internals/interfaces"

// removes every node that does not belong to the biggest SCC and does not include a station
func RemoveUnconnectedNodes(g interfaces.Graph, disconnectedComponents [][]interfaces.SCC) {
	biggestScc := 0
	for _, components := range disconnectedComponents {
		for _, component := range components {
			if biggestScc <= len(component.Nodes()) {
				biggestScc = len(component.Nodes())
			}
		}
	}
	nodesToBeRemoved := make([]interfaces.Node, 0)
	for _, components := range disconnectedComponents {
		for _, component := range components {
			if len(component.Nodes()) < biggestScc && !component.HasStation() {
				nodesToBeRemoved = append(nodesToBeRemoved, component.Nodes()...)
			}
		}
	}
	g.RemoveNodes(nodesToBeRemoved)
}
