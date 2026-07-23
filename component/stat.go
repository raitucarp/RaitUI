package component

import (
	"raitui/core"
	"raitui/theme"
)

func Stat() *core.Element {
	return VStack().
		Gap("4")
}

func StatLabel(text string) *core.Element {
	return Text(text).
		FontSize(13).
		TextColor(theme.Gray500)
}

func StatNumber(text string) *core.Element {
	return Text(text).
		FontSize(28).
		TextColor(theme.Gray800)
}

func StatHelpText(text string) *core.Element {
	return Text(text).
		FontSize(12).
		TextColor(theme.Gray400)
}
