package visualization

import (
	"edaa/internals/interfaces"
	"edaa/internals/visualization/primitives"
	"fmt"
	"github.com/fogleman/gg"
	"time"
)

func DrawQuad(quad interfaces.Quad) {
	nodes := quad.Nodes()

	start := time.Now()
	ctx := gg.NewContext(1000, 1000)

	xStandardizer := (quad.BrLon() - quad.TlLon()) * 1000
	yStandardizer := (quad.BrLat() - quad.TlLat()) * 1000

	for _, n := range nodes {
		//println(n.Longitude()-quad.TlLon(), xStandardizer, (n.Longitude()-quad.TlLon())*xStandardizer)
		//println(n.Latitude()-quad.TlLat(), yStandardizer, (n.Latitude()-quad.TlLat())*yStandardizer)
		//println()

		x := (n.Longitude() - quad.TlLon()) * xStandardizer
		y := (n.Latitude() - quad.TlLat()) * yStandardizer
		primitives.DrawNode(n, ctx, x, y)
		ctx.SetRGB(0, 0, 0)
		ctx.Fill()

		for _, e := range n.OutEdges() {
			primitives.DrawEdge(e, ctx, quad)
		}

		ctx.SetRGB(1, 0, 0)
		ctx.Fill()
	}

	ctx.SavePNG("out.png")

	fmt.Println(time.Since(start).Milliseconds())
}
