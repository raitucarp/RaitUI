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
	ctx.SetMinWindowSize(500, 350)
	ctx.Run(root, "RaitUI - Box Example", 800, 600)
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
					raitui.Text("RaitUI Demo").TextColor(theme.White).FontSize(18),
				),

			raitui.HStack().Width("100%").Gap("12").Children(
				statCard("Revenue", "$12,430", "+12.5%", theme.Green500),
				statCard("Users", "1,284", "+8.2%", theme.Blue500),
				statCard("Orders", "562", "-3.1%", theme.Purple500),
			),

			raitui.Box().
				Width("100%").FlexGrow(1).
				Padding("20").Gap("12").
				BackgroundColor(theme.White).BorderRadius(10).
				BoxShadow(0, 2, 8, 0, rgba(0, 0, 0, 15)).
				Children(
					raitui.VStack().Gap("8").Children(
						raitui.Text("Recent Activity").TextColor(theme.Gray800).FontSize(16),
						row("Server deploy completed", theme.Green500),
						row("2 new signups", theme.Blue500),
						row("Database backup scheduled", theme.Yellow500),
						row("Disk usage at 85%", theme.Red500),
					),
				),
		)
}

func statCard(label, value, change string, accent color.NRGBA) *core.Element {
	chg := theme.Green600
	if change[0] == '-' {
		chg = theme.Red600
	}
	return raitui.Box().
		FlexGrow(1).Padding("16").
		BackgroundColor(theme.White).BorderRadius(10).
		BoxShadow(0, 2, 8, 0, rgba(0, 0, 0, 10)).
		Children(
			raitui.VStack().Gap("4").Children(
				raitui.Text(label).TextColor(theme.Gray500).FontSize(12),
				raitui.Text(value).TextColor(accent).FontSize(22),
				raitui.Text(change).TextColor(chg).FontSize(13),
			),
		)
}

func row(label string, accent color.NRGBA) *core.Element {
	return raitui.Box().
		Width("100%").PaddingX("12").PaddingY("10").
		BackgroundColor(fade(accent, 0.15)).BorderRadius(6).
		Children(
			raitui.Text(label).TextColor(theme.Gray700).FontSize(13),
		)
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
