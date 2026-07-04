package props

import (
	"image/color"
	"raitui/theme"
)

type ShadowPreset struct {
	OffsetX, OffsetY, Blur, Spread float32
	Color                          color.NRGBA
}

var (
	ShadowXS   = ShadowPreset{0, 1, 2, 0, colorWithAlpha(theme.Black, 10)}
	ShadowSM   = ShadowPreset{0, 1, 3, 0, colorWithAlpha(theme.Black, 15)}
	ShadowMD   = ShadowPreset{0, 4, 6, -1, colorWithAlpha(theme.Black, 15)}
	ShadowLG   = ShadowPreset{0, 10, 15, -3, colorWithAlpha(theme.Black, 15)}
	ShadowXL   = ShadowPreset{0, 20, 25, -5, colorWithAlpha(theme.Black, 15)}
	Shadow2XL  = ShadowPreset{0, 25, 50, -12, colorWithAlpha(theme.Black, 25)}
	ShadowInner = ShadowPreset{0, 2, 4, 0, colorWithAlpha(theme.Black, 10)}
)

func colorWithAlpha(c theme.Color, a uint8) theme.Color {
	c.A = a
	return c
}
