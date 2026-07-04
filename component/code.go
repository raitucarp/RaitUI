package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Code(content string) *core.Element {
	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)
	elem.FontSize(13)
	elem.TextColor(theme.Gray700)
	elem.BackgroundColor(theme.Gray100)
	elem.BorderColor(theme.Gray200)
	elem.BorderRadius(4)

	runes := len([]rune(content))
	w := float32(runes)*9 + 16
	h := float32(24)

	elem.GNode.SetWidth(w).SetMinWidth(w)
	elem.GNode.SetHeight(h).SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetBorder(goda.EdgeAll, 1)

	return elem
}
