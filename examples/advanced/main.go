package main

import (
	"image/color"

	"raitui"
	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

func main() {
	raitui.SetState("clicks", 0)
	root := buildUI()
	ctx := core.NewContext(theme.Gray50)
	ctx.SetMinWindowSize(600, 400)
	ctx.Run(root, "RaitUI - Advanced", 900, 640)
}

func buildUI() *core.Element {
	getClicks := raitui.StateInt("clicks", 0)
	setClicks := raitui.Setter[int]("clicks")
	return raitui.VStack().
		Width("100%").Height("100%").
		MinWidth("600").MinHeight("400").
		Padding("20").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			header(),
			body(getClicks, setClicks),
		)
}

func header() *core.Element {
	return raitui.Box().
		Width("860").Padding("16").
		BackgroundColor(theme.White).BorderRadius(10).
		BoxShadow(0, 2, 8, 0, rgba(0, 0, 0, 10)).
		Children(
			raitui.HStack().
				Width("100%").Gap("10").
				JustifyContent(props.JustifySpaceBetween).
				AlignItems(props.AlignCenter).
				Children(
					raitui.Text("RaitUI — Fixed Size Layout").TextColor(theme.Gray800).FontSize(16),
					dots(),
				),
		)
}

func dots() *core.Element {
	return raitui.HStack().Gap("10").Children(
		dot(theme.Blue500), dot(theme.Green500), dot(theme.Purple500),
	)
}

func dot(c color.NRGBA) *core.Element {
	return raitui.Box().Width("30").Height("30").BackgroundColor(c).BorderRadius(15)
}

func body(getClicks func() int, setClicks func(int)) *core.Element {
	return raitui.HStack().Gap("16").Children(
		sidebar(),
		mainContent(getClicks, setClicks),
	)
}

func sidebar() *core.Element {
	return raitui.Box().
		Width("180").Padding("8").
		BackgroundColor(theme.White).BorderRadius(10).
		BoxShadow(0, 1, 4, 0, rgba(0, 0, 0, 10)).
		Children(
			raitui.VStack().Gap("4").Children(
				item("Dashboard", true),
				item("Analytics", false),
				item("Users", false),
				item("Settings", false),
			),
		)
}

func item(label string, active bool) *core.Element {
	bg := theme.Transparent
	txt := theme.Gray600
	if active {
		bg = theme.Blue500
		txt = theme.White
	}
	return raitui.Box().
		Width("100%").PaddingX("14").PaddingY("9").
		BackgroundColor(bg).BorderRadius(8).
		Children(
			raitui.Text(label).TextColor(txt).FontSize(13),
		)
}

func mainContent(getClicks func() int, setClicks func(int)) *core.Element {
	return raitui.VStack().Width("560").Gap("16").Children(
		clickCard(getClicks, setClicks),
		statsRow(),
		tagSection(),
	)
}

func clickCard(getClicks func() int, setClicks func(int)) *core.Element {
	card := raitui.Box().
		Width("100%").Padding("20").
		BackgroundColor(theme.White).BorderRadius(10).
		BoxShadow(0, 2, 12, 0, rgba(0, 0, 0, 15))

	card.OnClick(func() { setClicks(getClicks() + 1) })
	card.OnHover(func(entered bool) {
		if entered {
			card.BackgroundColor(theme.Blue50)
		} else {
			card.BackgroundColor(theme.White)
		}
	})
	card.Children(
		raitui.VStack().Gap("8").Children(
			raitui.Text("Interactive Card").TextColor(theme.Gray800).FontSize(16),
			raitui.Text("Click card to test OnClick. Hover for OnHover.").TextColor(theme.Gray500).FontSize(12),
		),
	)
	return card
}

func statsRow() *core.Element {
	return raitui.HStack().Width("100%").Gap("16").Children(
		statCard("Revenue", "$45,200", "+12.5%", theme.Green600),
		statCard("Users", "2,840", "+8.2%", theme.Blue600),
		statCard("Orders", "1,204", "-3.1%", theme.Red600),
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
				raitui.Text(value).TextColor(theme.Gray800).FontSize(22),
				raitui.Text(change).TextColor(chg).FontSize(13),
			),
		)
}

func tagSection() *core.Element {
	return raitui.Box().
		Width("100%").Padding("16").
		BackgroundColor(theme.White).BorderRadius(10).
		BoxShadow(0, 1, 4, 0, rgba(0, 0, 0, 10)).
		Children(
			raitui.VStack().Gap("10").Children(
				raitui.Text("Features").TextColor(theme.Gray800).FontSize(16),
				raitui.HStack().Gap("8").FlexWrap(props.WrapWrap).Children(
					tag("ID/Class", theme.Blue500),
					tag("Shadow", theme.Purple500),
					tag("Border", theme.Green500),
					tag("OnClick", theme.Orange500),
					tag("OnHover", theme.Pink500),
					tag("State", theme.Cyan500),
					tag("Radius", theme.Teal500),
					tag("Flex", theme.Yellow600),
				),
			),
		)
}

func tag(label string, c color.NRGBA) *core.Element {
	return raitui.Box().
		PaddingX("10").PaddingY("5").
		BackgroundColor(c).BorderRadius(6).
		Children(
			raitui.Text(label).TextColor(theme.White).FontSize(12),
		)
}

func rgba(r, g, b, a uint8) color.NRGBA {
	return color.NRGBA{R: r, G: g, B: b, A: a}
}
