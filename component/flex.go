package component

import (
	goda "goda"
	"raitui/core"
)

func Flex(dir goda.FlexDirection) *core.Element {
	elem := core.NewElement(core.TypeFlex)
	elem.FlexDirection(dir)
	elem.AlignItems(goda.AlignFlexStart)
	elem.FlexShrink(0)
	return elem
}
