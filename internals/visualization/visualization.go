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

	//xStandardizer := 1000 / (quad.BrLon() - quad.TlLon())
	//yStandardizer :=  1000 / (quad.BrLat() - quad.TlLat())

	for _, n := range nodes {

		/*x := (n.Longitude() - quad.TlLon()) * xStandardizer
		y := (n.Latitude() - quad.TlLat()) * yStandardizer

		primitives.DrawNode(n, ctx, x, y)
		ctx.SetRGB(0, 0, 0)
		ctx.Fill()*/

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
	for _, e := range path {
		primitives.DrawEdge(e, ctx, quad)
	}
	ctx.SetRGB(1, 0, 0)
	ctx.SetLineWidth(10)
	ctx.Stroke()

	ctx.SavePNG("images/"+name+".png")
}
