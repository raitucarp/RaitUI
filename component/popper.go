package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Popper() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BackgroundColor(theme.White)
	elem.BorderRadius(8)
	elem.BoxShadow(0, 4, 12, 0, colorWithAlpha(theme.Black, 15))
	elem.GNode.SetBorder(goda.EdgeAll, 1)
	elem.BorderColor(theme.Gray200)
	elem.GNode.SetFlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetMinWidth(200)
	return elem
}
