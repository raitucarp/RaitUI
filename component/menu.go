package component

import (
	"raitui/core"
	"raitui/theme"
)

func Menu() *core.Element {
	elem := core.NewElement(core.TypeMenu)
	elem.BackgroundColor(theme.White)
	elem.BorderRadius(8)
	elem.BoxShadow(0, 4, 12, 0, colorWithAlpha(theme.Black, 15))
	elem.BorderWidth("1")
	elem.BorderColor(theme.Gray200)
	elem.Padding("4")
	elem.Gap("2")
	elem.FlexShrink(0)
	elem.MinWidth("160")
	return elem
}

func MenuItem(label string) *core.Element {
	t := Text(label).TextColor(theme.Gray700).FontSize(14)

	item := Box().
		PaddingX("12").PaddingY("8").
		BorderRadius(6).
		Children(t)

	item.OnHover(func(entered bool) {
		if entered {
			item.BackgroundColor(theme.Gray100)
		} else {
			item.BackgroundColor(theme.Transparent)
		}
	})

	return item
}
