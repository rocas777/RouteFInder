package primitives

import (
	"edaa/internals/interfaces"
	"github.com/fogleman/gg"
)

func DrawEdge(edge interfaces.Edge, ctx *gg.Context, quad interfaces.Quad) {
	xStandardizer := (quad.BrLon() - quad.TlLon()) * 1000
	yStandardizer := (quad.BrLat() - quad.TlLat()) * 1000

	sx := edge.From().Longitude() * xStandardizer
	sy := edge.From().Latitude() * yStandardizer

	dx := edge.To().Longitude() * xStandardizer
	dy := edge.To().Latitude() * yStandardizer

	ctx.DrawLine(sx, sy, dx, dy)
}
