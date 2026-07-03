package main

import (
	"image/color"

	"raitui"
	"raitui/core"
	"raitui/theme"
)

func main() {
	root := raitui.VStack().
		Width("400").Height("300").
		Padding("20").Gap("12").
		BackgroundColor(theme.Gray50).
		Children(
			raitui.Box().
				Width("200").Height("80").
				BackgroundColor(theme.Red500).BorderRadius(8).
				Children(
					raitui.Text("Hello RaitUI!").
						TextColor(theme.White).FontSize(14),
				),
			raitui.Text("Second line").TextColor(theme.Gray800),
		)

	bg := color.NRGBA{R: 255, G: 0, B: 255, A: 255}
	ctx := core.NewContext(bg)
	ctx.Debug = true
	ctx.Run(root, "Minimal — MAGENTA BG", 500, 400)
}
