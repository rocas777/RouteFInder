package genetics

import (
	"edaa/internals/algorithms/path/astar"
	kdtree2 "edaa/internals/dataStructures/kdtree"
	"edaa/internals/graph"
	"edaa/internals/interfaces"
	"edaa/internals/types"
	"edaa/internals/utils"
	"math"
	"math/rand"
	"os"
	"time"
)

type solution struct {
	path   []interfaces.Edge
	weight float64
}
func GeneticPath(g interfaces.Graph, start interfaces.Node, end interfaces.Node, kdtree *kdtree2.KDTree,time bool) ([]interfaces.Edge, float64, int) {
	if time{
		p,t,i :=  _geneticPath(g, start, end , kdtree,costTime)
		return p,t,i
	}else{
		p,t,i :=  _geneticPath(g, start, end , kdtree,costCost)
		return p,t,i
	}
}


func _geneticPath(g interfaces.Graph, start interfaces.Node, end interfaces.Node, kdtree *kdtree2.KDTree,cost func([]interfaces.Edge) float64) ([]interfaces.Edge, float64, int) {
	//initTime := time.Now()

	bp, _, _ := GetBestSolution(g, start, end)
	a := cost(bp)
	middleStopLat := (start.Latitude() + end.Latitude()) / 2
	middleStopLon := (start.Longitude() + end.Longitude()) / 2

	firstStopLat := (start.Latitude() + middleStopLat) / 2
	firstStopLon := (start.Longitude() + middleStopLon) / 2

	thirdStopLat := (middleStopLat + end.Latitude()) / 2
	thirdStopLon := (middleStopLon + end.Longitude()) / 2

	fn, _ := kdtree.GetClosest(graph.NewNormalNode(firstStopLat, firstStopLon, "", "", ""))
	mn, _ := kdtree.GetClosest(graph.NewNormalNode(middleStopLat, middleStopLon, "", "", ""))
	tn, _ := kdtree.GetClosest(graph.NewNormalNode(thirdStopLat, thirdStopLon, "", "", ""))

	w := 10.0
	p1, _, _ := GetSolution(g, start, mn, w)
	p2, _, _ := GetSolution(g, mn, end, w)

	p11, _, _ := GetSolution(g, start, fn, w)
	p12, _, _ := GetSolution(g, fn, mn, w)

	p21, _, _ := GetSolution(g, mn, tn, w)
	p22, _, _ := GetSolution(g, tn, end, w)

	pt, _, _ := GetSolution(g, start, end, w)
	solt := solution{
		path:   pt,
		weight: cost(pt),
	}

	path1 := make([]interfaces.Edge,0)
	path1 = append(path1, p1...)
	path1 = append(path1,p2...)
	sol1 := solution{
		path:   path1,
		weight: cost(path1),
	}


	path2 := make([]interfaces.Edge,0)
	path2 = append(path2, p1...)
	path2 = append(path2,p21...)
	path2 = append(path2,p22...)
	sol2 := solution{
		path:   path2,
		weight: cost(path2),
	}

	path3 := make([]interfaces.Edge,0)
	path3 = append(path3, p11...)
	path3 = append(path3,p12...)
	path3 = append(path3,p2...)
	sol3 := solution{
		path:   path3,
		weight: cost(path3),
	}

	path4 := make([]interfaces.Edge,0)
	path4 = append(path4, p11...)
	path4 = append(path4,p12...)
	path4 = append(path4,p21...)
	path4 = append(path4,p22...)
	sol4 := solution{
		path:   path4,
		weight: cost(path4),
	}
	rand.Seed(time.Now().UnixNano())

	sol := []solution{
		sol1, sol2, sol3, sol4,solt,
	}
	var bv float64 = 0
	wait := 10
	var last []float64
	for {
		last = append(last, bv)
		if len(last) >= wait {
			last = last[1:]
			l := last[0]
			finished := true
			for _, v := range last {
				if v != l {
					finished = false
				}
			}
			if finished {
				break
			}
		}
		sol, bv = geneticPass(g, sol, bv,cost)
		for _, s := range sol {
			if s.path[len(s.path)-1].To().Id() != end.Id(){
				panic("QUIIII!")
			}
			/*if a > bv{
				panic("ooo")
			}*/
		}
	}

	//fmt.Println()
	//fmt.Println("Genetic:")
	//fmt.Println("Best path:", bv)
	//fmt.Println("dest:", sol[0].path[len(sol[0].path)-1].To().Id())
	//fmt.Println(time.Since(initTime))

	//initTime = time.Now()

	//fmt.Println()
	//fmt.Println("Optimal:")
	//fmt.Println("Best Result:", a)
	//fmt.Println(start.Id(), end.Id())
	//fmt.Println(time.Since(initTime))

	b := math.Inf(1)
	pos := 0
	for i, s := range sol {
		c := cost(s.path)
		if cost(s.path) < b{
			pos = i
			b = c
		}
	}

	return sol[pos].path, bv, checkBitSetVar(sol[pos].path[len(sol[pos].path)-1].To().Id() == end.Id() && a <= bv)
}

