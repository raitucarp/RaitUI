package component

import (
	goda "goda"
	"raitui/core"
)

func HStack() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	elem.GNode.SetFlexShrink(0)
	return elem
}
