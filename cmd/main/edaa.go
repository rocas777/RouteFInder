package main

import (
	kdtree2 "edaa/internals/dataStructures/kdtree"
	"edaa/internals/exports/reuse"
	"edaa/internals/graph"
	"errors"
	"log"
	"os/exec"

	//"edaa/internals/algorithms/path/astar"
	//"edaa/internals/g"
	"edaa/internals/interfaces"
	"edaa/internals/menuHelper"
	//"edaa/internals/utils"
	//"time"
	"bufio"
	"os"
	"strings"
)

var g interfaces.Graph
var kdtree *kdtree2.KDTree

func menu() bool {
	println("")
	println("Select one of the following")
	println("1 - Setup")
	println("2 - Load")
	println("3 - Connectivity Analysis")
	println("4 - Find Path")
	println("5 - See Map")
	println("6 - Load Raw")
	println("7 - Export")
	println("8 - Exit")
	reader := bufio.NewReader(os.Stdin)
	opt, _ := reader.ReadString('\n')
	opt = strings.TrimSpace(opt)

	switch opt {
	case "1":
		g = &graph.Graph{}
		g = menuHelper.Setup()
	case "2":
		if _, err := os.Stat("data/reuse/edges.csv"); err == nil {
			g = &graph.Graph{}
			graph.InitReuse(g)
		} else if errors.Is(err, os.ErrNotExist) {
			println("Could not load, file does not exist")
		}
	case "3":
		if g == nil {
			println("Must setup or load graph first!!!!")
		} else {
			menuHelper.Connectivity(g)
		}
	case "4":
		if g == nil {
			println("Must setup or load graph first!!!!")
		} else {
			if g == nil {
				println("Must setup or load graph first!!!!")
			} else {
				if kdtree == nil {
					kdtree = kdtree2.NewKDTree(g)
				}
				menuHelper.PathFinder(g, kdtree)
			}
		}
	case "5":
		go func() {
			println("Loading... be patient")
			cmd := exec.Command("python3", "networkx/view.py")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			log.Println(cmd.Run())
		}()
	case "6":
		g = &graph.Graph{}
		println("")
		println("Initiating graph...")
		g.Init()
	case "7":
		reuse.ExportEdges(g, "data/reuse/edges.csv")
		reuse.ExportNodes(g, "data/reuse/nodes.csv")
	case "8":
		println("Cya")
		return false
	}
	return true
}

func main() {
	for menu() {
	}
}