func checkBitSetVar(mybool bool) int {
	if mybool {
		return 1
	}
	return 0 //you just saved youself an else here!
}

func getRandDiff(diff int, limit int) int {
	out := rand.Intn(limit)
	for out == diff {
		out = rand.Intn(limit)
	}
	return out
}

func geneticPass(g interfaces.Graph, solutions []solution, last float64,cost func([]interfaces.Edge) float64) ([]solution, float64) {
	var nextGen []solution
	var lastGen []solution
	var children []solution
	tournments := int(math.Floor(float64(len(solutions) / 2.0)))

	lastGen = solutions
	for i := 0; i < tournments; i++ {
		p1 := rand.Intn(len(lastGen))
		p2 := getRandDiff(p1, len(lastGen))
		if lastGen[p1].weight > lastGen[p2].weight {
			nextGen = append(nextGen, lastGen[p2])
		} else {
			nextGen = append(nextGen, lastGen[p1])
		}

		helper := lastGen
		lastGen = []solution{}
		for p := range helper {
			if p != p1 && p != p2 {
				lastGen = append(lastGen, solutions[p])
			}
		}
	}

	lastGen = solutions
	for i := 0; i < tournments; i++ {
		m := rand.Intn(len(lastGen))
		f := getRandDiff(m, len(lastGen))

		children = append(children, getChild(lastGen[m], lastGen[f],cost))

		helper := lastGen
		lastGen = []solution{}
		for p := range helper {
			if p != m && p != f {
				lastGen = append(lastGen, solutions[p])
			}
		}
	}
	nextGen = append(nextGen, children...)

	for _, s := range nextGen {
		if rand.Intn(50) <= 1 {
			mutate(g, &s,cost)
		}
	}

	best := math.Inf(1)
	for _, s := range nextGen {
		if s.weight < best {
			best = s.weight
		}
	}

	if best == last {
		toMutate := rand.Intn(len(nextGen))
		nt := nextGen[toMutate]
		mutate(g, &nt,cost)
		nextGen = append(nextGen, nt)
	}

	return nextGen, best
}
func getChild(m solution, f solution,cost func([]interfaces.Edge) float64) solution {
	bef := f.path[len(f.path)-1].To().Id()

	out := solution{
		path:   nil,
		weight: 0,
	}

	mNodes := make(map[string]int)
	for mP, mE := range m.path {
		mNodes[mE.To().Id()] = mP
	}

	for i := len(f.path) - 1; i >= 0; i-- {
		if v, ok := mNodes[f.path[i].From().Id()]; ok {
			out.path = make([]interfaces.Edge,0)
			out.path = append(out.path,m.path[:v+1]...)
			out.path = append(out.path,f.path[i:]...)
			break
		}
	}
	cur := out.path[len(out.path)-1].To().Id()

	if cur != bef {
		os.Exit(1)
	}

	out.weight = cost(out.path)
	return out
}

