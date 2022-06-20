package astar

import (
	"edaa/internals/interfaces"
	"edaa/internals/types"
)

type astarEdge struct {
	from     interfaces.Node
	to       interfaces.Node
	edgeType types.EdgeType
	weight   float64
}

func (a astarEdge) From() interfaces.Node {
	return a.from
}

func (a astarEdge) SetFrom(from interfaces.Node) {
	panic("implement me")
}

func (a astarEdge) To() interfaces.Node {
	return a.to
}

func (a astarEdge) SetTo(to interfaces.Node) {
	panic("implement me")
}

func (a astarEdge) Weight() float64 {
	return a.weight
}

func (a astarEdge) SetWeight(weight float64) {
	panic("implement me")
}

func (a astarEdge) EdgeType() types.EdgeType {
	return a.edgeType
}
