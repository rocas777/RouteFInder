package tile_server

import (
	"edaa/internals/interfaces"
	"edaa/internals/visualization"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
)

var path []interfaces.Edge

var killer = make(chan interface{})

func AddPath(edges []interfaces.Edge) {
	os.RemoveAll("images/")
	os.MkdirAll("images/", os.ModePerm)
	path = edges
}

func ClearPath() {
	os.RemoveAll("images/")
	os.MkdirAll("images/", os.ModePerm)
	path = []interfaces.Edge{}
}

func TileServer(q interfaces.Quad, kill chan interface{}) {
	killer = kill
	rand.Seed(time.Now().UnixNano())

	adder := strconv.Itoa(rand.Int())

	ClearPath()
	r := mux.NewRouter()
	r.HandleFunc("/{z}/{x}/{y}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		xs := vars["x"]
		x, _ := strconv.ParseInt(xs, 10, 32)

		ys := vars["y"]
		y, _ := strconv.ParseInt(ys[:len(ys)-4], 10, 32)

		zs := vars["z"]
		z, _ := strconv.ParseInt(zs, 10, 64)

		tiles := int64(math.Sqrt(math.Pow(4, float64(z))))
		var startX int64 = 0
		var endX = tiles
		var startY int64 = 0
		var endY = tiles
		var i int64 = 0
		exploringNode := q
		imgN := ""
		for ; i < z; i++ {
			if reflect.ValueOf(exploringNode).IsNil() {
				return
			}
			if x < (endX+startX)/2 {
				if y < (endY+startY)/2 {
					exploringNode = exploringNode.NW()
					endX = (endX + startX) / 2
					endY = (endY + startY) / 2
					imgN += "NW-"
				} else {
					exploringNode = exploringNode.SW()
					endX = (endX + startX) / 2
					startY = (endY + startY) / 2
					imgN += "SW-"
				}
			} else {
				if y < (endY+startY)/2 {
					exploringNode = exploringNode.NE()
					startX = (endX + startX) / 2
					endY = (endY + startY) / 2
					imgN += "NE-"
				} else {
					exploringNode = exploringNode.SE()
					startX = (endX + startX) / 2
					startY = (endY + startY) / 2
					imgN += "SE-"
				}
			}
		}

		if reflect.ValueOf(exploringNode).IsNil() {
			return
		}

		a, b := exploringNode.GetNodesPos()
		imgN += strconv.Itoa(b-a) + "-" + adder

		if _, err := os.Stat("images/" + imgN + ".png"); errors.Is(err, os.ErrNotExist) {
			visualization.DrawQuad(exploringNode, q, imgN, path)
		}

		img, err := os.Open("images/" + imgN + ".png")
		if err != nil {
			log.Fatal(err) // perhaps handle this nicer
		}
		defer img.Close()
		w.Header().Set("Content-Type", "image/png") // <-- set the content-type header
		io.Copy(w, img)

	})
	if kill != nil {
		go http.ListenAndServe("localhost:8000", r)
		<-kill
	} else {
		http.ListenAndServe("localhost:8000", r)
	}
}
