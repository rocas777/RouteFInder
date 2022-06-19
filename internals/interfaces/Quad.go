package interfaces

type Quad interface {
	NW() Quad
	NE() Quad
	SW() Quad
	SE() Quad

	Nodes() []Node

	TlLat() float64
	TlLon() float64
	BrLat() float64
	BrLon() float64
}
