package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func Tabs() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func TabList() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetFlexShrink(0)
	elem.Gap("2")
	elem.GNode.SetBorder(goda.EdgeBottom, 2)
	elem.BorderColor(theme.Gray200)
	return elem
}

func Tab(label string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetFlexShrink(0)
	elem.PaddingX("16").PaddingY("10")
	elem.BorderRadius(6)

	t := Text(label)
	t.FontSize(14)
	t.TextColor(theme.Gray500)
	t.GNode.SetFlexShrink(0)
	elem.Children(t)

	elem.OnHover(func(entered bool) {
		if entered {
			elem.BackgroundColor(theme.Gray100)
		} else {
			elem.BackgroundColor(theme.Transparent)
		}
	})

	return elem
}

func TabPanels() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetFlexShrink(0)
	elem.Padding("16")
	return elem
}

func TabPanel() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.Gap("8")
	elem.Visible(false)
	return elem
}
