package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func CloseButton() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BackgroundColor(theme.Transparent)
	elem.BorderRadius(6)
	elem.GNode.SetWidth(32).SetMinWidth(32)
	elem.GNode.SetHeight(32).SetMinHeight(32)
	elem.GNode.SetFlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.GNode.SetFlexShrink(0)

	t := Text("\u00d7")
	t.FontSize(18)
	t.TextColor(theme.Gray500)
	t.GNode.SetFlexShrink(0)
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
