package main

import (
	"image/color"

	"raitui"
	"raitui/core"
	"raitui/theme"
)

func main() {
	root := buildUI()
	ctx := core.NewContext(theme.Gray50)
	ctx.SetMinWindowSize(400, 400)
	ctx.Run(root, "RaitUI - Text Components", 700, 580)
}

func buildUI() *core.Element {
	return raitui.VStack().
		Width("100%").Height("100%").MinWidth("400").MinHeight("400").
		Padding("24").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			raitui.Text("Typography Components").TextColor(theme.Gray800).FontSize(18),

			raitui.Separator().Width("100%"),

			raitui.VStack().Gap("10").Children(
				label("Heading 1-4"),
				raitui.Heading("Heading Level 1", 1),
				raitui.Heading("Heading Level 2", 2),
				raitui.Heading("Heading Level 3", 3),
				raitui.Heading("Heading Level 4", 4),
			),

			raitui.VStack().Gap("10").Children(
				label("Inline Text Components"),
				raitui.HStack().Gap("8").Children(
					raitui.Code("import raitui"),
					raitui.Kbd("Ctrl+C"),
					raitui.Mark("highlighted text"),
					link("github.com"),
				),
			),

			raitui.VStack().Gap("10").Children(
				label("Blockquote"),
				raitui.Blockquote(
					"This is a blockquote component.",
					"It has a blue left border accent.",
				),
			),

			raitui.VStack().Gap("10").Children(
				label("Unordered List"),
				raitui.List([]string{"First item", "Second item", "Third item"}, false),
			),

			raitui.VStack().Gap("10").Children(
				label("Ordered List"),
				raitui.List([]string{"Step one", "Step two", "Step three"}, true),
			),
		)
}

func label(text string) *core.Element {
	return raitui.Text(text).TextColor(theme.Gray400).FontSize(11)
}

func link(text string) *core.Element {
	l := raitui.Link(text)
	l.OnClick(func() {})
	return l
}

var _ = color.NRGBA{}
