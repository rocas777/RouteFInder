package quadtree

import "edaa/internals/interfaces"

type quad struct {
	nw *quad
	ne *quad
	sw *quad
	se *quad

	nodes []interfaces.Node

	tlLat float64
	tlLon float64
	brLat float64
	brLon float64

	start int
	end   int
}

func (q* quad) NW() interfaces.Quad {
	return q.nw
}

func (q* quad) NE() interfaces.Quad {
	return q.ne
}

func (q* quad) SW() interfaces.Quad {
	return q.sw
}

func (q* quad) SE() interfaces.Quad {
	return q.se
}

func (q* quad) Nodes() []interfaces.Node {
	return q.nodes
}

func (q* quad) TlLat() float64 {
	return q.tlLat
}

func (q* quad) TlLon() float64 {
	return q.tlLon
}

func (q* quad) BrLat() float64 {
	return q.brLat
}

func (q* quad) BrLon() float64 {
	return q.brLon
}


func (q* quad) GetNodesPos() (int,int) {
	return q.start,q.end
}

func newQuad(tlLat float64, tlLon float64, brLat float64, brLon float64) *quad {
	return &quad{tlLat: tlLat, tlLon: tlLon, brLat: brLat, brLon: brLon}
}



type QuadTree struct {
	Root *quad
}
