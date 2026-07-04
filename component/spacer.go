package component

import (
		"raitui/core"
)

func Spacer() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetFlexGrow(1)
	return elem
}
