package interfaces

type SCC interface {
	Nodes() []Node
	SetNodes(nodes []Node)
	HasStation() bool
	SetHasStation(hasStation bool)
}
