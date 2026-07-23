package component

import (
	"raitui/core"
	"raitui/theme"
)

func Breadcrumb() *core.Element {
	return HStack().
		Gap("4")
}

func BreadcrumbItem(label string) *core.Element {
	t := Text(label).FontSize(14).TextColor(theme.Gray500)

	return HStack().
		Gap("4").
		Children(t)
}

func BreadcrumbLink(label string) *core.Element {
	return Link(label)
}

func BreadcrumbSeparator() *core.Element {
	return Text("/").
		FontSize(14).
		TextColor(theme.Gray400)
}
