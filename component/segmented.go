package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func SegmentedControl() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.BackgroundColor(theme.Gray100)
	elem.BorderRadius(8)
	elem.Padding("2")
	elem.Gap("1")
	elem.GNode.SetFlexShrink(0)
	return elem
}

func Segment(label string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.PaddingX("16").PaddingY("8")
	elem.BorderRadius(6)
	elem.GNode.SetFlexShrink(0)

	t := Text(label)
	t.FontSize(14)
	t.TextColor(theme.Gray600)
	t.GNode.SetFlexShrink(0)
	elem.Children(t)

	elem.OnHover(func(entered bool) {
		if entered {
			elem.BackgroundColor(theme.Gray200)
		} else {
			elem.BackgroundColor(theme.Transparent)
		}
	})

	return elem
}
