package graph

import (
	"edaa/internals/interfaces"
	"edaa/internals/types"
)

type Node struct {
	outEdges   []interfaces.Edge
	inEdges    []interfaces.Edge
	latitude   float64
	longitude  float64
	name       string
	zone       string
	id         string
	referenced bool
	isStation  bool
	fromLandmarks [12]float64
}

func (n *Node) OutEdges() []interfaces.Edge {
	return n.outEdges
}

func (n *Node) SetOutEdges(outEdges []interfaces.Edge) {
	n.outEdges = outEdges
}

func (n *Node) InEdges() []interfaces.Edge {
	return n.inEdges
}

func (n *Node) SetInEdges(inEdges []interfaces.Edge) {
	n.inEdges = inEdges
}

func (n *Node) Latitude() float64 {
	return n.latitude
}

func (n *Node) SetLatitude(latitude float64) {
	n.latitude = latitude
}

func (n *Node) Longitude() float64 {
	return n.longitude
}

func (n *Node) SetLongitude(longitude float64) {
	n.longitude = longitude
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) SetName(name string) {
	n.name = name
}

func (n *Node) Zone() string {
	return n.zone
}

func (n *Node) SetZone(zone string) {
	n.zone = zone
}

func (n *Node) Id() string {
	return n.id
}

func (n *Node) SetId(id string) {
	n.id = id
}

func (n *Node) Referenced() bool {
	return n.referenced
}

func (n *Node) SetReferenced(referenced bool) {
	n.referenced = referenced
}

func (n *Node) IsStation() bool {
	return n.isStation
}

func (n *Node) SetIsStation(isStation bool) {
	n.isStation = isStation
}

func NewStationNode(latitude float64, longitude float64, name string, zone string, code string) *Node {
	return &Node{latitude: latitude, longitude: longitude, name: name, zone: zone, id: code, isStation: true, referenced: false}
}

func NewNormalNode(latitude float64, longitude float64, name string, zone string, code string) *Node {
	return &Node{latitude: latitude, longitude: longitude, name: name, zone: zone, id: code, isStation: false, referenced: false}
}

func (n *Node) AddDestination(destination interfaces.Node, weight float64) {
	edgeType := types.Road

	if n.Id()[0] == destination.Id()[0] {
		if n.Id()[0] == 'm' {
			edgeType = types.Metro
		} else if n.Id()[0] == 'b' {
			edgeType = types.Bus
		}
	}

	if weight == 0 {
		println(n.id, destination.Id())
		panic("")
	}
	destination.SetReferenced(true)
	n.referenced = true
	n.outEdges = append(n.outEdges, NewEdge(n, destination, weight, edgeType))
	destination.SetInEdges(append(destination.InEdges(), NewEdge(n, destination, weight, edgeType)))
}

func (n *Node) RemoveOutEdge(node interfaces.Node) {
	for i, edge := range n.outEdges {
		if edge.To().Id() == node.Id() {
			n.outEdges = append(n.outEdges[:i], n.outEdges[i+1:]...)
			return
		}
	}
}
func (n *Node) RemoveInEdge(node interfaces.Node) {
	for i, edge := range n.inEdges {
		if edge.To().Id() == node.Id() {
			n.inEdges = append(n.inEdges[:i], n.inEdges[i+1:]...)
			return
		}
	}
}

func (n *Node) RemoveConnections(nodeToRemove interfaces.Node) {
	n.RemoveInEdge(nodeToRemove)
	n.RemoveOutEdge(nodeToRemove)
}

func (n *Node) AddFromLandmark(landmark int, distance float64) () {
	n.fromLandmarks[landmark] = distance
}

func (n *Node) GetFromLandmarks() [12]float64 {
	return n.fromLandmarks
}
