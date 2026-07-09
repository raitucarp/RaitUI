package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func Stat() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.Gap("4")
	return elem
}

func StatLabel(text string) *core.Element {
	t := Text(text)
	t.FontSize(13)
	t.TextColor(theme.Gray500)
	t.GNode.SetFlexShrink(0)
	return t
}

func StatNumber(text string) *core.Element {
	t := Text(text)
	t.FontSize(28)
	t.TextColor(theme.Gray800)
	t.GNode.SetFlexShrink(0)
	return t
}

func StatHelpText(text string) *core.Element {
	t := Text(text)
	t.FontSize(12)
	t.TextColor(theme.Gray400)
	t.GNode.SetFlexShrink(0)
	return t
}
