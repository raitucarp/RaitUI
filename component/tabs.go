package component

import (
	"raitui/core"
	"raitui/theme"
)

func Tabs() *core.Element {
	return VStack()
}

func TabList() *core.Element {
	return HStack().
		Gap("2").
		BorderBottom("2").
		BorderColor(theme.Gray200)
}

func Tab(label string) *core.Element {
	t := Text(label).FontSize(14).TextColor(theme.Gray500)

	tab := Center().
		PaddingX("16").
		PaddingY("10").
		BorderRadius(6).
		Children(t)

	tab.OnHover(func(entered bool) {
		if entered {
			tab.BackgroundColor(theme.Gray100)
		} else {
			tab.BackgroundColor(theme.Transparent)
		}
	})

	return tab
}

func TabPanels() *core.Element {
	return VStack().
		Padding("16")
}

func TabPanel() *core.Element {
	return VStack().
		Gap("8").
		Visible(false)
}
