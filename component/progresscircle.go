package component

import "raitui/core"

func ProgressCircle(value float32) *core.Element {
	if value < 0 {
		value = 0
	}
	if value > 100 {
		value = 100
	}
	elem := core.NewElement(core.TypeProgress)
	elem.GNode.SetWidth(48).SetMinWidth(48)
	elem.GNode.SetHeight(48).SetMinHeight(48)
	elem.GNode.SetFlexShrink(0)
	elem.SetProgressValue(value / 100)
	return elem
}
