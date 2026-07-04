package main

import (
	"raitui"
	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

var (
	BadgeSolid   = raitui.BadgeSolid
	BadgeSubtle  = raitui.BadgeSubtle
	BadgeOutline = raitui.BadgeOutline
)

func main() {
	root := buildUI()
	ctx := core.NewContext(theme.Gray50)
	ctx.SetMinWindowSize(450, 400)
	ctx.Run(root, "RaitUI - Data Components", 720, 700)
}

func buildUI() *core.Element {
	return VStack().
		Width("100%").Height("100%").MinWidth("450").MinHeight("400").
		Padding("24").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			Text("Data Display Components").TextColor(theme.Gray800).FontSize(18),
			Separator().Width("100%"),

			sectionLabel("Badge — Subtle"),
			HStack().Gap("6").FlexWrap(props.WrapWrap).Children(
				BadgeSubtle("Gray", "gray"),
				BadgeSubtle("Red", "red"),
				BadgeSubtle("Green", "green"),
				BadgeSubtle("Blue", "blue"),
				BadgeSubtle("Orange", "orange"),
				BadgeSubtle("Purple", "purple"),
				BadgeSubtle("Yellow", "yellow"),
				BadgeSubtle("Teal", "teal"),
			),

			sectionLabel("Badge — Solid"),
			HStack().Gap("6").FlexWrap(props.WrapWrap).Children(
				BadgeSolid("Gray", "gray"),
				BadgeSolid("Red", "red"),
				BadgeSolid("Green", "green"),
				BadgeSolid("Blue", "blue"),
				BadgeSolid("Orange", "orange"),
				BadgeSolid("Purple", "purple"),
			),

			sectionLabel("Badge — Outline"),
			HStack().Gap("6").FlexWrap(props.WrapWrap).Children(
				BadgeOutline("Gray", "gray"),
				BadgeOutline("Red", "red"),
				BadgeOutline("Green", "green"),
				BadgeOutline("Blue", "blue"),
				BadgeOutline("Purple", "purple"),
			),

			sectionLabel("Cards"),
			HStack().Gap("12").Children(
				Card().FlexGrow(1).Padding("16").Gap("8").Children(
					Text("Card Title").TextColor(theme.Gray800).FontSize(15),
					Text("Description text goes here.").TextColor(theme.Gray500).FontSize(13),
				),
				Card().FlexGrow(1).Padding("16").Gap("8").Children(
					Text("Stats Card").TextColor(theme.Gray800).FontSize(15),
					Text("$12,430").TextColor(theme.Blue500).FontSize(24),
				),
			),

			sectionLabel("Alerts"),
			Alert("info").Children(Text("Information: Update available.").FontSize(13)),
			Alert("success").Children(Text("Success: Deploy completed.").FontSize(13)),
			Alert("warning").Children(Text("Warning: Disk usage at 85%.").FontSize(13)),
			Alert("error").Children(Text("Error: Connection failed.").FontSize(13)),

			sectionLabel("Avatar & Spinner"),
			HStack().Gap("12").AlignItems(props.AlignCenter).Children(
				Avatar("John Doe"),
				Avatar("Alice Smith").BackgroundColor(theme.Green500),
				Avatar("Bob").BackgroundColor(theme.Purple500),
				Spacer(),
				Spinner(),
				Text("Loading...").TextColor(theme.Gray500).FontSize(13),
			),

			sectionLabel("Progress"),
			Progress(25),
			Progress(60),
			Progress(90),
		)
}

func sectionLabel(text string) *core.Element {
	return Text(text).TextColor(theme.Gray400).FontSize(11)
}
