package component

import "raitui/theme"

func colorWithAlpha(c theme.Color, a uint8) theme.Color {
	c.A = a
	return c
}
