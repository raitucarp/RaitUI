package component

import (
	goda "goda"
	"raitui/core"
)

func Center() *core.Element {
	elem := core.NewElement(core.TypeCenter)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.GNode.SetFlexShrink(0)
	return elem
}
