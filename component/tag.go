package component

import (
	"raitui/core"
	"raitui/theme"
)

func Tag() *core.Element {
	return HStack().
		BackgroundColor(theme.Gray100).
		BorderRadius(4).
		PaddingY("4").
		PaddingX("10").
		Gap("6")
}

func TagLabel(label string) *core.Element {
	return Text(label).
		TextColor(theme.Gray700).
		FontSize(13)
}

func TagCloseButton(onClose func()) *core.Element {
	t := Text("\u00d7").FontSize(11).TextColor(theme.Gray500)

	btn := Center().
		Width("14").Height("14").
		MinWidth("14").MinHeight("14").
		BorderRadius(7).
		Children(t)

	if onClose != nil {
		btn.OnClick(onClose)
	}

	return btn
}
