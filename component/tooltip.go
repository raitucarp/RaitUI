package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Tooltip(label string) *core.Element {
	elem := core.NewElement(core.TypeTooltip)
	elem.BackgroundColor(theme.Gray700)
	elem.BorderRadius(6)
	elem.GNode.SetFlexShrink(0)

	runes := len([]rune(label))
	w := float32(runes)*9 + 16
	h := float32(28)

	elem.GNode.SetWidth(w).SetMinWidth(w)
	elem.GNode.SetHeight(h).SetMinHeight(h)

	elem.GNode.SetFlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)

	t := Text(label)
	t.TextColor(theme.White)
	t.FontSize(12)
	t.TextAlign(core.AlignCenter)
	elem.Children(t)

	return elem
}
