package component

import (
	"raitui/core"
	"raitui/theme"
)

func Accordion() *core.Element {
	return VStack().
		Gap("1")
}

func AccordionItem() *core.Element {
	return VStack().
		BorderWidth("1").
		BorderColor(theme.Gray200).
		BorderRadius(8)
}

func AccordionHeader(label string) *core.Element {
	t := Text(label).FontSize(15).TextColor(theme.Gray700)

	header := HStack().
		Padding("12").PaddingX("16").
		Children(t)

	header.OnHover(func(entered bool) {
		if entered {
			header.BackgroundColor(theme.Gray100)
		} else {
			header.BackgroundColor(theme.Transparent)
		}
	})

	return header
}

func AccordionPanel() *core.Element {
	return VStack().
		Padding("12").PaddingX("16").
		Gap("8").
		Visible(false)
}
