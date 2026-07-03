package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func Button(label string) *core.Element {
	elem := core.NewElement(core.TypeButton)
	elem.BackgroundColor(theme.Blue500)
	elem.BorderRadius(6)
	elem.GNode.SetFlexShrink(0)

	runes := len([]rune(label))
	w := float32(runes)*10 + 28
	h := float32(36)
	elem.GNode.SetWidth(w).SetMinWidth(w)
	elem.GNode.SetHeight(h).SetMinHeight(h)
	elem.SetTextContent(label)
	elem.TextColor(theme.White).FontSize(14)

	elem.OnHover(func(entered bool) {
		if entered {
			elem.BackgroundColor(theme.Blue600)
		} else {
			elem.BackgroundColor(theme.Blue500)
		}
	})

	return elem
}

func OutlineButton(label string) *core.Element {
	elem := core.NewElement(core.TypeButton)
	elem.BackgroundColor(theme.Transparent)
	elem.BorderColor(theme.Blue500)
	elem.BorderRadius(6)
	elem.GNode.SetBorder(goda.EdgeAll, 1.5)
	elem.GNode.SetFlexShrink(0)

	runes := len([]rune(label))
	w := float32(runes)*10 + 28
	h := float32(36)
	elem.GNode.SetWidth(w).SetMinWidth(w)
	elem.GNode.SetHeight(h).SetMinHeight(h)
	elem.SetTextContent(label)
	elem.TextColor(theme.Blue500).FontSize(14)

	elem.OnHover(func(entered bool) {
		if entered {
			elem.BackgroundColor(theme.Blue50)
		} else {
			elem.BackgroundColor(theme.Transparent)
		}
	})

	return elem
}

func Checkbox(label string, checked bool) *core.Element {
	elem := core.NewElement(core.TypeCheckbox)
	elem.SetTextContent(label)
	elem.SetChecked(checked)

	runes := len([]rune(label))
	w := float32(runes)*9 + 40
	h := float32(28)

	elem.GNode.SetWidth(w)
	elem.GNode.SetMinWidth(w)
	elem.GNode.SetHeight(h)
	elem.GNode.SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	elem.OnClick(func() {
		elem.SetChecked(!elem.IsChecked())
	})

	return elem
}

func Switch(label string, checked bool) *core.Element {
	elem := core.NewElement(core.TypeSwitch)
	elem.SetTextContent(label)
	elem.SetChecked(checked)

	runes := len([]rune(label))
	w := float32(runes)*9 + 60
	h := float32(28)

	elem.GNode.SetWidth(w)
	elem.GNode.SetMinWidth(w)
	elem.GNode.SetHeight(h)
	elem.GNode.SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	elem.OnClick(func() {
		elem.SetChecked(!elem.IsChecked())
	})

	return elem
}

func Input(placeholder string) *core.Element {
	elem := core.NewElement(core.TypeInput)
	elem.SetPlaceholder(placeholder)
	elem.GNode.SetWidth(200)
	elem.GNode.SetMinHeight(36)
	elem.GNode.SetHeight(36)
	elem.GNode.SetFlexShrink(0)

	return elem
}

func TextArea(placeholder string) *core.Element {
	elem := core.NewElement(core.TypeInput)
	elem.SetPlaceholder(placeholder)
	elem.GNode.SetWidth(300)
	elem.GNode.SetMinHeight(80)
	elem.GNode.SetHeight(80)
	elem.GNode.SetFlexShrink(0)

	return elem
}

func colorWithAlpha(c theme.Color, a uint8) theme.Color {
	c.A = a
	return c
}
