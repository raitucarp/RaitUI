package core

import (
	"image/color"
	"math"
)

func RGBA(r, g, b uint8) color.NRGBA { return color.NRGBA{R: r, G: g, B: b, A: 255} }

func RGBAHex(s string) color.NRGBA { return hexColor(s) }

func hexColor(hex string) color.NRGBA {
	if len(hex) == 0 || hex[0] != '#' {
		return color.NRGBA{R: 128, G: 128, B: 128, A: 255}
	}
	hex = hex[1:]
	var r, g, b uint8
	switch len(hex) {
	case 6:
		v := uint32(0)
		for _, c := range hex {
			v = v*16 + uint32(hexDigit(byte(c)))
		}
		r = uint8(v >> 16)
		g = uint8(v >> 8)
		b = uint8(v)
	case 3:
		v := uint32(0)
		for _, c := range hex {
			v = v*16 + uint32(hexDigit(byte(c)))
		}
		r = uint8((v>>8)&0xF) * 17
		g = uint8((v>>4)&0xF) * 17
		b = uint8(v&0xF) * 17
	default:
		return color.NRGBA{R: 128, G: 128, B: 128, A: 255}
	}
	return color.NRGBA{R: r, G: g, B: b, A: 255}
}

func hexDigit(c byte) uint8 {
	switch {
	case c >= '0' && c <= '9':
		return c - '0'
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func darkenColor(c color.NRGBA, amount float64) color.NRGBA {
	return color.NRGBA{
		R: uint8(math.Max(0, float64(c.R)*(1-amount))),
		G: uint8(math.Max(0, float64(c.G)*(1-amount))),
		B: uint8(math.Max(0, float64(c.B)*(1-amount))),
		A: c.A,
	}
}
