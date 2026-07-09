package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func DataList() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.Gap("8")
	return elem
}

func DataListItem() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.Gap("16")
	elem.PaddingY("8")
	elem.GNode.SetBorder(goda.EdgeBottom, 1)
	elem.BorderColor(theme.Gray100)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func DataListLabel(text string) *core.Element {
	t := Text(text)
	t.FontSize(14)
	t.TextColor(theme.Gray500)
	t.GNode.SetWidth(120).SetMinWidth(120)
	t.GNode.SetFlexShrink(0)
	return t
}

func DataListValue(text string) *core.Element {
	t := Text(text)
	t.FontSize(14)
	t.TextColor(theme.Gray800)
	t.GNode.SetFlexShrink(0)
	return t
}
