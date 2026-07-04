package theme

import "image/color"

type Theme struct {
	Colors     map[string]Color
	Scheme     string
	Semantic   map[string]Color
}

func Default() *Theme {
	return &Theme{
		Scheme: "light",
		Colors: DefaultColors(),
		Semantic: DefaultSemantic(),
	}
}

func DefaultColors() map[string]Color {
	return map[string]Color{
		"white":       White,
		"black":       Black,
		"transparent": Transparent,

		"gray.50":  Gray50,  "gray.100": Gray100, "gray.200": Gray200,
		"gray.300": Gray300, "gray.400": Gray400, "gray.500": Gray500,
		"gray.600": Gray600, "gray.700": Gray700, "gray.800": Gray800,
		"gray.900": Gray900,

		"red.50":  Red50,  "red.100": Red100, "red.200": Red200,
		"red.300": Red300, "red.400": Red400, "red.500": Red500,
		"red.600": Red600, "red.700": Red700, "red.800": Red800,
		"red.900": Red900,

		"orange.50":  Orange50,  "orange.100": Orange100, "orange.200": Orange200,
		"orange.300": Orange300, "orange.400": Orange400, "orange.500": Orange500,
		"orange.600": Orange600, "orange.700": Orange700, "orange.800": Orange800,
		"orange.900": Orange900,

		"yellow.50":  Yellow50,  "yellow.100": Yellow100, "yellow.200": Yellow200,
		"yellow.300": Yellow300, "yellow.400": Yellow400, "yellow.500": Yellow500,
		"yellow.600": Yellow600, "yellow.700": Yellow700, "yellow.800": Yellow800,
		"yellow.900": Yellow900,

		"green.50":  Green50,  "green.100": Green100, "green.200": Green200,
		"green.300": Green300, "green.400": Green400, "green.500": Green500,
		"green.600": Green600, "green.700": Green700, "green.800": Green800,
		"green.900": Green900,

		"teal.50":  Teal50,  "teal.100": Teal100, "teal.200": Teal200,
		"teal.300": Teal300, "teal.400": Teal400, "teal.500": Teal500,
		"teal.600": Teal600, "teal.700": Teal700, "teal.800": Teal800,
		"teal.900": Teal900,

		"blue.50":  Blue50,  "blue.100": Blue100, "blue.200": Blue200,
		"blue.300": Blue300, "blue.400": Blue400, "blue.500": Blue500,
		"blue.600": Blue600, "blue.700": Blue700, "blue.800": Blue800,
		"blue.900": Blue900,

		"cyan.50":  Cyan50,  "cyan.100": Cyan100, "cyan.200": Cyan200,
		"cyan.300": Cyan300, "cyan.400": Cyan400, "cyan.500": Cyan500,
		"cyan.600": Cyan600, "cyan.700": Cyan700, "cyan.800": Cyan800,
		"cyan.900": Cyan900,

		"purple.50":  Purple50,  "purple.100": Purple100, "purple.200": Purple200,
		"purple.300": Purple300, "purple.400": Purple400, "purple.500": Purple500,
		"purple.600": Purple600, "purple.700": Purple700, "purple.800": Purple800,
		"purple.900": Purple900,

		"pink.50":  Pink50,  "pink.100": Pink100, "pink.200": Pink200,
		"pink.300": Pink300, "pink.400": Pink400, "pink.500": Pink500,
		"pink.600": Pink600, "pink.700": Pink700, "pink.800": Pink800,
		"pink.900": Pink900,
	}
}

func DefaultSemantic() map[string]Color {
	return map[string]Color{
		"primary":    Blue500,
		"secondary":  Gray300,
		"success":    Green500,
		"warning":    Yellow500,
		"error":      Red500,
		"info":       Blue400,
		"background": Gray50,
		"foreground": Gray800,
		"muted":      Gray400,
		"border":     Gray200,
		"surface":    White,
	}
}

func Dark() *Theme {
	return &Theme{
		Scheme: "dark",
		Colors: DefaultColors(),
		Semantic: map[string]Color{
			"primary":    Blue300,
			"secondary":  Gray500,
			"success":    Green300,
			"warning":    Yellow300,
			"error":      Red300,
			"info":       Blue200,
			"background": Gray800,
			"foreground": Gray50,
			"muted":      Gray500,
			"border":     Gray600,
			"surface":    Gray700,
		},
	}
}

var current = Default()

func Current() *Theme { return current }

func SetTheme(t *Theme) { current = t }

func Token(key string) Color {
	if c, ok := current.Colors[key]; ok {
		return c
	}
	if c, ok := current.Semantic[key]; ok {
		return c
	}
	return color.NRGBA{R: 128, G: 128, B: 128, A: 255}
}

func T(key string) Color { return Token(key) }
