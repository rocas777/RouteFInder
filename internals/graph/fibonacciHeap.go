package graph

import (
	"edaa/internals/utils"
	"edaa/internals/interfaces"
	"fmt"
)

type FibHeap struct {
	Roots			[]string
	NodesMap		map[string]*FibHeapNode
	minRoot			string
}

func NewFibHeap() *FibHeap {
	return &FibHeap{NodesMap: make(map[string]*FibHeapNode), minRoot: ""}
}

func (heap *FibHeap) AddNode(n interfaces.Node, key float64) *FibHeapNode {
	var heapNode *FibHeapNode = NewFibHeapNode(key, n.Id());
	
	heap.Roots = append(heap.Roots, heapNode.Code)
	heap.NodesMap[heapNode.Code] = heapNode;
	
	if (heap.minRoot == "" || key < heap.NodesMap[heap.minRoot].Key) {
		heap.minRoot = heapNode.Code;
	}
	return heapNode;
}

func (heap *FibHeap) PopMin() string {
	var min string = heap.minRoot;
	heap.Roots = utils.RemoveString(heap.Roots, heap.minRoot);
	
	//promote children to roots
	for _, c := range heap.NodesMap[min].Children {
		var n *FibHeapNode = heap.NodesMap[c];
		n.BecomeRoot()
		heap.Roots = append(heap.Roots, c)
	}
	
	//merge roots of same degree
	var i, j = heap.sameDegreeRoots();
	for i != "" {
		var nodeI *FibHeapNode = heap.NodesMap[i];
		var nodeJ *FibHeapNode = heap.NodesMap[j];
		
		if nodeI.Key < nodeJ.Key {
			nodeJ.BecomeChild(i);
			nodeI.AddChild(j, nodeJ.Degree);
			heap.Roots = utils.RemoveString(heap.Roots, j);
		} else {
			nodeI.BecomeChild(j);
			nodeJ.AddChild(i, nodeI.Degree);
			heap.Roots = utils.RemoveString(heap.Roots, i);
		}
		
		i, j = heap.sameDegreeRoots();
	}
	
	//update minRoot
	heap.minRoot = ""
	for _, r := range heap.Roots {
		if heap.minRoot == "" || heap.NodesMap[r].Key < heap.NodesMap[heap.minRoot].Key {
			heap.minRoot = r;
		}
	}
	
	delete(heap.NodesMap, min)
	return min;
}

func (heap *FibHeap) sameDegreeRoots() (string, string) {
	for i := 0; i < len(heap.Roots)-1; i++ {
		for j := i+1; j < len(heap.Roots); j++ {
			if heap.NodesMap[heap.Roots[i]].Degree == heap.NodesMap[heap.Roots[j]].Degree {
				return heap.Roots[i], heap.Roots[j];
			}
		}
	}
	
	return "", "";
}

func (heap *FibHeap) DecreaseKey(code string, newKey float64) {
	var n *FibHeapNode = heap.NodesMap[code];
	
	n.Key = newKey;
	
	if n.Parent != "" && newKey < heap.NodesMap[n.Parent].Key {
		for {
			var p *FibHeapNode = heap.NodesMap[n.Parent];
			
			//remove n from p children
			p.RemoveChild(n.Code);
			heap.updateDegree(p);
			
			//add n to roots, update minRoot
			heap.Roots = append(heap.Roots, n.Code);
			n.BecomeRoot()
			if n.Key < heap.NodesMap[heap.minRoot].Key {
				heap.minRoot = n.Code;
			}
			
			n = p;
			
			if !p.Loser {
				if !p.Root {
					p.Loser = true;
				}
				break;
			}
		}
	} else {
		if heap.minRoot == "" || n.Key < heap.NodesMap[heap.minRoot].Key {
			heap.minRoot = n.Code;
		}
	}
}

func (heap *FibHeap) updateDegree(n *FibHeapNode) {
	var d = 0;
	
	for _, c := range n.Children {
		if heap.NodesMap[c].Degree+1 > d {
			d = heap.NodesMap[c].Degree+1
		}
	}
	
	n.Degree = d;
}

func (heap *FibHeap) PrintFibHeap() {
	fmt.Println("minRoot:", heap.minRoot, "with k:", heap.NodesMap[heap.minRoot].Key, "all roots:", heap.Roots)
	for _, r := range heap.Roots {
		heap.NodesMap[r].PrintNode(heap)
	}
}