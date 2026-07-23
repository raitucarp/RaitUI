package component

import (
	"raitui/core"
	"raitui/theme"
)

func Link(content string) *core.Element {
	return Text(content).
		TextColor(theme.Blue500).
		FontSize(14).
		SetUnderline(true)
}
