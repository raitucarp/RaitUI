package core

import "image/color"

type ElementType int

const (
	TypeBox ElementType = iota
	TypeText
	TypeVStack
	TypeHStack
	TypeFlex
	TypeCenter
	TypeSeparator
	TypeCheckbox
	TypeSwitch
	TypeInput
	TypeButton
	TypeSpinner
	TypeProgress
	TypeAvatar
)

type Edge int

const (
	EdgeAll Edge = iota
	EdgeTop
	EdgeRight
	EdgeBottom
	EdgeLeft
	EdgeHorizontal
	EdgeVertical
)

type BorderStyle int

const (
	BorderSolid BorderStyle = iota
	BorderDashed
	BorderDotted
	BorderNone
)

type TextAlign int

const (
	AlignLeft TextAlign = iota
	AlignCenter
	AlignRight
)

type CursorType int

const (
	CursorDefault CursorType = iota
	CursorPointer
	CursorText
	CursorMove
	CursorNotAllowed
	CursorCrosshair
	CursorGrab
	CursorGrabbing
	CursorResize
)

type Shadow struct {
	OffsetX float32
	OffsetY float32
	Blur    float32
	Spread  float32
	Color   color.NRGBA
}
