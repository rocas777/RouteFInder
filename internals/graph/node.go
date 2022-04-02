package graph

type Node struct {
	Edges         []*Edge
	IncomingEdges []*Edge
	latitude      float64
	longitude     float64
	name          string
	zone          string
	Code          string
	Referenced    bool
	IsStation     bool
	Visited		  bool
	Distance	  float64
	Previous	  string
}

func NewStationNode(latitude float64, longitude float64, name string, zone string, code string) *Node {
	return &Node{latitude: latitude, longitude: longitude, name: name, zone: zone, Code: code, IsStation: true, Referenced: false}
}

func NewNormalNode(latitude float64, longitude float64, name string, zone string, code string) *Node {
	return &Node{latitude: latitude, longitude: longitude, name: name, zone: zone, Code: code, IsStation: false}
}

func (n *Node) AddDestination(destination *Node, weight float64) {
	destination.Referenced = true
	n.Referenced = true
	n.Edges = append(n.Edges, NewEdge(n, destination, weight))
	n.IncomingEdges = append(n.IncomingEdges, NewEdge(destination, n, weight))
}

func (n *Node) RemoveEdge(node *Node) {
	for i, edge := range n.Edges {
		if edge.To.Code == node.Code {
			n.Edges = append(n.Edges[:i], n.Edges[i+1:]...)
			return
		}
	}
}
func (n *Node) RemoveIncomingEdge(node *Node) {
	for i, edge := range n.IncomingEdges {
		if edge.To.Code == node.Code {
			n.Edges = append(n.IncomingEdges[:i], n.IncomingEdges[i+1:]...)
			return
		}
	}
}
