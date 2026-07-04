package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Badge(label string) *core.Element {
	return BadgeVariant(label, "subtle", "gray")
}

func BadgeSolid(label, colorScheme string) *core.Element {
	return BadgeVariant(label, "solid", colorScheme)
}

func BadgeSubtle(label, colorScheme string) *core.Element {
	return BadgeVariant(label, "subtle", colorScheme)
}

func BadgeOutline(label, colorScheme string) *core.Element {
	return BadgeVariant(label, "outline", colorScheme)
}

func BadgeVariant(label, variant, colorScheme string) *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetFlexDirection(goda.FlexDirectionRow)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)

	runes := len([]rune(label))
	w := float32(runes)*8 + 14
	h := float32(22)

	elem.GNode.SetWidth(w).SetMinWidth(w)
	elem.GNode.SetHeight(h).SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	bg, txt, border := badgeColors(variant, colorScheme)
	elem.BackgroundColor(bg)
	elem.BorderRadius(4)

	if variant == "outline" {
		elem.BorderColor(border)
		elem.GNode.SetBorder(goda.EdgeAll, 1)
	}

	t := Text(label)
	t.TextColor(txt)
	t.FontSize(11)
	elem.Children(t)

	return elem
}

func badgeColors(variant, scheme string) (bg, txt, border theme.Color) {
	palette := map[string]struct{ soft, strong, border theme.Color }{
		"gray":   {theme.Gray100, theme.Gray700, theme.Gray300},
		"red":    {theme.Red100, theme.Red700, theme.Red300},
		"green":  {theme.Green100, theme.Green700, theme.Green300},
		"blue":   {theme.Blue100, theme.Blue700, theme.Blue300},
		"orange": {theme.Orange100, theme.Orange700, theme.Orange300},
		"purple": {theme.Purple100, theme.Purple700, theme.Purple300},
		"yellow": {theme.Yellow100, theme.Yellow700, theme.Yellow300},
		"teal":   {theme.Teal100, theme.Teal700, theme.Teal300},
		"pink":   {theme.Pink100, theme.Pink700, theme.Pink300},
		"cyan":   {theme.Cyan100, theme.Cyan700, theme.Cyan300},
	}

	c, ok := palette[scheme]
	if !ok {
		c = palette["gray"]
	}

	switch variant {
	case "solid":
		return c.strong, theme.White, c.strong
	case "outline":
		return theme.Transparent, c.strong, c.border
	default:
		return c.soft, c.strong, c.border
	}
}
