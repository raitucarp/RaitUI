package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Badge(label string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetFlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)

	runes := len([]rune(label))
	w := float32(runes)*8 + 14
	h := float32(22)

	elem.GNode.SetWidth(w).SetMinWidth(w)
	elem.GNode.SetHeight(h).SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	elem.BackgroundColor(theme.Gray100)
	elem.BorderRadius(4)

	t := Text(label)
	t.TextColor(theme.Gray700)
	t.FontSize(11)
	elem.Children(t)

	return elem
}
