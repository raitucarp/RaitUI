package component

import (
		"raitui/core"
)

func Input(placeholder string) *core.Element {
	elem := core.NewElement(core.TypeInput)
	elem.SetPlaceholder(placeholder)
	elem.GNode.SetWidth(200)
	elem.GNode.SetMinHeight(36)
	elem.GNode.SetHeight(36)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func TextArea(placeholder string) *core.Element {
	elem := core.NewElement(core.TypeInput)
	elem.SetPlaceholder(placeholder)
	elem.GNode.SetWidth(300)
	elem.GNode.SetMinHeight(80)
	elem.GNode.SetHeight(80)
	elem.GNode.SetFlexShrink(0)
	return elem
}
