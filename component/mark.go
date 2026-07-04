package component

import (
		"raitui/core"
	"raitui/theme"
)

func Mark(content string) *core.Element {
	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)
	elem.TextColor(theme.Gray800)
	elem.BackgroundColor(theme.Yellow100)
	elem.BorderRadius(2)
	elem.FontSize(14)

	runes := len([]rune(content))
	w := float32(runes)*9 + 6
	h := float32(22)

	elem.GNode.SetWidth(w).SetMinWidth(w)
	elem.GNode.SetHeight(h).SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	return elem
}
