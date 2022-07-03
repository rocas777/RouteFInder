package main

import (
	"edaa/internals/algorithms/path/landmarks"
	kdtree2 "edaa/internals/dataStructures/kdtree"
	quadtree2 "edaa/internals/dataStructures/quadtree"
	"edaa/internals/exports/reuse"
	"edaa/internals/graph"
	"edaa/internals/types"
	tile_server "edaa/internals/visualization/tile-server"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"os/exec"
	"strconv"
	time2 "time"

	//"edaa/internals/algorithms/path/astar"
	//"edaa/internals/g"
	"edaa/internals/interfaces"
	"edaa/internals/menuHelper"

	//"edaa/internals/utils"
	//"time"
	"bufio"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
)

var g interfaces.Graph
var kdtree *kdtree2.KDTree
var quadtree *quadtree2.QuadTree
var l *landmarks.Dijkstra

var mLat = 999.0
var MLat = -999.0
var mLon = 999.0
var MLon = -999.0

var first = 0

var kill chan interface{}

func server (){
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		sLatS := c.Request.URL.Query()["slat"][0]
		sLonS := c.Request.URL.Query()["slon"][0]
		dLatS := c.Request.URL.Query()["dlat"][0]
		dLonS := c.Request.URL.Query()["dlon"][0]
		m := c.Request.URL.Query()["method"][0]

		sLat,_ := strconv.ParseFloat(sLatS,64)
		sLon,_ := strconv.ParseFloat(sLonS,64)
		dLat,_ := strconv.ParseFloat(dLatS,64)
		dLon,_ := strconv.ParseFloat(dLonS,64)

		top := 38.194482
		bottom := 37.95421
		topR := 41.37813
		bottomR := 41.08061

		pS := (sLat - bottom) / (top - bottom)
		pD := (dLat - bottom) / (top - bottom)

		sLat = bottomR + (topR-bottomR) * pS
		dLat = bottomR + (topR-bottomR) * pD


		var time float64
		var cost float64
		var path []interfaces.Edge

		start := time2.Now()

		switch m{
		case "d":
			time,cost,path = menuHelper.DijkstraServer(g,kdtree,sLat,sLon,dLat,dLon)
			fmt.Println("Price:",cost)
			pathTime := time
			fmt.Printf("Time: %d:%d\n", int(pathTime/60), int((pathTime/60-math.Floor(pathTime/60))*60))
			fmt.Println("Alg Time:",time2.Since(start).Milliseconds())
		case "a":
			time,cost,path = menuHelper.AStartServer(g,kdtree,sLat,sLon,dLat,dLon)
			t := time
			time = cost
			cost = t
			fmt.Println("Price:",cost)
			pathTime := time
			fmt.Printf("Time: %d:%d\n", int(pathTime/60), int((pathTime/60-math.Floor(pathTime/60))*60))
			fmt.Println("Alg Time:",time2.Since(start).Milliseconds())
		case "alt":
			time,cost,path = menuHelper.ALTServer(g,kdtree,sLat,sLon,dLat,dLon,l)
			t := time
			time = cost
			cost = t
			fmt.Println("Price:",cost)
			pathTime := time
			fmt.Printf("Time: %d:%d\n", int(pathTime/60), int((pathTime/60-math.Floor(pathTime/60))*60))
			fmt.Println("Alg Time:",time2.Since(start).Milliseconds())
		case "gt":
			time,cost,path = menuHelper.GeneticTimeServer(g,kdtree,sLat,sLon,dLat,dLon)
			t := time
			time = cost
			cost = t
			fmt.Println("Price:",cost,"€")
			pathTime := time
			fmt.Printf("Time: %d:%d\n", int(pathTime/60), int((pathTime/60-math.Floor(pathTime/60))*60))
			fmt.Println("Alg Time:",time2.Since(start).Milliseconds(),"ms")
		case "gp":
			time,cost,path = menuHelper.GeneticPriceServer(g,kdtree,sLat,sLon,dLat,dLon)
			t := time
			time = cost
			cost = t
			fmt.Println("Price:",cost)
			pathTime := time
			fmt.Printf("Time: %d:%d\n", int(pathTime/60), int((pathTime/60-math.Floor(pathTime/60))*60))
			fmt.Println("Alg Time:",time2.Since(start).Milliseconds())
		}

		GetLegs(path)

		c.JSON(http.StatusOK, gin.H{"price":cost,"time":fmt.Sprintf("Time: %d:%d\n", int(time/60), int((time/60-math.Floor(time/60))*60)),"alg_time":time2.Since(start).Milliseconds()})

	})
	r.Run(":8081")
}

type leg struct {
	start string
	end string
	method string
}

