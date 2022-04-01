package interfaces

type NeighbourFinder interface {
	GetClosest(target Node) (Node, float64)
}
