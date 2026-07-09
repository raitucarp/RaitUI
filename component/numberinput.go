package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func NumberInput() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignStretch)
	elem.GNode.SetWidth(140).SetMinWidth(120)
	elem.GNode.SetFlexShrink(0)

	field := core.NewElement(core.TypeInput)
	field.SetPlaceholder("0")
	field.GNode.SetFlexGrow(1)
	field.GNode.SetMinHeight(36)
	field.GNode.SetHeight(36)
	field.GNode.SetFlexShrink(0)

	stepper := core.NewElement(core.TypeVStack)
	stepper.FlexDirection(goda.FlexDirectionColumn)
	stepper.GNode.SetWidth(30).SetMinWidth(30)
	stepper.GNode.SetFlexShrink(0)

	up := core.NewElement(core.TypeBox)
	up.BackgroundColor(theme.Gray100)
	up.GNode.SetBorder(goda.EdgeBottom, 1)
	up.BorderColor(theme.Gray300)
	up.GNode.SetFlexGrow(1)
	up.GNode.SetFlexDirection(goda.FlexDirectionRow)
	up.GNode.SetJustifyContent(goda.JustifyCenter)
	up.GNode.SetAlignItems(goda.AlignCenter)
	up.Children(Text("\u25b2").FontSize(8).TextColor(theme.Gray600))

	down := core.NewElement(core.TypeBox)
	down.BackgroundColor(theme.Gray100)
	down.GNode.SetFlexGrow(1)
	down.GNode.SetFlexDirection(goda.FlexDirectionRow)
	down.GNode.SetJustifyContent(goda.JustifyCenter)
	down.GNode.SetAlignItems(goda.AlignCenter)
	down.Children(Text("\u25bc").FontSize(8).TextColor(theme.Gray600))

	stepper.Children(up, down)
	elem.Children(field, stepper)

	return elem
}
