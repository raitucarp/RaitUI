package component

import (
	"raitui/core"
	"raitui/theme"
)

func Alert(status string) *core.Element {
	elem := VStack().
		BorderRadius(6).
		BorderLeft("3").
		Padding("12")

	switch status {
	case "success":
		elem.BackgroundColor(theme.Green50)
		elem.BorderColor(theme.Green500)
	case "warning":
		elem.BackgroundColor(theme.Yellow50)
		elem.BorderColor(theme.Yellow500)
	case "error":
		elem.BackgroundColor(theme.Red50)
		elem.BorderColor(theme.Red500)
	default:
		elem.BackgroundColor(theme.Blue50)
		elem.BorderColor(theme.Blue500)
	}

	return elem
}

func AlertTitle(title string) *core.Element {
	return Text(title).FontSize(14).TextColor(theme.Gray900)
}

func AlertDescription(desc string) *core.Element {
	return Text(desc).FontSize(13).TextColor(theme.Gray600)
}
