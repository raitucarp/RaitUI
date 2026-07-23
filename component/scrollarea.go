package component

import (
	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

func ScrollAreaRoot() *core.Element {
	return Box().
		Overflow(props.OverflowScroll).
		BackgroundColor(theme.White).
		BorderColor(theme.Gray200).
		BorderWidth("1").
		BorderRadius(8).
		FlexShrink(1).
		FlexGrow(1)
}

func ScrollAreaViewport() *core.Element {
	return Box().
		Overflow(props.OverflowScroll).
		FlexDirection(props.FlexDirectionColumn).
		FlexGrow(1).
		FlexShrink(1)
}

func ScrollAreaContent() *core.Element {
	return VStack()
}

func ScrollAreaScrollbar() *core.Element {
	return Box().
		Width("6").MinWidth("6").
		BackgroundColor(theme.Gray100).
		BorderRadius(3)
}

func ScrollAreaThumb() *core.Element {
	return Box().
		Width("6").MinWidth("6").
		Height("30").MinHeight("30").
		BackgroundColor(theme.Gray300).
		BorderRadius(3)
}
