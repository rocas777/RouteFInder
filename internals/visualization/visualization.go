package visualization

import (
	"edaa/internals/interfaces"
	"edaa/internals/types"
	"edaa/internals/visualization/primitives"
	"github.com/fogleman/gg"
)

func DrawQuad(quad interfaces.Quad, original interfaces.Quad,name string,path []interfaces.Edge) {
	s,e := quad.GetNodesPos()

	nodes := original.Nodes()[s:e]

	ctx := gg.NewContext(1000, 1000)

	xStandardizer := 1000 / (quad.BrLon() - quad.TlLon())
	yStandardizer :=  1000 / (quad.BrLat() - quad.TlLat())

	for _, n := range nodes {

		for _, e := range n.OutEdges() {
			if e.EdgeType() == types.Road {
				primitives.DrawEdge(e, ctx, quad)
			}
		}
		ctx.SetRGB(0, 0, 0)
		ctx.Stroke()
		for _, e := range n.InEdges() {
			if e.EdgeType() == types.Road {
				primitives.DrawEdge(e, ctx, quad)
			}
		}
		ctx.SetRGB(0, 0, 0)
		ctx.Stroke()
	}

	if len(path) > 0 {
		sN := path[0].From()
		EN := path[len(path)-1].To()

		x := (sN.Longitude() - quad.TlLon()) * xStandardizer
		y := (sN.Latitude() - quad.TlLat()) * yStandardizer
		primitives.DrawNodeI(sN, ctx, x, y)
		primitives.DrawNodeI(sN, ctx, x, y)

		x = (EN.Longitude() - quad.TlLon()) * xStandardizer
		y = (EN.Latitude() - quad.TlLat()) * yStandardizer
		primitives.DrawNodeI(EN, ctx, x, y)
		primitives.DrawNodeI(EN, ctx, x, y)

		ctx.SetRGB(1, 1, 0)
		ctx.Fill()

		for _, e := range path {
			if e.EdgeType() == types.Road {
				primitives.DrawEdge(e, ctx, quad)
			}
		}
		ctx.SetRGB(1, 0, 0)
		ctx.SetLineWidth(10)
		ctx.Stroke()

		for _, e := range path {
			if e.EdgeType() == types.Bus {
				primitives.DrawEdge(e, ctx, quad)
			}
		}
		ctx.SetRGB(0, 1, 0)
		ctx.SetLineWidth(10)
		ctx.Stroke()

		for _, e := range path {
			if e.EdgeType() == types.Metro {
				primitives.DrawEdge(e, ctx, quad)
			}
		}
		ctx.SetRGB(0, 0, 1)
		ctx.SetLineWidth(10)
		ctx.Stroke()
	}

	ctx.SavePNG("images/"+name+".png")
}
