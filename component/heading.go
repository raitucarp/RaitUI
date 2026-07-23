package component

import (
	"raitui/core"
	"raitui/theme"
)

func Heading(content string, level int) *core.Element {
	sizes := []float32{24, 20, 18, 16, 14, 12}
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}

	return Text(content).
		FontSize(sizes[level-1]).
		TextColor(theme.Gray800)
}
