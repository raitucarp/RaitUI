package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Kbd(content string) *core.Element {
	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)
	elem.TextColor(theme.Gray700)
	elem.BackgroundColor(theme.Gray50)
	elem.BorderColor(theme.Gray300)
	elem.BorderRadius(4)
	elem.FontSize(12)

	runes := len([]rune(content))
	w := float32(runes)*8 + 12
	h := float32(22)

	elem.GNode.SetWidth(w).SetMinWidth(w)
	elem.GNode.SetHeight(h).SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetBorder(goda.EdgeAll, 1)
	elem.BoxShadow(0, 1, 0, 0, theme.Gray300)

	return elem
}
