package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func RadioGroup() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.Gap("8")
	return elem
}

func Radio(label string) *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.Gap("8")
	elem.GNode.SetFlexShrink(0)

	dot := core.NewElement(core.TypeBox)
	dot.BorderColor(theme.Gray300)
	dot.GNode.SetBorder(goda.EdgeAll, 2)
	dot.GNode.SetWidth(18).SetMinWidth(18)
	dot.GNode.SetHeight(18).SetMinHeight(18)
	dot.BorderRadius(9)
	dot.GNode.SetFlexShrink(0)

	inner := core.NewElement(core.TypeBox)
	inner.BackgroundColor(theme.Blue500)
	inner.GNode.SetWidth(8).SetMinWidth(8)
	inner.GNode.SetHeight(8).SetMinHeight(8)
	inner.BorderRadius(4)
	inner.Visible(false)

	dot.Children(inner)
	elem.Children(dot, Text(label).FontSize(14).TextColor(theme.Gray700))

	elem.OnClick(func() {
		inner.Visible(!inner.IsVisible())
	})

	return elem
}
