package component

import (
	"raitui/core"
	"raitui/theme"
)

func Mark(content string) *core.Element {
	t := Text(content).FontSize(14).TextColor(theme.Gray800)

	return Box().
		BackgroundColor(theme.Yellow100).
		BorderRadius(2).
		PaddingX("3").
		Children(t)
}
