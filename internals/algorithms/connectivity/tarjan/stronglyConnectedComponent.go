package tarjan

import "edaa/internals/interfaces"

type SCC struct {
	nodes      []interfaces.Node
	hasStation bool
}

func (S *SCC) Nodes() []interfaces.Node {
	return S.nodes
}

func (S *SCC) SetNodes(nodes []interfaces.Node) {
	S.nodes = nodes
}

func (S *SCC) HasStation() bool {
	return S.hasStation
}

func (S *SCC) SetHasStation(hasStation bool) {
	S.hasStation = hasStation
}

func NewSCC(nodes []interfaces.Node, hasStation bool) *SCC {
	return &SCC{nodes: nodes, hasStation: hasStation}
}
