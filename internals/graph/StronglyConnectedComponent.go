package graph

type SCC struct {
	Nodes      []*Node
	HasStation bool
}

func NewSCC(nodes []*Node, hasStation bool) *SCC {
	return &SCC{Nodes: nodes, HasStation: hasStation}
}
