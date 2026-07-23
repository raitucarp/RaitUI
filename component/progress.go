package component

import (
	"raitui/core"
)

func Progress(value float32) *core.Element {
	if value < 0 {
		value = 0
	}
	if value > 100 {
		value = 100
	}

	elem := core.NewElement(core.TypeProgress).
		Width("200").MinWidth("60").
		Height("8").MinHeight("8").
		FlexShrink(0)

	elem.SetProgressValue(value / 100)
	return elem
}
