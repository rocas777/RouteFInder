package graph

import (
	"edaa/internals/interfaces"
	"edaa/internals/types"
)

type Edge struct {
	from     interfaces.Node
	to       interfaces.Node
	weight   float64
	edgeType types.EdgeType
}

func (e *Edge) EdgeType() types.EdgeType {
	return e.edgeType
}

func (e *Edge) From() interfaces.Node {
	return e.from
}

func (e *Edge) SetFrom(from interfaces.Node) {
	e.from = from
}

func (e *Edge) To() interfaces.Node {
	return e.to
}

func (e *Edge) SetTo(to interfaces.Node) {
	e.to = to
}

func (e *Edge) Weight() float64 {
	return e.weight
}

func (e *Edge) SetWeight(weight float64) {
	e.weight = weight
}

func NewEdge(from interfaces.Node, to interfaces.Node, weight float64, edgeType types.EdgeType) *Edge {
	return &Edge{from: from, to: to, weight: weight, edgeType: edgeType}
}
