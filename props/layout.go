package props

import goda "goda"

type Justify = goda.Justify
type Align = goda.Align
type FlexDirection = goda.FlexDirection
type Wrap = goda.Wrap
type PositionType = goda.PositionType
type Display = goda.Display
type Overflow = goda.Overflow
type BoxSizing = goda.BoxSizing

var (
	JustifyFlexStart    = goda.JustifyFlexStart
	JustifyCenter       = goda.JustifyCenter
	JustifyFlexEnd      = goda.JustifyFlexEnd
	JustifySpaceBetween = goda.JustifySpaceBetween
	JustifySpaceAround  = goda.JustifySpaceAround
	JustifySpaceEvenly  = goda.JustifySpaceEvenly

	AlignFlexStart    = goda.AlignFlexStart
	AlignCenter       = goda.AlignCenter
	AlignFlexEnd      = goda.AlignFlexEnd
	AlignStretch      = goda.AlignStretch
	AlignBaseline     = goda.AlignBaseline

	FlexDirectionRow            = goda.FlexDirectionRow
	FlexDirectionColumn         = goda.FlexDirectionColumn
	FlexDirectionRowReverse     = goda.FlexDirectionRowReverse
	FlexDirectionColumnReverse  = goda.FlexDirectionColumnReverse

	WrapWrap        = goda.WrapWrap
	WrapNoWrap      = goda.WrapNoWrap
	WrapWrapReverse = goda.WrapWrapReverse

	PositionStatic   = goda.PositionTypeStatic
	PositionRelative = goda.PositionTypeRelative
	PositionAbsolute = goda.PositionTypeAbsolute

	DisplayFlex = goda.DisplayFlex
	DisplayNone = goda.DisplayNone

	OverflowVisible = goda.OverflowVisible
	OverflowHidden  = goda.OverflowHidden
	OverflowScroll  = goda.OverflowScroll

	BoxSizingBorderBox  = goda.BoxSizingBorderBox
	BoxSizingContentBox = goda.BoxSizingContentBox
)
