package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Tag() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BackgroundColor(theme.Gray100)
	elem.BorderRadius(4)
	elem.GNode.SetFlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.GNode.SetFlexShrink(0)
	elem.PaddingX("10").PaddingY("4")
	elem.Gap("6")
	return elem
}

func TagLabel(label string) *core.Element {
	t := Text(label)
	t.TextColor(theme.Gray700)
	t.FontSize(13)
	t.GNode.SetFlexShrink(0)
	return t
}

func TagCloseButton(onClose func()) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetWidth(14).SetMinWidth(14)
	elem.GNode.SetHeight(14).SetMinHeight(14)
	elem.BorderRadius(7)
	elem.GNode.SetFlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.GNode.SetFlexShrink(0)

	t := Text("\u00d7")
	t.FontSize(11)
	t.TextColor(theme.Gray500)
	t.GNode.SetFlexShrink(0)
	elem.Children(t)

	if onClose != nil {
		elem.OnClick(onClose)
	}

	return elem
}