func costTime(path []interfaces.Edge) float64 {
	adder := 0.0
	for _, v := range path {
		adder += v.Weight()
	}
	return adder
}


var prices map[int]float64 = map[int]float64{
	2 : 1.25,
	3 : 1.6,
	4 : 2.0,
	5 : 2.4,
	6 : 2.85,
	7 : 3.25,
	8 : 3.65,
	9 : 4.05,
}

func costCost(path []interfaces.Edge) float64{
	cost,w,_ := costCostNoPenalty(path)
	if w >= 5*60{
		return cost + w / 60 * 10
	}
	return cost
}

func costCostNoPenalty(path []interfaces.Edge) (float64,float64,float64){
	walk := 0.0
	transport := 0.0
	cost := 0.0
	lastBus := true
	zones := make(map[string]interface{})
	for _, v := range path {
		if v.To().IsStation() {
			zones[v.To().Zone()] = "";
		}
		if lastBus && v.EdgeType() == types.Road{
			cost += 2.0
		}
		if v.EdgeType() != types.Road{
			lastBus = true;
			transport += v.Weight()
		}else{
			lastBus = false
			walk += v.Weight()
		}
	}
	/*fmt.Println(cost)
	fmt.Println(prices[len(zones)])
	fmt.Println(walk)
	fmt.Println(transport)
	fmt.Println()*/
	return math.Min(cost,prices[len(zones)]),walk,transport
}

func GetSolution(g interfaces.Graph, start interfaces.Node, end interfaces.Node, multiplier float64) (path []interfaces.Edge, t float64, explored int) {
	as := astar.NewAstar(g, func(from interfaces.Node, to interfaces.Node) float64 {
		a := utils.GetDistanceBetweenNodes(from, to) / (20)
		return a * multiplier
	})
	path, t, explored = as.Path(start, end)
	return
}

func GetBestSolution(g interfaces.Graph, start interfaces.Node, end interfaces.Node) (path []interfaces.Edge, t float64, explored int) {
	as := astar.NewAstar(g, func(from interfaces.Node, to interfaces.Node) float64 {
		a := utils.GetDistanceBetweenNodes(from, to) / (20)
		return a - a
	})
	path, t, explored = as.Path(start, end)
	return
}

func mutate(g interfaces.Graph, sol *solution,cost func([]interfaces.Edge) float64) {
	bf := sol.path[0].From().Id()
	bt := sol.path[len(sol.path)-1].To().Id()

	pi := len(sol.path)

	p1 := rand.Intn(len(sol.path))
	p2 := getRandDiff(p1, len(sol.path))

	if p1 > p2 {
		h := p1
		p1 = p2
		p2 = h
	}
	//p2 = int(math.Min(float64(p1+300), float64(p2)))
	if p1+1 == p2 {
		p2++
	}
	if p2 >= len(sol.path) {
		p2 = len(sol.path) - 1
	}
	w := float64(len(sol.path)) / (float64(len(sol.path)) - float64(p2-p1))
	w = w * 10
	nP, _, _ := GetSolution(g, sol.path[p1].To(), sol.path[p2].From(), w)

	if len(nP) == 0 {
		return
	}

	temp := make([]interfaces.Edge, len(sol.path))
	copy(temp, sol.path)

	sol.path = make([]interfaces.Edge,0)
	sol.path = append(sol.path,temp[:p1+1]...)
	sol.path = append(sol.path,nP...)
	sol.path = append(sol.path,temp[p2:]...)
	sol.weight = cost(sol.path)

	ef := sol.path[0].From().Id()
	et := sol.path[len(sol.path)-1].To().Id()

	if ef != bf || et != bt {
		println(pi)
		panic("")
	}
}
