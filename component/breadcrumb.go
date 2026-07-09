package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func Breadcrumb() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.Gap("4")
	elem.GNode.SetFlexShrink(0)
	return elem
}

func BreadcrumbItem(label string) *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.Gap("4")
	elem.GNode.SetFlexShrink(0)

	t := Text(label)
	t.FontSize(14)
	t.TextColor(theme.Gray500)
	t.GNode.SetFlexShrink(0)
	elem.Children(t)

	return elem
}

func BreadcrumbLink(label string) *core.Element {
	t := Text(label)
	t.FontSize(14)
	t.TextColor(theme.Blue500)
	t.GNode.SetFlexShrink(0)
	t.SetUnderline(true)
	return t
}

func BreadcrumbSeparator() *core.Element {
	t := Text("/")
	t.FontSize(14)
	t.TextColor(theme.Gray400)
	t.GNode.SetFlexShrink(0)
	return t
}
