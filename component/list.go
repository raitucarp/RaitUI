package component

import (
	"raitui/core"
	"raitui/theme"
)

func List(items []string, ordered bool) *core.Element {
	vst := VStack()
	vst.Gap("4")

	for i, item := range items {
		var prefix string
		if ordered {
			prefix = itoa(i+1) + ". "
		} else {
			prefix = "\u2022 "
		}

		t := Text(prefix + item)
		t.TextColor(theme.Gray700)
		t.FontSize(14)
		vst.AppendChild(t)
	}

	return vst
}

func ListItem(content string) *core.Element {
	return Text(content).
		TextColor(theme.Gray700).
		FontSize(14)
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	return s
}
