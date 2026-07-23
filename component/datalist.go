package component

import (
	"raitui/core"
	"raitui/theme"
)

func DataList() *core.Element {
	return VStack().
		Gap("8")
}

func DataListItem() *core.Element {
	return HStack().
		Gap("16").
		PaddingY("8").
		BorderBottom("1").
		BorderColor(theme.Gray100)
}

func DataListLabel(text string) *core.Element {
	return Text(text).
		FontSize(14).
		TextColor(theme.Gray500).
		Width("120").MinWidth("120")
}

func DataListValue(text string) *core.Element {
	return Text(text).
		FontSize(14).
		TextColor(theme.Gray800)
}
