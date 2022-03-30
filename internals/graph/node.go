package graph

type Node struct {
	Edges     []*Edge
	latitude  float64
	longitude float64
	isStation bool
	name      string
	zone      string
	Code      string
}

func NewStationNode(latitude float64, longitude float64, name string, zone string, code string) *Node {
	return &Node{latitude: latitude, longitude: longitude, name: name, zone: zone, Code: code, isStation: true}
}

func NewNormalNode(latitude float64, longitude float64, name string, zone string, code string) *Node {
	return &Node{latitude: latitude, longitude: longitude, name: name, zone: zone, Code: code, isStation: false}
}

func (n *Node) AddDestination(destination *Node, weight float64) {
	n.Edges = append(n.Edges, NewEdge(n, destination, weight))
}
