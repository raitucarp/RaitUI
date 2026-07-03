package theme

import "image/color"

type Color = color.NRGBA

func createRamp(base string, steps map[string]string) map[string]Color {
	result := make(map[string]Color)
	for name, hex := range steps {
		result[name] = parseHex(hex)
	}
	return result
}

func parseHex(s string) Color {
	return hexColor(s)
}

func hexColor(s string) Color {
	if len(s) == 0 || s[0] != '#' {
		return Color{R: 128, G: 128, B: 128, A: 255}
	}
	s = s[1:]
	var r, g, b uint8
	switch len(s) {
	case 6:
		v := uint32(0)
		for _, c := range s {
			v = v*16 + uint32(hexDigit(byte(c)))
		}
		r = uint8(v >> 16)
		g = uint8(v >> 8)
		b = uint8(v)
	case 3:
		v := uint32(0)
		for _, c := range s {
			v = v*16 + uint32(hexDigit(byte(c)))
		}
		r = uint8((v>>8)&0xF) * 17
		g = uint8((v>>4)&0xF) * 17
		b = uint8(v&0xF) * 17
	default:
		return Color{R: 128, G: 128, B: 128, A: 255}
	}
	return Color{R: r, G: g, B: b, A: 255}
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

var (
	Gray50  = parseHex("#F7FAFC")
	Gray100 = parseHex("#EDF2F7")
	Gray200 = parseHex("#E2E8F0")
	Gray300 = parseHex("#CBD5E1")
	Gray400 = parseHex("#A0AEC0")
	Gray500 = parseHex("#718096")
	Gray600 = parseHex("#4A5568")
	Gray700 = parseHex("#2D3748")
	Gray800 = parseHex("#1A202C")
	Gray900 = parseHex("#171923")

	Red50   = parseHex("#FFF5F5")
	Red100  = parseHex("#FED7D7")
	Red200  = parseHex("#FEB2B2")
	Red300  = parseHex("#FC8181")
	Red400  = parseHex("#F56565")
	Red500  = parseHex("#E53E3E")
	Red600  = parseHex("#C53030")
	Red700  = parseHex("#9B2C2C")
	Red800  = parseHex("#822727")
	Red900  = parseHex("#63171B")

	Orange50  = parseHex("#FFFAF0")
	Orange100 = parseHex("#FEEBC8")
	Orange200 = parseHex("#FBD38D")
	Orange300 = parseHex("#F6AD55")
	Orange400 = parseHex("#ED8936")
	Orange500 = parseHex("#DD6B20")
	Orange600 = parseHex("#C05621")
	Orange700 = parseHex("#9C4221")
	Orange800 = parseHex("#7B341E")
	Orange900 = parseHex("#652B19")

	Yellow50  = parseHex("#FFFFF0")
	Yellow100 = parseHex("#FEFCBF")
	Yellow200 = parseHex("#FAF089")
	Yellow300 = parseHex("#F6E05E")
	Yellow400 = parseHex("#ECC94B")
	Yellow500 = parseHex("#D69E2E")
	Yellow600 = parseHex("#B7791F")
	Yellow700 = parseHex("#975A16")
	Yellow800 = parseHex("#744210")
	Yellow900 = parseHex("#5F370E")

	Green50  = parseHex("#F0FFF4")
	Green100 = parseHex("#C6F6D5")
	Green200 = parseHex("#9AE6B4")
	Green300 = parseHex("#68D391")
	Green400 = parseHex("#48BB78")
	Green500 = parseHex("#38A169")
	Green600 = parseHex("#2F855A")
	Green700 = parseHex("#276749")
	Green800 = parseHex("#22543D")
	Green900 = parseHex("#1C4532")

	Teal50  = parseHex("#E6FFFA")
	Teal100 = parseHex("#B2F5EA")
	Teal200 = parseHex("#81E6D9")
	Teal300 = parseHex("#4FD1C5")
	Teal400 = parseHex("#38B2AC")
	Teal500 = parseHex("#319795")
	Teal600 = parseHex("#2C7A7B")
	Teal700 = parseHex("#285E61")
	Teal800 = parseHex("#234E52")
	Teal900 = parseHex("#1D4044")

	Blue50  = parseHex("#EBF8FF")
	Blue100 = parseHex("#BEE3F8")
	Blue200 = parseHex("#90CDF4")
	Blue300 = parseHex("#63B3ED")
	Blue400 = parseHex("#4299E1")
	Blue500 = parseHex("#3182CE")
	Blue600 = parseHex("#2B6CB0")
	Blue700 = parseHex("#2C5282")
	Blue800 = parseHex("#2A4365")
	Blue900 = parseHex("#1A365D")

	Cyan50  = parseHex("#EDFDFD")
	Cyan100 = parseHex("#C4F1F9")
	Cyan200 = parseHex("#9DECF9")
	Cyan300 = parseHex("#76E4F7")
	Cyan400 = parseHex("#0BC5EA")
	Cyan500 = parseHex("#00B5D8")
	Cyan600 = parseHex("#00A3C4")
	Cyan700 = parseHex("#0987A0")
	Cyan800 = parseHex("#086F83")
	Cyan900 = parseHex("#065666")

	Purple50  = parseHex("#FAF5FF")
	Purple100 = parseHex("#E9D8FD")
	Purple200 = parseHex("#D6BCFA")
	Purple300 = parseHex("#B794F4")
	Purple400 = parseHex("#9F7AEA")
	Purple500 = parseHex("#805AD5")
	Purple600 = parseHex("#6B46C1")
	Purple700 = parseHex("#553C9A")
	Purple800 = parseHex("#44337A")
	Purple900 = parseHex("#322659")

	Pink50  = parseHex("#FFF5F7")
	Pink100 = parseHex("#FED7E2")
	Pink200 = parseHex("#FBB6CE")
	Pink300 = parseHex("#F687B3")
	Pink400 = parseHex("#ED64A6")
	Pink500 = parseHex("#D53F8C")
	Pink600 = parseHex("#B83280")
	Pink700 = parseHex("#97266D")
	Pink800 = parseHex("#702459")
	Pink900 = parseHex("#521B41")

	White       = Color{R: 255, G: 255, B: 255, A: 255}
	Black       = Color{R: 0, G: 0, B: 0, A: 255}
	Transparent = Color{R: 0, G: 0, B: 0, A: 0}

	Colors = map[string]Color{
		"gray.50":   Gray50,
		"gray.100":  Gray100,
		"gray.200":  Gray200,
		"gray.300":  Gray300,
		"gray.400":  Gray400,
		"gray.500":  Gray500,
		"gray.600":  Gray600,
		"gray.700":  Gray700,
		"gray.800":  Gray800,
		"gray.900":  Gray900,
		"red.50":    Red50,
		"red.100":   Red100,
		"red.200":   Red200,
		"red.300":   Red300,
		"red.400":   Red400,
		"red.500":   Red500,
		"red.600":   Red600,
		"red.700":   Red700,
		"red.800":   Red800,
		"red.900":   Red900,
		"orange.50":  Orange50,
		"orange.100": Orange100,
		"orange.200": Orange200,
		"orange.300": Orange300,
		"orange.400": Orange400,
		"orange.500": Orange500,
		"orange.600": Orange600,
		"orange.700": Orange700,
		"orange.800": Orange800,
		"orange.900": Orange900,
		"yellow.50":  Yellow50,
		"yellow.100": Yellow100,
		"yellow.200": Yellow200,
		"yellow.300": Yellow300,
		"yellow.400": Yellow400,
		"yellow.500": Yellow500,
		"yellow.600": Yellow600,
		"yellow.700": Yellow700,
		"yellow.800": Yellow800,
		"yellow.900": Yellow900,
		"green.50":   Green50,
		"green.100":  Green100,
		"green.200":  Green200,
		"green.300":  Green300,
		"green.400":  Green400,
		"green.500":  Green500,
		"green.600":  Green600,
		"green.700":  Green700,
		"green.800":  Green800,
		"green.900":  Green900,
		"teal.50":    Teal50,
		"teal.100":   Teal100,
		"teal.200":   Teal200,
		"teal.300":   Teal300,
		"teal.400":   Teal400,
		"teal.500":   Teal500,
		"teal.600":   Teal600,
		"teal.700":   Teal700,
		"teal.800":   Teal800,
		"teal.900":   Teal900,
		"blue.50":    Blue50,
		"blue.100":   Blue100,
		"blue.200":   Blue200,
		"blue.300":   Blue300,
		"blue.400":   Blue400,
		"blue.500":   Blue500,
		"blue.600":   Blue600,
		"blue.700":   Blue700,
		"blue.800":   Blue800,
		"blue.900":   Blue900,
		"cyan.50":    Cyan50,
		"cyan.100":   Cyan100,
		"cyan.200":   Cyan200,
		"cyan.300":   Cyan300,
		"cyan.400":   Cyan400,
		"cyan.500":   Cyan500,
		"cyan.600":   Cyan600,
		"cyan.700":   Cyan700,
		"cyan.800":   Cyan800,
		"cyan.900":   Cyan900,
		"purple.50":  Purple50,
		"purple.100": Purple100,
		"purple.200": Purple200,
		"purple.300": Purple300,
		"purple.400": Purple400,
		"purple.500": Purple500,
		"purple.600": Purple600,
		"purple.700": Purple700,
		"purple.800": Purple800,
		"purple.900": Purple900,
		"pink.50":    Pink50,
		"pink.100":   Pink100,
		"pink.200":   Pink200,
		"pink.300":   Pink300,
		"pink.400":   Pink400,
		"pink.500":   Pink500,
		"pink.600":   Pink600,
		"pink.700":   Pink700,
		"pink.800":   Pink800,
		"pink.900":   Pink900,
	}
)
