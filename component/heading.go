package component

import (
		"raitui/core"
	"raitui/theme"
)

func Heading(content string, level int) *core.Element {
	sizes := []float32{24, 20, 18, 16, 14, 12}
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}

	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)
	elem.TextColor(theme.Gray800)

	size := sizes[level-1]
	elem.FontSize(size)

	runes := len([]rune(content))
	w := float32(runes) * size * 0.65
	h := size * 1.5
	if w < 50 {
		w = 50
	}

	elem.GNode.SetWidth(w)
	elem.GNode.SetMinWidth(w)
	elem.GNode.SetHeight(h)
	elem.GNode.SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	return elem
}
