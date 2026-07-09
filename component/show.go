package component

import "raitui/core"

func Show(condition bool) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.Visible(condition)
	return elem
}

func VisuallyHidden() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetWidth(1).SetMinWidth(1)
	elem.GNode.SetHeight(1).SetMinHeight(1)
	elem.Opacity(0)
	return elem
}
