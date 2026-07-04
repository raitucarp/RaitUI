package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func Popper() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BackgroundColor(theme.White)
	elem.BorderRadius(8)
	elem.BoxShadow(0, 4, 12, 0, colorWithAlpha(theme.Black, 15))
	elem.GNode.SetBorder(goda.EdgeAll, 1)
	elem.BorderColor(theme.Gray200)
	elem.GNode.SetFlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetMinWidth(200)
	elem.GNode.SetPositionType(goda.PositionTypeAbsolute)
	elem.GNode.SetPadding(goda.EdgeAll, 8)
	return elem
}

func PopoverAt(target, pop *core.Element, placement core.Placement) *core.Element {
	target.GNode.SetPositionType(goda.PositionTypeRelative)
	target.Children(pop)
	pop.Visible(false)
	pop.SetTrigger(target)
	pop.TooltipPlacement(placement)
	target.OnClick(func() { pop.Visible(!pop.IsVisible()) })
	return target
}

func WithPopper(target, pop *core.Element) *core.Element {
	target.GNode.SetPositionType(goda.PositionTypeRelative)
	target.Children(pop)
	pop.Visible(false)
	pop.SetTrigger(target)
	target.OnClick(func() { pop.Visible(!pop.IsVisible()) })
	return target
}
