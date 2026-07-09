package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Splitter() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BackgroundColor(theme.Gray200)
	elem.GNode.SetWidth(4).SetMinWidth(4)
	elem.GNode.SetHeightPercent(100)
	elem.GNode.SetFlexShrink(0)
	return elem
}

func Float() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetPositionType(goda.PositionTypeAbsolute)
	elem.GNode.SetFlexShrink(0)
	return elem
}
