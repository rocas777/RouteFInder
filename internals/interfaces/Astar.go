package interfaces

type AStar interface {
	Path(source Node, destination Node) []Node
	Dist(source Node) []struct {
		node Node
		dist float64
	}
}
