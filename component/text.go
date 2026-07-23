package component

import (
	"strconv"

	"raitui/core"
)

func Text(content string) *core.Element {
	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)

	runes := len([]rune(content))
	w := float32(runes) * 9
	h := float32(22)

	if w < 22 {
		w = 22
	}

	wStr := strconv.FormatFloat(float64(w), 'f', 0, 32)
	hStr := strconv.FormatFloat(float64(h), 'f', 0, 32)

	elem.Width(wStr)
	elem.MinWidth(wStr)
	elem.Height(hStr)
	elem.MinHeight(hStr)
	elem.FlexShrink(0)

	return elem
}
