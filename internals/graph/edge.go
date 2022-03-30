package graph

type Edge struct {
	from   *Node
	to     *Node
	weight float64
}

func NewEdge(from *Node, to *Node, weight float64) *Edge {
	return &Edge{from: from, to: to, weight: weight}
}
