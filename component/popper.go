package component

import (
	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

func Popper() *core.Element {
	return Box().
		BackgroundColor(theme.White).
		BorderRadius(8).
		BoxShadow(0, 4, 12, 0, colorWithAlpha(theme.Black, 15)).
		BorderWidth("1").
		BorderColor(theme.Gray200).
		FlexDirection(props.FlexDirectionColumn).
		MinWidth("200").
		Position(props.PositionAbsolute).
		Padding("8")
}

func PopoverAt(target, pop *core.Element, placement core.Placement) *core.Element {
	target.Position(props.PositionRelative)
	target.Children(pop)
	pop.Visible(false)
	pop.SetTrigger(target)
	pop.TooltipPlacement(placement)
	target.OnClick(func() { pop.Visible(!pop.IsVisible()) })
	return target
}

func WithPopper(target, pop *core.Element) *core.Element {
	target.Position(props.PositionRelative)
	target.Children(pop)
	pop.Visible(false)
	pop.SetTrigger(target)
	target.OnClick(func() { pop.Visible(!pop.IsVisible()) })
	return target
}
