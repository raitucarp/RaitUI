package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func ScrollAreaRoot() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetOverflow(goda.OverflowScroll)
	elem.BackgroundColor(theme.White)
	elem.BorderColor(theme.Gray200)
	elem.GNode.SetBorder(goda.EdgeAll, 1)
	elem.BorderRadius(8)
	elem.GNode.SetFlexShrink(1)
	elem.GNode.SetFlexGrow(1)
	return elem
}

func ScrollAreaViewport() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetOverflow(goda.OverflowScroll)
	elem.GNode.SetFlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetFlexGrow(1)
	elem.GNode.SetFlexShrink(1)
	return elem
}

func ScrollAreaContent() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	return elem
}

func ScrollAreaScrollbar() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetWidth(6).SetMinWidth(6)
	elem.BackgroundColor(theme.Gray100)
	elem.BorderRadius(3)
	return elem
}

func ScrollAreaThumb() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetWidth(6).SetMinWidth(6)
	elem.GNode.SetHeight(30).SetMinHeight(30)
	elem.BackgroundColor(theme.Gray300)
	elem.BorderRadius(3)
	return elem
}
