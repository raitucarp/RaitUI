package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Alert(status string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BorderRadius(6)
	elem.GNode.SetFlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetBorder(goda.EdgeLeft, 3)
	elem.Padding("12")

	switch status {
	case "success":
		elem.BackgroundColor(theme.Green50)
		elem.BorderColor(theme.Green500)
	case "warning":
		elem.BackgroundColor(theme.Yellow50)
		elem.BorderColor(theme.Yellow500)
	case "error":
		elem.BackgroundColor(theme.Red50)
		elem.BorderColor(theme.Red500)
	default:
		elem.BackgroundColor(theme.Blue50)
		elem.BorderColor(theme.Blue500)
	}

	return elem
}
