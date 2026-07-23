package component

import (
	"raitui/core"
	"raitui/theme"
)

func EmptyState() *core.Element {
	return Center().
		Padding("32").
		Gap("12")
}

func EmptyStateTitle(title string) *core.Element {
	return Text(title).
		FontSize(18).
		TextColor(theme.Gray800)
}

func EmptyStateDescription(desc string) *core.Element {
	return Text(desc).
		FontSize(14).
		TextColor(theme.Gray500)
}
