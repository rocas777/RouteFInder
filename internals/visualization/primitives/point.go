package primitives

import (
	"edaa/internals/interfaces"
	"github.com/fogleman/gg"
)

func DrawNode(node interfaces.Node, ctx *gg.Context, x, y float64) {
	ctx.DrawCircle(x, y, 1)
}

func DrawNodeI(node interfaces.Node, ctx *gg.Context, x, y float64) {
	ctx.DrawCircle(x, y, 50)
}