package graph

type Node struct {
	Edges         []*Edge
	IncomingEdges []*Edge
	Latitude      float64
	Longitude     float64
	name          string
	zone          string
	Code          string
	Referenced    bool
	IsStation     bool
}

func (n *Node) Lat() float64 {
	return n.Latitude
}

func (n *Node) Lon() float64 {
	return n.Longitude
}

func NewStationNode(latitude float64, longitude float64, name string, zone string, code string) *Node {
	return &Node{Latitude: latitude, Longitude: longitude, name: name, zone: zone, Code: code, IsStation: true, Referenced: false}
}

func NewNormalNode(latitude float64, longitude float64, name string, zone string, code string) *Node {
	return &Node{Latitude: latitude, Longitude: longitude, name: name, zone: zone, Code: code, IsStation: false}
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
