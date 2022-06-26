package primitives

import (
	"edaa/internals/interfaces"
	"github.com/fogleman/gg"
)

func DrawEdge(edge interfaces.Edge, ctx *gg.Context, quad interfaces.Quad) {
	xStandardizer := 1000 / (quad.BrLon() - quad.TlLon())
	yStandardizer :=  1000 / (quad.BrLat() - quad.TlLat())

	sx := (edge.From().Longitude() - quad.TlLon()) * xStandardizer
	sy := (edge.From().Latitude() - quad.TlLat()) * yStandardizer

	dx := (edge.To().Longitude() - quad.TlLon()) * xStandardizer
	dy := (edge.To().Latitude() - quad.TlLat()) * yStandardizer

	ctx.DrawLine(sx, sy, dx, dy)
}
