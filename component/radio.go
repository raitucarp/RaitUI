package component

import (
	"raitui/core"
	"raitui/theme"
)

func RadioGroup() *core.Element {
	return VStack().
		Gap("8")
}

func Radio(label string) *core.Element {
	elem := HStack().
		Gap("8")

	dot := Box().
		BorderColor(theme.Gray300).
		BorderWidth("2").
		Width("18").MinWidth("18").
		Height("18").MinHeight("18").
		BorderRadius(9)

	inner := Box().
		BackgroundColor(theme.Blue500).
		Width("8").MinWidth("8").
		Height("8").MinHeight("8").
		BorderRadius(4).
		Visible(false)

	dot.Children(inner)
	elem.Children(dot, Text(label).FontSize(14).TextColor(theme.Gray700))

	elem.OnClick(func() {
		inner.Visible(!inner.IsVisible())
	})

	return elem
}
