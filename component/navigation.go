package component

import (
	"fmt"

	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

func Drawer() *core.Element {
	elem := Box().
		BackgroundColor(theme.White).
		BoxShadow(0, 0, 20, 0, colorWithAlpha(theme.Black, 20)).
		FlexDirection(props.FlexDirectionColumn).
		Position(props.PositionAbsolute).
		Height("100%").
		MinWidth("300").
		MaxWidth("400").
		Padding("24").Gap("16").
		Visible(false)
	elem.PositionTop("0")
	elem.PositionRight("0")
	return elem
}

func Collapsible() *core.Element {
	return VStack().
		BorderWidth("1").
		BorderColor(theme.Gray200).
		BorderRadius(8)
}

func CollapsibleTrigger(label string) *core.Element {
	t := Text(label).FontSize(15).TextColor(theme.Gray700)

	trigger := HStack().
		Padding("12").PaddingX("16").
		Children(t)

	trigger.OnHover(func(entered bool) {
		if entered {
			trigger.BackgroundColor(theme.Gray100)
		} else {
			trigger.BackgroundColor(theme.Transparent)
		}
	})

	return trigger
}

func CollapsibleContent() *core.Element {
	return VStack().
		Padding("12").PaddingX("16").
		Gap("8").
		Visible(false)
}

func Pagination() *core.Element {
	return HStack().
		Gap("4")
}

func PageButton(label string) *core.Element {
	t := Text(label).FontSize(14).TextColor(theme.Gray600)

	btn := Center().
		Width("36").MinWidth("36").
		Height("36").MinHeight("36").
		BorderRadius(6).
		Children(t)

	btn.OnHover(func(entered bool) {
		if entered {
			btn.BackgroundColor(theme.Gray100)
		} else {
			btn.BackgroundColor(theme.Transparent)
		}
	})

	return btn
}

func Steps() *core.Element {
	return HStack().
		Gap("0")
}

func Step(active bool) *core.Element {
	return HStack().
		Gap("8").
		FlexShrink(0)
}

func StepIndicator(number int, active bool) *core.Element {
	t := Text(fmt.Sprintf("%d", number)).FontSize(12)

	indicator := Center().
		Width("28").MinWidth("28").
		Height("28").MinHeight("28").
		BorderRadius(14).
		Children(t)

	if active {
		indicator.BackgroundColor(theme.Blue500)
		t.TextColor(theme.White)
	} else {
		indicator.BackgroundColor(theme.Gray200)
		t.TextColor(theme.Gray600)
	}

	return indicator
}

func StepSeparator() *core.Element {
	return Box().
		Width("40").MinWidth("40").
		Height("2").MinHeight("2").
		BackgroundColor(theme.Gray200).
		FlexShrink(0).
		AlignSelf(props.AlignCenter)
}
