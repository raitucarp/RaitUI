package component

import (
	"raitui/core"
	"raitui/theme"
)

func Code(content string) *core.Element {
	t := Text(content).FontSize(13).TextColor(theme.Gray700)

	return Box().
		BackgroundColor(theme.Gray100).
		BorderColor(theme.Gray200).
		BorderRadius(4).
		BorderWidth("1").
		PaddingY("2").
		PaddingX("8").
		Children(t)
}
