package quadtree

import (
	"edaa/internals/interfaces"
	"fmt"
	"math"
	"runtime"
)

func NewQuadTree(g interfaces.Graph) *QuadTree {
	PrintMemUsage()

	top_left_lat := -math.Inf(1)
	top_left_lon := math.Inf(1)

	bottom_right_lat := math.Inf(1)
	bottom_right_lon := -math.Inf(1)

	for _, n := range g.Nodes() {
		if top_left_lat < n.Latitude() {
			top_left_lat = n.Latitude()
		}
		if top_left_lon > n.Longitude() {
			top_left_lon = n.Longitude()
		}
		if bottom_right_lat > n.Latitude() {
			bottom_right_lat = n.Latitude()
		}
		if bottom_right_lon < n.Longitude() {
			bottom_right_lon = n.Longitude()
		}
	}

	//fmt.Println(top_left_lat, top_left_lon)
	//fmt.Println(bottom_right_lat, bottom_right_lon)

	q := &quad{
		nodes: append([]interfaces.Node{}, g.Nodes()...),
		tlLat: top_left_lat,
		tlLon: top_left_lon,
		brLat: bottom_right_lat,
		brLon: bottom_right_lon,
	}

	_,q.nodes = sortInQuad(g, q, 0, 0)

	PrintMemUsage()

	return &QuadTree{Root: q}
}

var i = 0

func sortInQuad(g interfaces.Graph, q *quad, depth int, init int) (int,[]interfaces.Node) {
	if q.nodes == nil {
		q.start = init
		q.end = init
		return 0,[]interfaces.Node{}
	}
	if len(q.nodes) == 1 {
		n := q.nodes[0]

		q.start = init
		q.end = init+1

		q.nodes = nil
		return 1,[]interfaces.Node{n}
	}

	//fmt.Println(depth, q.tlLat-q.brLat, q.tlLon-q.brLon)
	mLat := (q.brLat + q.tlLat) / 2.0
	mLon := (q.brLon + q.tlLon) / 2.0

	q.nw = newQuad(q.tlLat, q.tlLon, mLat, mLon)
	q.ne = newQuad(q.tlLat, mLon, mLat, q.brLon)

	q.sw = newQuad(mLat, q.tlLon, q.brLat, mLon)
	q.se = newQuad(mLat, mLon, q.brLat, q.brLon)

	for _, n := range q.nodes {
		if isNw(n.Latitude(), n.Longitude(), mLat, mLon) {
			q.nw.nodes = append(q.nw.nodes, n)
		} else if isNe(n.Latitude(), n.Longitude(), mLat, mLon) {
			q.ne.nodes = append(q.ne.nodes, n)
		} else if isSw(n.Latitude(), n.Longitude(), mLat, mLon) {
			q.sw.nodes = append(q.sw.nodes, n)
		} else if isSe(n.Latitude(), n.Longitude(), mLat, mLon) {
			q.se.nodes = append(q.se.nodes, n)
		}
	}


	if depth < 18 {
		q.nodes = nil

		adder1,ln1 := sortInQuad(g, q.nw, depth+1, init)
		adder2,ln2 := sortInQuad(g, q.ne, depth+1, init+adder1)
		adder3,ln3 := sortInQuad(g, q.sw, depth+1, init+adder1+adder2)
		adder4,ln4 := sortInQuad(g, q.se, depth+1, init+adder1+adder2+adder3)

		q.start = init
		q.end = init + adder1+adder2+adder3+adder4

		n := append(append(append(ln1, ln2...), ln3...), ln4...)

		return adder1+adder2+adder3+adder4,n
	}else{

		q.start = init
		q.end = init + len(q.nodes)

		n := make([]interfaces.Node,len(q.nodes))
		copy(n,q.nodes)
		q.nodes = nil
		return len(q.nodes),n
	}
}

func isNw(lat, lon, mLat, mLon float64) bool {
	return lat > mLat && lon < mLon
}

func isNe(lat, lon, mLat, mLon float64) bool {
	return lat > mLat && lon >= mLon
}

func isSw(lat, lon, mLat, mLon float64) bool {
	return lat < mLat && lon < mLon
}

func isSe(lat, lon, mLat, mLon float64) bool {
	return lat < mLat && lon >= mLon
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

