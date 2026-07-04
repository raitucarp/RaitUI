package main

import (
	"image/color"

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
	return VStack().
		Width("100%").Height("100%").MinWidth("400").MinHeight("400").
		Padding("24").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			Text("Typography Components").TextColor(theme.Gray800).FontSize(18),

			Separator().Width("100%"),

			VStack().Gap("10").Children(
				label("Heading 1-4"),
				Heading("Heading Level 1", 1),
				Heading("Heading Level 2", 2),
				Heading("Heading Level 3", 3),
				Heading("Heading Level 4", 4),
			),

			VStack().Gap("10").Children(
				label("Inline Text Components"),
				HStack().Gap("8").Children(
					Code("import raitui"),
					Kbd("Ctrl+C"),
					Mark("highlighted text"),
					link("github.com"),
				),
			),

			VStack().Gap("10").Children(
				label("Blockquote"),
				Blockquote(
					"This is a blockquote component.",
					"It has a blue left border accent.",
				),
			),

			VStack().Gap("10").Children(
				label("Unordered List"),
				List([]string{"First item", "Second item", "Third item"}, false),
			),

			VStack().Gap("10").Children(
				label("Ordered List"),
				List([]string{"Step one", "Step two", "Step three"}, true),
			),
		)
}

func label(text string) *core.Element {
	return Text(text).TextColor(theme.Gray400).FontSize(11)
}

func link(text string) *core.Element {
	l := Link(text)
	l.OnClick(func() {})
	return l
}

var _ = color.NRGBA{}
