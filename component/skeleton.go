package component

import (
	"raitui/core"
	"raitui/theme"
)

func Skeleton() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BackgroundColor(theme.Gray200)
	elem.BorderRadius(4)
	elem.GNode.SetWidth(200).SetMinWidth(60)
	elem.GNode.SetHeight(16).SetMinHeight(16)
	elem.GNode.SetFlexShrink(0)
	return elem
}
