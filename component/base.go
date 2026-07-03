package component

import (
	goda "goda"

	"raitui/core"
)

func Box() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func VStack() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func HStack() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func Text(content string) *core.Element {
	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)

	runes := len([]rune(content))
	w := float32(runes) * 9
	h := float32(22)

	if w < 22 {
		w = 22
	}

	elem.GNode.SetWidth(w)
	elem.GNode.SetMinWidth(w)
	elem.GNode.SetHeight(h)
	elem.GNode.SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	return elem
}

func Flex(dir goda.FlexDirection) *core.Element {
	elem := core.NewElement(core.TypeFlex)
	elem.FlexDirection(dir)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func Center() *core.Element {
	elem := core.NewElement(core.TypeCenter)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func Container() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func Separator() *core.Element {
	elem := core.NewElement(core.TypeSeparator)
	elem.GNode.SetWidth(100)
	elem.GNode.SetHeight(1)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func Spacer() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetFlexGrow(1)
	return elem
}
