package component

import (
	goda "goda"
	"raitui/core"
)

func Container() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	elem.GNode.SetFlexShrink(0)
	return elem
}
