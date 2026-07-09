package component

import (
	goda "goda"
	"raitui/core"
)

func EmptyState() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.Padding("32").Gap("12")
	return elem
}
