package component

import (
	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

func Tooltip(label string) *core.Element {
	elem := core.NewElement(core.TypeTooltip)
	elem.BackgroundColor(theme.Gray700)
	elem.BorderRadius(6)
	elem.FlexShrink(0)
	elem.Position(props.PositionAbsolute)
	elem.Padding("5")
	elem.FlexDirection(props.FlexDirectionRow)
	elem.JustifyContent(props.JustifyFlexStart)
	elem.AlignItems(props.AlignCenter)

	t := Text(label).TextColor(theme.White).FontSize(12)
	elem.Children(t)

	return elem
}

func TooltipPlacement(e *core.Element, p core.Placement) *core.Element {
	return e.TooltipPlacement(p)
}

func WithTooltip(target, tooltip *core.Element) *core.Element {
	target.Position(props.PositionRelative)
	target.Children(tooltip)
	tooltip.Visible(false)
	tooltip.SetTrigger(target)
	target.OnHover(func(entered bool) {
		tooltip.Visible(entered)
	})
	return target
}
