package component

import (
		"raitui/core"
)

func Checkbox(label string, checked bool) *core.Element {
	elem := core.NewElement(core.TypeCheckbox)
	elem.SetTextContent(label)
	elem.SetChecked(checked)

	runes := len([]rune(label))
	w := float32(runes)*9 + 40
	h := float32(28)

	elem.GNode.SetWidth(w).SetMinWidth(w)
	elem.GNode.SetHeight(h).SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	elem.OnClick(func() {
		elem.SetChecked(!elem.IsChecked())
	})

	return elem
}
