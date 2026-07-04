package component

import (
	goda "goda"
	"raitui/core"
)

func Flex(dir goda.FlexDirection) *core.Element {
	elem := core.NewElement(core.TypeFlex)
	elem.FlexDirection(dir)
	elem.GNode.SetAlignItems(goda.AlignFlexStart)
	elem.GNode.SetFlexShrink(0)
	return elem
}
