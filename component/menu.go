package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Menu() *core.Element {
	elem := core.NewElement(core.TypeMenu)
	elem.BackgroundColor(theme.White)
	elem.BorderRadius(8)
	elem.BoxShadow(0, 4, 12, 0, colorWithAlpha(theme.Black, 15))
	elem.GNode.SetBorder(goda.EdgeAll, 1)
	elem.BorderColor(theme.Gray200)
	elem.GNode.SetFlexDirection(goda.FlexDirectionColumn)
	elem.Padding("4")
	elem.Gap("2")
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetMinWidth(160)
	return elem
}

func MenuItem(label string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.PaddingX("12").PaddingY("8")
	elem.BorderRadius(6)
	elem.GNode.SetFlexShrink(0)

	t := Text(label)
	t.TextColor(theme.Gray700)
	t.FontSize(14)
	elem.Children(t)

	elem.OnHover(func(entered bool) {
		if entered {
			elem.BackgroundColor(theme.Gray100)
		} else {
			elem.BackgroundColor(theme.Transparent)
		}
	})

	return elem
}
