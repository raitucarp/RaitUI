package component

import (
	"raitui/core"
	"raitui/theme"
)

func Table() *core.Element {
	return VStack().
		BorderWidth("1").
		BorderColor(theme.Gray200).
		BorderRadius(8)
}

func TableHead() *core.Element {
	return VStack().
		BackgroundColor(theme.Gray50).
		BorderBottom("1").
		BorderColor(theme.Gray200)
}

func TableBody() *core.Element {
	return VStack()
}

func TableRow() *core.Element {
	return HStack().
		BorderBottom("1").
		BorderColor(theme.Gray100)
}

func TableCell(label string) *core.Element {
	t := Text(label).FontSize(14).TextColor(theme.Gray700)

	return Box().
		Width("180").MinWidth("100").
		Padding("12").PaddingX("16").
		Children(t)
}

func TableHeaderCell(label string) *core.Element {
	t := Text(label).FontSize(13).TextColor(theme.Gray600)

	return Box().
		Width("180").MinWidth("100").
		Padding("12").PaddingX("16").
		Children(t)
}
