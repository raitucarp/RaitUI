package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func Table() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetBorder(goda.EdgeAll, 1)
	elem.BorderColor(theme.Gray200)
	elem.BorderRadius(8)
	return elem
}

func TableHead() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.BackgroundColor(theme.Gray50)
	elem.GNode.SetBorder(goda.EdgeBottom, 1)
	elem.BorderColor(theme.Gray200)
	return elem
}

func TableBody() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	return elem
}

func TableRow() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetBorder(goda.EdgeBottom, 1)
	elem.BorderColor(theme.Gray100)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func TableCell(label string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetWidth(180).SetMinWidth(100)
	elem.Padding("12").PaddingX("16")
	elem.GNode.SetFlexShrink(0)

	t := Text(label)
	t.FontSize(14)
	t.TextColor(theme.Gray700)
	t.GNode.SetFlexShrink(0)
	elem.Children(t)

	return elem
}

func TableHeaderCell(label string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetWidth(180).SetMinWidth(100)
	elem.Padding("12").PaddingX("16")
	elem.GNode.SetFlexShrink(0)

	t := Text(label)
	t.FontSize(13)
	t.TextColor(theme.Gray600)
	t.GNode.SetFlexShrink(0)
	elem.Children(t)

	return elem
}
