package main

import (
	"context"
	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
	"os"
	"runtime"
)

func main() {
	file, err := os.Open("data/road/portugal-latest.pbf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := osmpbf.New(context.Background(), file, runtime.GOMAXPROCS(-1))
	defer scanner.Close()

	for scanner.Scan() {
		switch o := scanner.Object().(type) {
		case *osm.Node:
			println("dd", o)
		case *osm.Way:
			println("dd", o)
		case *osm.Relation:
			println("dd", o)
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
