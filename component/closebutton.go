package component

import (
	"raitui/core"
	"raitui/theme"
)

func CloseButton() *core.Element {
	t := Text("\u00d7").FontSize(18).TextColor(theme.Gray500)

	btn := Center().
		Width("32").Height("32").
		MinWidth("32").MinHeight("32").
		BackgroundColor(theme.Transparent).
		BorderRadius(6).
		Children(t)

	btn.OnHover(func(entered bool) {
		if entered {
			btn.BackgroundColor(theme.Gray100)
		} else {
			btn.BackgroundColor(theme.Transparent)
		}
	})

	return btn
}
