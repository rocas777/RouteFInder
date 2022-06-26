package genetics

import (
	"edaa/internals/algorithms/path/astar"
	kdtree2 "edaa/internals/dataStructures/kdtree"
	"edaa/internals/graph"
	"edaa/internals/interfaces"
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

func GeneticPath(g interfaces.Graph, start interfaces.Node, end interfaces.Node, kdtree *kdtree2.KDTree) ([]interfaces.Edge, float64, int) {
	//initTime := time.Now()

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
	p1, t1, _ := GetSolution(g, start, mn, w)
	p2, _, _ := GetSolution(g, mn, end, w)

	p11, _, _ := GetSolution(g, start, fn, w)
	p12, _, _ := GetSolution(g, fn, mn, w)

	p21, _, _ := GetSolution(g, mn, tn, w)
	p22, _, _ := GetSolution(g, tn, end, w)

	path1 := append(p1, p2...)
	sol1 := solution{
		path:   path1,
		weight: cost(path1),
	}

	path2 := append(append(p1, p21...), p22...)
	sol2 := solution{
		path:   path2,
		weight: cost(path2),
	}

	path3 := append(append(p11, p12...), p2...)
	sol3 := solution{
		path:   path3,
		weight: cost(path3),
	}

	path4 := append(append(append(p11, p12...), p21...), p22...)
	sol4 := solution{
		path:   path4,
		weight: cost(path4),
	}
	rand.Seed(time.Now().UnixNano())

	sol := []solution{
		sol1, sol2, sol3, sol4,
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
		sol, bv = geneticPass(g, sol, bv)
	}

	//fmt.Println()
	//fmt.Println("Genetic:")
	//fmt.Println("Best path:", bv)
	//fmt.Println("dest:", sol[0].path[len(sol[0].path)-1].To().Id())
	//fmt.Println(time.Since(initTime))

	//initTime = time.Now()
	bp, _, _ := GetBestSolution(g, start, end)

	//fmt.Println()
	//fmt.Println("Optimal:")
	//fmt.Println("Best Result:", cost(bp))
	//fmt.Println(start.Id(), end.Id())
	//fmt.Println(time.Since(initTime))

	return p1, t1, checkBitSetVar(sol[0].path[len(sol[0].path)-1].To().Id() == end.Id() && cost(bp) <= bv)
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

func geneticPass(g interfaces.Graph, solutions []solution, last float64) ([]solution, float64) {
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

		children = append(children, getChild(lastGen[m], lastGen[f]))

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
			mutate(g, &s)
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
		mutate(g, &nt)
		nextGen = append(nextGen, nt)
	}

	return nextGen, best
}
func getChild(m solution, f solution) solution {
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
			out.path = append(m.path[:v+1], f.path[i:]...)
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

func cost(path []interfaces.Edge) float64 {
	adder := 0.0
	for _, v := range path {
		adder += v.Weight()
	}
	return adder
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

func mutate(g interfaces.Graph, sol *solution) {
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

	temp := make([]interfaces.Edge, len(sol.path))
	copy(temp, sol.path)

	sol.path = append(append(temp[:p1+1], nP...), temp[p2:]...)
	sol.weight = cost(sol.path)
}
