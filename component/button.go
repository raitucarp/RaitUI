package component

import (
	"raitui/core"
	"raitui/theme"
)

func Button(label string) *core.Element {
	textEl := Text(label).TextColor(theme.White).FontSize(14)

	btn := Center().
		BackgroundColor(theme.Blue500).
		BorderRadius(6).
		PaddingY("8").
		PaddingX("16").
		Children(textEl)

	btn.OnHover(func(entered bool) {
		if entered {
			btn.BackgroundColor(theme.Blue600)
		} else {
			btn.BackgroundColor(theme.Blue500)
		}
	})

	return btn
}

func OutlineButton(label string) *core.Element {
	textEl := Text(label).TextColor(theme.Blue500).FontSize(14)

	btn := Center().
		BackgroundColor(theme.Transparent).
		BorderColor(theme.Blue500).
		BorderRadius(6).
		BorderWidth("1.5").
		PaddingY("8").
		PaddingX("16").
		Children(textEl)

	btn.OnHover(func(entered bool) {
		if entered {
			btn.BackgroundColor(theme.Blue50)
		} else {
			btn.BackgroundColor(theme.Transparent)
		}
	})

	return btn
}
