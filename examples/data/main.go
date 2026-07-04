package main

import (
	"raitui/core"
	"raitui/props"
	"raitui/theme"
)

func main() {
	root := buildUI()
	ctx := core.NewContext(theme.Gray50)
	ctx.SetMinWindowSize(400, 350)
	ctx.Run(root, "RaitUI - Data Components", 700, 640)
}

func buildUI() *core.Element {
	return VStack().
		Width("100%").Height("100%").MinWidth("400").MinHeight("350").
		Padding("24").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			Text("Data Display Components").TextColor(theme.Gray800).FontSize(18),
			Separator().Width("100%"),

		VStack().Gap("10").Children(
			Text("Badges").TextColor(theme.Gray400).FontSize(11),
			HStack().Gap("8").FlexWrap(props.WrapWrap).Children(
				Badge("Default"),
				Badge("Blue"),
				Badge("Red"),
				Badge("Green"),
			),
		),

			VStack().Gap("10").Children(
				Text("Cards").TextColor(theme.Gray400).FontSize(11),
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
			),

			VStack().Gap("10").Children(
				Text("Alerts").TextColor(theme.Gray400).FontSize(11),
				Alert("info").Children(Text("Information: Update available.").FontSize(13)),
				Alert("success").Children(Text("Success: Deploy completed.").FontSize(13)),
				Alert("warning").Children(Text("Warning: Disk usage at 85%.").FontSize(13)),
				Alert("error").Children(Text("Error: Connection failed.").FontSize(13)),
			),

			VStack().Gap("10").Children(
				Text("Avatar & Spinner").TextColor(theme.Gray400).FontSize(11),
				HStack().Gap("10").AlignItems(props.AlignCenter).Children(
					Avatar("John Doe"),
					Avatar("Alice Smith").BackgroundColor(theme.Green500),
					Avatar("Bob").BackgroundColor(theme.Purple500),
					Spacer(),
					Spinner(),
					Text("Loading...").TextColor(theme.Gray500).FontSize(13),
				),
			),

			VStack().Gap("10").Children(
				Text("Progress").TextColor(theme.Gray400).FontSize(11),
				Progress(25),
				Progress(60),
				Progress(90),
			),
		)
}
