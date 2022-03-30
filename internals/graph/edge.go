package graph

type Edge struct {
	From   *Node
	To     *Node
	Weight float64
}

func NewEdge(from *Node, to *Node, weight float64) *Edge {
	return &Edge{From: from, To: to, Weight: weight}
}
