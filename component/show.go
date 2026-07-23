package component

import "raitui/core"

func Show(condition bool) *core.Element {
	return Box().Visible(condition)
}

func VisuallyHidden() *core.Element {
	return Box().
		Width("1").
		MinWidth("1").
		Height("1").
		MinHeight("1").
		Opacity(0)
}
