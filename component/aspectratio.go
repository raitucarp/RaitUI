package component

import (
	"raitui/core"
)

func AspectRatio(ratio float32) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetAspectRatio(ratio)
	elem.GNode.SetFlexShrink(0)
	return elem
}
