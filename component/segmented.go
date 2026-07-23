package component

import (
	"raitui/core"
	"raitui/theme"
)

func SegmentedControl() *core.Element {
	return HStack().
		BackgroundColor(theme.Gray100).
		BorderRadius(8).
		Padding("2").
		Gap("1")
}

func Segment(label string) *core.Element {
	t := Text(label).FontSize(14).TextColor(theme.Gray600)

	seg := Center().
		PaddingX("16").PaddingY("8").
		BorderRadius(6).
		Children(t)

	seg.OnHover(func(entered bool) {
		if entered {
			seg.BackgroundColor(theme.Gray200)
		} else {
			seg.BackgroundColor(theme.Transparent)
		}
	})

	return seg
}
