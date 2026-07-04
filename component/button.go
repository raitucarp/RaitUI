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
