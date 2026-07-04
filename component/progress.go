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

	elem := core.NewElement(core.TypeProgress)
	elem.GNode.SetWidth(200).SetMinWidth(60)
	elem.GNode.SetHeight(8).SetMinHeight(8)
	elem.GNode.SetFlexShrink(0)

	elem.SetProgressValue(value / 100)
	return elem
}
