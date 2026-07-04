package component

import (
	goda "goda"
	"raitui/core"
)

func VStack() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	elem.GNode.SetFlexShrink(0)
	return elem
}
