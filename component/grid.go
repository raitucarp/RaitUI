package component

import (
	goda "goda"
	"raitui/core"
)

func Grid() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetFlexWrap(goda.WrapWrap)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	return elem
}

func GridItem() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetFlexGrow(1)
	elem.GNode.SetFlexBasis(200)
	return elem
}

func SimpleGrid() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetFlexWrap(goda.WrapWrap)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	elem.Gap("16")
	return elem
}

func Icon(name string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetWidth(24).SetMinWidth(24)
	elem.GNode.SetHeight(24).SetMinHeight(24)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetFlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.SetTextContent(name)
	return elem
}
