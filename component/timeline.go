package component

import (
	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

func Timeline() *core.Element {
	return VStack().
		Gap("0")
}

func TimelineItem() *core.Element {
	return HStack().
		Gap("12")
}

func TimelineDot() *core.Element {
	return Box().
		BackgroundColor(theme.Blue500).
		BorderRadius(9999).
		Width("16").MinWidth("16").
		Height("16").MinHeight("16").
		FlexShrink(0).
		AlignSelf(props.AlignCenter)
}

func TimelineContent() *core.Element {
	return VStack().
		PaddingY("4").
		PaddingBottom("24").
		PaddingLeft("16").
		Gap("4").
		FlexGrow(1).
		BorderLeft("2").
		BorderColor(theme.Gray200)
}
