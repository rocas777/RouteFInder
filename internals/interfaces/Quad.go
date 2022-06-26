package interfaces

type Quad interface {
	NW() Quad
	NE() Quad
	SW() Quad
	SE() Quad

	Nodes() []Node
	GetNodesPos() (int,int)

	TlLat() float64
	TlLon() float64
	BrLat() float64
	BrLon() float64
}
