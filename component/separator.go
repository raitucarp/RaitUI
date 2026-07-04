package component

import (
		"raitui/core"
)

func Separator() *core.Element {
	elem := core.NewElement(core.TypeSeparator)
	elem.GNode.SetWidth(100)
	elem.GNode.SetHeight(1)
	elem.GNode.SetFlexShrink(0)
	return elem
}
