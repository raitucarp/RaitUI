package component

import (
	"fmt"

	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func Drawer() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BackgroundColor(theme.White)
	elem.BoxShadow(0, 0, 20, 0, colorWithAlpha(theme.Black, 20))
	elem.GNode.SetFlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetPositionType(goda.PositionTypeAbsolute)
	elem.GNode.SetHeightPercent(100)
	elem.GNode.SetMinWidth(300)
	elem.GNode.SetMaxWidth(400)
	elem.Padding("24").Gap("16")
	elem.Visible(false)
	elem.GNode.SetEdgePosition(goda.EdgeTop, 0)
	elem.GNode.SetEdgePosition(goda.EdgeRight, 0)
	return elem
}

func Collapsible() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetBorder(goda.EdgeAll, 1)
	elem.BorderColor(theme.Gray200)
	elem.BorderRadius(8)
	return elem
}

func CollapsibleTrigger(label string) *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.Padding("12").PaddingX("16")

	t := Text(label)
	t.FontSize(15)
	t.TextColor(theme.Gray700)
	t.GNode.SetFlexShrink(0)
	elem.Children(t)

	elem.OnHover(func(entered bool) {
		if entered {
			elem.BackgroundColor(theme.Gray100)
		} else {
			elem.BackgroundColor(theme.Transparent)
		}
	})

	return elem
}

func CollapsibleContent() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.Padding("12").PaddingX("16")
	elem.Gap("8")
	elem.Visible(false)
	return elem
}

func Pagination() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.Gap("4")
	return elem
}

func PageButton(label string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetWidth(36).SetMinWidth(36)
	elem.GNode.SetHeight(36).SetMinHeight(36)
	elem.BorderRadius(6)
	elem.GNode.SetFlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.GNode.SetFlexShrink(0)

	t := Text(label)
	t.FontSize(14)
	t.TextColor(theme.Gray600)
	t.GNode.SetFlexShrink(0)
	elem.Children(t)

	elem.OnHover(func(entered bool) {
		if entered {
			elem.BackgroundColor(theme.Gray100)
		} else {
			elem.BackgroundColor(theme.Transparent)
		}
	})

	return elem
}

func Steps() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	elem.Gap("0")
	return elem
}

func Step(active bool) *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.Gap("8")
	elem.FlexShrink(0)

	return elem
}

func StepIndicator(number int, active bool) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetWidth(28).SetMinWidth(28)
	elem.GNode.SetHeight(28).SetMinHeight(28)
	elem.BorderRadius(14)
	elem.GNode.SetFlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.GNode.SetFlexShrink(0)

	if active {
		elem.BackgroundColor(theme.Blue500)
	} else {
		elem.BackgroundColor(theme.Gray200)
	}

	t := Text(fmt.Sprintf("%d", number))
	t.FontSize(12)
	if active {
		t.TextColor(theme.White)
	} else {
		t.TextColor(theme.Gray600)
	}
	t.GNode.SetFlexShrink(0)
	elem.Children(t)

	return elem
}

func StepSeparator() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetWidth(40).SetMinWidth(40)
	elem.GNode.SetHeight(2).SetMinHeight(2)
	elem.BackgroundColor(theme.Gray200)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetAlignSelf(goda.AlignCenter)
	return elem
}
