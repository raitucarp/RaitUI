package main

import (
	"image/color"

	"raitui"
	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

func main() {
	root := buildUI()
	ctx := core.NewContext(theme.Gray50)
	ctx.SetMinWindowSize(500, 350)
	ctx.Run(root, "RaitUI - Layout Components", 800, 600)
}

func buildUI() *core.Element {
	return raitui.VStack().
		Width("100%").Height("100%").
		MinWidth("500").MinHeight("350").
		Padding("20").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			raitui.Box().
				Width("100%").PaddingX("20").PaddingY("16").
				BackgroundColor(theme.Blue500).BorderRadius(10).
				Children(
					raitui.HStack().
						Width("100%").
						JustifyContent(props.JustifySpaceBetween).
						AlignItems(props.AlignCenter).
						Children(
							raitui.Text("Layout Components").TextColor(theme.White).FontSize(18),
							raitui.HStack().Gap("8").Children(
								dot(theme.White, 0.3),
								dot(theme.White, 0.5),
								dot(theme.White, 0.3),
							),
						),
				),

			raitui.HStack().Gap("16").FlexWrap(props.WrapWrap).Children(
				raitui.Spacer(),
				buildCenterDemo(),
				buildFlexDemo(),
				buildSeparatorDemo(),
				raitui.Spacer(),
			),

			buildSeparatorLine(),

			raitui.VStack().Gap("12").Children(
				raitui.Text("Container Demo").TextColor(theme.Gray600).FontSize(12),

				raitui.Container().
					Width("100%").Padding("16").
					BackgroundColor(theme.White).BorderRadius(10).
					BoxShadow(0, 1, 4, 0, rgba(0, 0, 0, 10)).
					Children(
						raitui.VStack().Gap("6").Children(
							raitui.Text("Container (max-width aware)").TextColor(theme.Gray800).FontSize(15),
							raitui.Text("This box acts as a container. On narrow screens, content adapts.").TextColor(theme.Gray500).FontSize(12),
						),
					),
			),

			raitui.Spacer(),
		)
}

func buildCenterDemo() *core.Element {
	return raitui.Center().
		Width("200").Height("120").
		BackgroundColor(theme.White).BorderRadius(10).
		BoxShadow(0, 1, 4, 0, rgba(0, 0, 0, 10)).
		Children(
			raitui.Text("Center").TextColor(theme.Gray800).FontSize(14),
		)
}

func buildFlexDemo() *core.Element {
	return raitui.Flex(props.FlexDirectionColumn).
		Width("200").Height("120").
		Padding("16").Gap("8").
		BackgroundColor(theme.White).BorderRadius(10).
		BoxShadow(0, 1, 4, 0, rgba(0, 0, 0, 10)).
		Children(
			raitui.Text("Flex").TextColor(theme.Gray800).FontSize(14),
			raitui.HStack().Gap("6").Children(
				tag("col", theme.Blue500),
				tag("gap", theme.Green500),
			),
		)
}

func buildSeparatorDemo() *core.Element {
	return raitui.VStack().
		Width("200").Height("120").
		Padding("16").Gap("8").
		BackgroundColor(theme.White).BorderRadius(10).
		BoxShadow(0, 1, 4, 0, rgba(0, 0, 0, 10)).
		Children(
			raitui.Text("Separator").TextColor(theme.Gray800).FontSize(14),
			buildSeparatorLine(),
			raitui.Text("Below line").TextColor(theme.Gray500).FontSize(12),
		)
}

func buildSeparatorLine() *core.Element {
	return raitui.Separator().Width("100%")
}

func dot(c color.NRGBA, alpha float64) *core.Element {
	return raitui.Box().
		Width("28").Height("28").
		BackgroundColor(fade(c, alpha)).BorderRadius(14)
}

func tag(label string, c color.NRGBA) *core.Element {
	return raitui.Box().
		PaddingX("8").PaddingY("3").
		BackgroundColor(c).BorderRadius(4).
		Children(raitui.Text(label).TextColor(theme.White).FontSize(11))
}

func rgba(r, g, b, a uint8) color.NRGBA {
	return color.NRGBA{R: r, G: g, B: b, A: a}
}

func fade(c color.NRGBA, amt float64) color.NRGBA {
	return color.NRGBA{
		R: uint8(float64(c.R) + (255-float64(c.R))*amt),
		G: uint8(float64(c.G) + (255-float64(c.G))*amt),
		B: uint8(float64(c.B) + (255-float64(c.B))*amt),
		A: c.A,
	}
}
