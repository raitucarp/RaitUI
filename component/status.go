package component

import (
	"raitui/core"
	"raitui/theme"
)

func Status(variant string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BorderRadius(9999)
	elem.GNode.SetWidth(10).SetMinWidth(10)
	elem.GNode.SetHeight(10).SetMinHeight(10)
	elem.GNode.SetFlexShrink(0)

	switch variant {
	case "success":
		elem.BackgroundColor(theme.Green500)
	case "warning":
		elem.BackgroundColor(theme.Yellow500)
	case "error":
		elem.BackgroundColor(theme.Red500)
	case "info":
		elem.BackgroundColor(theme.Blue500)
	default:
		elem.BackgroundColor(theme.Gray400)
	}

	return elem
}
