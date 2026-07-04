package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Card() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BackgroundColor(theme.White)
	elem.BorderRadius(10)
	elem.BorderColor(theme.Gray200)
	elem.BoxShadow(0, 1, 4, 0, colorWithAlpha(theme.Black, 10))
	elem.GNode.SetBorder(goda.EdgeAll, 1)
	elem.GNode.SetFlexShrink(0)
	return elem
}
