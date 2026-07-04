package main

import (
	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

func main() {
	root := buildUI()
	ctx := core.NewContext(theme.Gray50)
	ctx.SetMinWindowSize(400, 300)
	ctx.Run(root, "RaitUI - Overlay Components", 700, 560)
}

func buildUI() *core.Element {
	return VStack().
		Width("100%").Height("100%").MinWidth("400").MinHeight("300").
		Padding("24").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			Text("Overlay Components").TextColor(theme.Gray800).FontSize(18),
			Separator().Width("100%"),

			Text("Menu").TextColor(theme.Gray400).FontSize(11),
			HStack().Gap("8").Children(
				Menu().Children(
					MenuItem("Dashboard"),
					MenuItem("Settings"),
					MenuItem("Profile"),
					Separator().Width("100%"),
					MenuItem("Logout"),
				),
				VStack().Gap("10").Children(
					Text("Menu items").TextColor(theme.Gray600).FontSize(13),
					Text("Hover items to see highlight").TextColor(theme.Gray400).FontSize(12),
				),
			),

			Text("Popper").TextColor(theme.Gray400).FontSize(11),
			HStack().Gap("8").Children(
				Popper().Padding("16").Children(
					Text("Popper Content").TextColor(theme.Gray700).FontSize(14),
					Text("This is a floating card.").TextColor(theme.Gray500).FontSize(12),
				),
			),

			Text("Tooltip").TextColor(theme.Gray400).FontSize(11),
			HStack().Gap("8").FlexWrap(props.WrapWrap).Children(
				Tooltip("Save changes"),
				Tooltip("Delete file"),
				Tooltip("Download report"),
			),

			Spacer(),

			Text("Dialog").TextColor(theme.Gray400).FontSize(11),
			Dialog("Confirm Action").Children(
				Text("Are you sure you want to proceed?").TextColor(theme.Gray600).FontSize(14),
				HStack().Gap("8").Children(
					Button("Cancel"),
					Button("Confirm").BackgroundColor(theme.Red500),
				),
			),
		)
}
