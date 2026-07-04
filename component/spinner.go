package component

import (
		"raitui/core"
)

func Spinner() *core.Element {
	elem := core.NewElement(core.TypeSpinner)
	elem.GNode.SetWidth(24).SetMinWidth(24)
	elem.GNode.SetHeight(24).SetMinHeight(24)
	elem.GNode.SetFlexShrink(0)
	return elem
}
