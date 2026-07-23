package component

import (
	"raitui/core"
	"raitui/theme"
)

func Avatar(name string) *core.Element {
	initials := avatarInitials(name)
	t := Text(initials).TextColor(theme.White).FontSize(15)

	return Center().
		Width("40").Height("40").
		MinWidth("40").MinHeight("40").
		BackgroundColor(theme.Blue500).
		BorderRadius(20).
		Children(t)
}

func avatarInitials(name string) string {
	rs := []rune(name)
	if len(rs) == 0 {
		return "?"
	}
	var parts []string
	current := ""
	for _, r := range rs {
		if r == ' ' {
			if current != "" {
				parts = append(parts, current)
				current = ""
			}
		} else {
			current += string(r)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}
	if len(parts) == 0 {
		return string(rs[0])
	}
	if len(parts) == 1 {
		return string([]rune(parts[0])[0])
	}
	return string([]rune(parts[0])[0]) + string([]rune(parts[len(parts)-1])[0])
}
