package interfaces

type Node interface {
	OutEdges() []Edge
	SetOutEdges(outEdges []Edge)

	InEdges() []Edge
	SetInEdges(inEdges []Edge)

	Latitude() float64
	SetLatitude(latitude float64)

	Longitude() float64
	SetLongitude(longitude float64)

	Name() string
	SetName(name string)

	Zone() string
	SetZone(name string)

	Id() string
	SetId(name string)

	Referenced() bool
	SetReferenced(referenced bool)

	IsStation() bool
	SetIsStation(isStation bool)
	
	Visited() bool
	SetVisited(visited bool)
	
	Distance() float64
	SetDistance(distance float64)
	
	Previous() string
	SetPrevious(previous string)

	AddDestination(destination Node, weight float64)
	RemoveConnections(nodeToRemove Node)
}
