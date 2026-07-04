package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func Tooltip(label string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BackgroundColor(theme.Gray700)
	elem.BorderRadius(6)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetPositionType(goda.PositionTypeAbsolute)
	elem.GNode.SetPadding(goda.EdgeAll, 5)

	elem.GNode.SetFlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetJustifyContent(goda.JustifyFlexStart)
	elem.GNode.SetAlignItems(goda.AlignCenter)

	t := core.NewElement(core.TypeText)
	t.SetTextContent(label)
	t.TextColor(theme.White)
	t.FontSize(12)
	t.GNode.SetHeight(16).SetMinHeight(16)

	width := float32(len([]rune(label))) * 7
	if width < 16 {
		width = 16
	}
	t.GNode.SetMeasureFunc(func(node *goda.Node, w float32, wm goda.MeasureMode, h float32, hm goda.MeasureMode) goda.Size {
		return goda.Size{Width: width, Height: 16}
	})

	elem.Children(t)

	return elem
}

func TooltipPlacement(e *core.Element, p core.Placement) *core.Element {
	return e.TooltipPlacement(p)
}

func WithTooltip(target, tooltip *core.Element) *core.Element {
	target.GNode.SetPositionType(goda.PositionTypeRelative)
	target.Children(tooltip)
	tooltip.Visible(false)
	tooltip.SetTrigger(target)
	target.OnHover(func(entered bool) {
		tooltip.Visible(entered)
	})
	return target
}
