package component

import (
	"raitui/core"
	"raitui/theme"
)

func Kbd(content string) *core.Element {
	t := Text(content).FontSize(12).TextColor(theme.Gray700)

	return Box().
		BackgroundColor(theme.Gray50).
		BorderColor(theme.Gray300).
		BorderRadius(4).
		BorderWidth("1").
		BoxShadow(0, 1, 0, 0, theme.Gray300).
		PaddingY("2").
		PaddingX("6").
		Children(t)
}
