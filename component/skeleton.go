package component

import (
	"raitui/core"
	"raitui/theme"
)

func Skeleton() *core.Element {
	return Box().
		BackgroundColor(theme.Gray200).
		BorderRadius(4).
		Width("200").
		MinWidth("60").
		Height("16").
		MinHeight("16")
}
