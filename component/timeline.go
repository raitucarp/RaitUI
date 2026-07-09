package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func Timeline() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.Gap("0")
	return elem
}

func TimelineItem() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.Gap("12")
	elem.GNode.SetFlexShrink(0)
	return elem
}

func TimelineDot() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BackgroundColor(theme.Blue500)
	elem.BorderRadius(9999)
	elem.GNode.SetWidth(16).SetMinWidth(16)
	elem.GNode.SetHeight(16).SetMinHeight(16)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetAlignSelf(goda.AlignCenter)
	return elem
}

func TimelineContent() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.PaddingY("4").PaddingBottom("24")
	elem.Gap("4")
	elem.FlexGrow(1)
	elem.GNode.SetBorder(goda.EdgeLeft, 2)
	elem.BorderColor(theme.Gray200)
	elem.PaddingLeft("16")
	return elem
}
