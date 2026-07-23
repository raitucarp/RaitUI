package component

import (
	"strconv"

	"raitui/core"
)

func Checkbox(label string, checked bool) *core.Element {
	elem := core.NewElement(core.TypeCheckbox)
	elem.SetTextContent(label)
	elem.SetChecked(checked)

	runes := len([]rune(label))
	w := float32(runes)*9 + 40
	wStr := strconv.FormatFloat(float64(w), 'f', 0, 32)

	elem.Width(wStr).MinWidth(wStr)
	elem.Height("28").MinHeight("28")
	elem.FlexShrink(0)

	elem.OnClick(func() {
		elem.SetChecked(!elem.IsChecked())
	})

	return elem
}
