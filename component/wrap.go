package component

import (
	goda "goda"
	"raitui/core"
)

func Wrap() *core.Element {
	elem := core.NewElement(core.TypeHStack)
	elem.FlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetFlexWrap(goda.WrapWrap)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	return elem
}