func GetLegs(path []interfaces.Edge)  {
	out := []leg{}
	var last types.EdgeType = types.Road
	var lastNode string = ""
	for _, edge := range path {
		if last != edge.EdgeType(){

			edgeT := ""
			if edge.EdgeType() == types.Road{
				edgeT = "Walk"
			}else if edge.EdgeType() == types.Metro{
				edgeT = "Metro"
			}else if edge.EdgeType() == types.Bus{
				edgeT = "Bus"
			}

			out = append(out, leg{
				start:  lastNode,
				end:    edge.To().Name(),
				method: edgeT,
			})
			last = edge.EdgeType()
			lastNode = edge.To().Name()
		}
	}
	fmt.Println(out)
}

func main() {
	go server()
	kill = make(chan interface{})
	for {
		fmt.Println("")
		fmt.Println("Select one of the following")
		fmt.Println("1 - Setup")
		fmt.Println("2 - Load")
		fmt.Println("3 - Load Landmarks")
		fmt.Println("4 - Connectivity Analysis")
		fmt.Println("5 - Find Path")
		fmt.Println("6 - Find Path Landmarks")
		fmt.Println("7 - Find Path Genetics")
		fmt.Println("8 - See Map")
		fmt.Println("9 - Load Raw")
		fmt.Println("10 - Export")
		fmt.Println("11 - Exit")
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
				fmt.Println("Could not load, file does not exist")
			}
			for _, node := range g.Nodes() {
				if node.Latitude() < mLat{
					mLat = node.Latitude()
				}
				if node.Latitude() > MLat{
					MLat = node.Latitude()
				}
				if node.Longitude() < mLon{
					mLon = node.Latitude()
				}
				if node.Longitude() > MLon{
					MLon = node.Latitude()
				}
			}
			if quadtree == nil {
				quadtree = quadtree2.NewQuadTree(g)
				go restartServer()
			}
		case "3":
			if g == nil {
				fmt.Println("Must setup or load graph first!!!!")
			} else {
				l = menuHelper.Landmarks(g)
			}
		case "4":
			if g == nil {
				fmt.Println("Must setup or load graph first!!!!")
			} else {
				menuHelper.Connectivity(g)
			}
		case "5":
			if _, err := os.Stat("data/reuse/edges.csv"); err == nil {
				g = &graph.Graph{}
				graph.InitReuse(g)
			} else if errors.Is(err, os.ErrNotExist) {
				fmt.Println("Could not load, file does not exist")
			}
			for _, node := range g.Nodes() {
				if node.Latitude() < mLat{
					mLat = node.Latitude()
				}
				if node.Latitude() > MLat{
					MLat = node.Latitude()
				}
				if node.Longitude() < mLon{
					mLon = node.Latitude()
				}
				if node.Longitude() > MLon{
					MLon = node.Latitude()
				}
			}
			if quadtree == nil {
				quadtree = quadtree2.NewQuadTree(g)
				go restartServer()
			}
			if g == nil {
				fmt.Println("Must setup or load graph first!!!!")
			} else {
				if g == nil {
					fmt.Println("Must setup or load graph first!!!!")
				} else {
					if kdtree == nil {
						kdtree = kdtree2.NewKDTree(g)
					}
					if quadtree == nil {
						quadtree = quadtree2.NewQuadTree(g)
						go restartServer()
					}
					menuHelper.PathFinder(g, kdtree)
				}
			}
		case "6":
			if g == nil {
				fmt.Println("Must setup or load graph first!!!!")
			} else {
				if l == nil {
					fmt.Println("Must load landmarks first!!!!")
				} else {
					if kdtree == nil {
						kdtree = kdtree2.NewKDTree(g)
					}
					if quadtree == nil {
						quadtree = quadtree2.NewQuadTree(g)
						go restartServer()
					}
					menuHelper.PathFinderLandmarks(g, kdtree, l)
				}
			}
		case "7":
			if g == nil {
				fmt.Println("Must setup or load graph first!!!!")
			} else {
				if g == nil {
					fmt.Println("Must setup or load graph first!!!!")
				} else {
					if kdtree == nil {
						kdtree = kdtree2.NewKDTree(g)
					}
					if quadtree == nil {
						quadtree = quadtree2.NewQuadTree(g)
						go restartServer()
					}
					menuHelper.PathFinderGenetics(g, kdtree)
				}
			}
		case "8":
			go func() {
				fmt.Println("Loading... be patient")
				cmd := exec.Command("python3", "networkx/view.py")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				log.Println(cmd.Run())
			}()
		case "9":
			g = &graph.Graph{}
			fmt.Println("")
			fmt.Println("Initiating graph...")
			g.Init()
		case "10":
			reuse.ExportEdges(g, "data/reuse/edges.csv")
			reuse.ExportNodes(g, "data/reuse/nodes.csv")
		case "11":
			fmt.Println("Cya")
			return
		}
	}
}

func restartServer(){
	if first != 0 {
		kill <- ""
	}
	first = 1
	tile_server.TileServer(quadtree.Root,kill)
}