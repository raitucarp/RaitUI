package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Dialog(title string) *core.Element {
	elem := core.NewElement(core.TypeDialog)
	elem.GNode.SetPositionType(goda.PositionTypeAbsolute)
	elem.GNode.SetFlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.Width("100%").Height("100%")

	backdrop := core.NewElement(core.TypeBox)
	backdrop.BackgroundColor(colorWithAlpha(theme.Black, 60))
	backdrop.Width("100%").Height("100%")
	backdrop.GNode.SetPositionType(goda.PositionTypeAbsolute)

	card := core.NewElement(core.TypeBox)
	card.BackgroundColor(theme.White)
	card.BorderRadius(12)
	card.BoxShadow(0, 20, 60, 0, colorWithAlpha(theme.Black, 40))
	card.Padding("24").Gap("16")
	card.GNode.SetMinWidth(400)
	card.GNode.SetMaxWidth(560)
	card.GNode.SetFlexDirection(goda.FlexDirectionColumn)

	if title != "" {
		t := Text(title)
		t.FontSize(18)
		t.TextColor(theme.Gray800)
		card.AppendChild(t)
	}

	elem.Children(backdrop)
	backdrop.OnClick(func() {})

	elem.AppendChild(card)

	return elem
}
