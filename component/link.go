package component

import (
		"raitui/core"
	"raitui/theme"
)

func Link(content string) *core.Element {
	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)
	elem.TextColor(theme.Blue500)
	elem.FontSize(14)
	elem.SetUnderline(true)

	runes := len([]rune(content))
	w := float32(runes) * 9
	h := float32(22)

	elem.GNode.SetWidth(w).SetMinWidth(w)
	elem.GNode.SetHeight(h).SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	return elem
}
