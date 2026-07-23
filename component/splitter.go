package component

import (
	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

func Splitter() *core.Element {
	return Box().
		BackgroundColor(theme.Gray200).
		Width("4").MinWidth("4").
		Height("100%").
		FlexShrink(0)
}

func Float() *core.Element {
	return Box().
		Position(props.PositionAbsolute)
}
