package component

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2/vector"

	"raitui/core"
)

func Spinner() *core.Element {
	return Canvas(func(ctx core.CanvasContext) {
		cx := ctx.W / 2
		cy := ctx.H / 2
		r := ctx.W / 2
		if ctx.H/2 < r {
			r = ctx.H / 2
		}
		if r < 4 {
			r = 4
		}

		clr := core.ColorAccent
		segments := 12
		for i := 0; i < segments; i++ {
			angle := float64(ctx.Frame)*0.15 - float64(i)*2*math.Pi/float64(segments)
			a := uint8(255 * (1 - float64(i)/float64(segments)))
			sc := color.NRGBA{R: clr.R, G: clr.G, B: clr.B, A: a}
			sx := cx + float32(math.Cos(angle))*r*0.6
			sy := cy + float32(math.Sin(angle))*r*0.6
			ex := cx + float32(math.Cos(angle))*r*0.9
			ey := cy + float32(math.Sin(angle))*r*0.9
			vector.StrokeLine(ctx.Screen, sx, sy, ex, ey, 2, sc, true)
		}
	}).Width("24").MinWidth("24").Height("24").MinHeight("24")
}
