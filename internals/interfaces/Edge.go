package interfaces

type Edge interface {
	From() Node
	SetFrom(from Node)

	To() Node
	SetTo(to Node)

	Weight() float64
	SetWeight(weight float64)
}
