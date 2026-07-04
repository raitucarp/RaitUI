package core

import (
	"image/color"
	"raitui/theme"
)

var (
	ColorTransparent = theme.Transparent
	ColorBlack       = theme.Black
	ColorWhite       = theme.White

	ColorTextPrimary   = theme.Gray800
	ColorTextSecondary = theme.Gray500
	ColorTextMuted     = theme.Gray400
	ColorTextInverse   = theme.White

	ColorBorder       = theme.Gray300
	ColorBorderFocus  = theme.Blue500

	ColorBgInput    = theme.White
	ColorBgTrack    = theme.Gray100
	ColorBgAccent   = theme.Blue500
	ColorBgSurface  = theme.White

	ColorAccent        = theme.Blue500
	ColorAccentHover   = theme.Blue600
	ColorAccentSubtle  = theme.Blue50

	ColorSuccess = theme.Green500
	ColorWarning = theme.Yellow500
	ColorError   = theme.Red500

	ColorDebugRed = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
)
