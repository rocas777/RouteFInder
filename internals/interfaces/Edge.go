package interfaces

import "edaa/internals/types"

type Edge interface {
	From() Node
	SetFrom(from Node)

	To() Node
	SetTo(to Node)

	Weight() float64
	SetWeight(weight float64)

	EdgeType() types.EdgeType
}
