package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func Accordion() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.Gap("1")
	return elem
}

func AccordionItem() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetBorder(goda.EdgeAll, 1)
	elem.BorderColor(theme.Gray200)
	elem.BorderRadius(8)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func AccordionHeader(label string) *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.GNode.SetFlexShrink(0)
	elem.Padding("12").PaddingX("16")

	t := Text(label)
	t.FontSize(15)
	t.TextColor(theme.Gray700)
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

func AccordionPanel() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.Padding("12").PaddingX("16")
	elem.Gap("8")
	elem.Visible(false)
	return elem
}
