package component

import (
	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

func NumberInput() *core.Element {
	elem := HStack().
		Width("140").MinWidth("120").
		FlexShrink(0)

	field := core.NewElement(core.TypeInput)
	field.SetPlaceholder("0")
	field.FlexGrow(1)
	field.MinHeight("36")
	field.Height("36")
	field.FlexShrink(0)

	stepper := VStack().
		Width("30").MinWidth("30").
		FlexShrink(0)

	up := Center().
		BackgroundColor(theme.Gray100).
		BorderBottom("1").
		BorderColor(theme.Gray300).
		FlexGrow(1).
		Children(Text("\u25b2").FontSize(8).TextColor(theme.Gray600))

	down := Center().
		BackgroundColor(theme.Gray100).
		FlexGrow(1).
		Children(Text("\u25bc").FontSize(8).TextColor(theme.Gray600))

	stepper.Children(up, down)
	elem.Children(field, stepper)

	elem.AlignItems(props.AlignStretch)

	return elem
}
