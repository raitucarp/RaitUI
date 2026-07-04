package component

import (
		"raitui/core"
	"raitui/theme"
)

func Avatar(name string) *core.Element {
	size := float32(40)
	elem := core.NewElement(core.TypeAvatar)
	elem.BackgroundColor(theme.Blue500)
	elem.BorderRadius(size / 2)
	elem.GNode.SetWidth(size).SetMinWidth(size)
	elem.GNode.SetHeight(size).SetMinHeight(size)
	elem.GNode.SetFlexShrink(0)

	initials := avatarInitials(name)
	t := Text(initials)
	t.TextColor(theme.White)
	t.FontSize(float32(size) * 0.38)
	elem.Children(t)

	return elem
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
