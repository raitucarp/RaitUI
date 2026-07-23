package component

import (
	"raitui/core"
	"raitui/theme"
)

func Card() *core.Element {
	return VStack().
		BackgroundColor(theme.White).
		BorderRadius(10).
		BorderColor(theme.Gray200).
		BoxShadow(0, 1, 4, 0, colorWithAlpha(theme.Black, 10)).
		BorderWidth("1")
}

func CardHeader() *core.Element {
	return Box().
		Padding("16").
		PaddingBottom("0")
}

func CardBody() *core.Element {
	return Box().
		Padding("16")
}

func CardFooter() *core.Element {
	return HStack().
		Padding("16").
		PaddingTop("0").
		Gap("12")
}
