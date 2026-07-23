package component

import (
	"raitui/core"
	"raitui/theme"
)

func Status(variant string) *core.Element {
	elem := Box().
		BorderRadius(9999).
		Width("10").
		MinWidth("10").
		Height("10").
		MinHeight("10")

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
