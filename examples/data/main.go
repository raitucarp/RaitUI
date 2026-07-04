package main

import (
	"raitui"
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
	return raitui.VStack().
		Width("100%").Height("100%").MinWidth("400").MinHeight("350").
		Padding("24").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			raitui.Text("Data Display Components").TextColor(theme.Gray800).FontSize(18),
			raitui.Separator().Width("100%"),

		raitui.VStack().Gap("10").Children(
			raitui.Text("Badges").TextColor(theme.Gray400).FontSize(11),
			raitui.HStack().Gap("8").FlexWrap(props.WrapWrap).Children(
				raitui.Badge("Default"),
				raitui.Badge("Blue"),
				raitui.Badge("Red"),
				raitui.Badge("Green"),
			),
		),

			raitui.VStack().Gap("10").Children(
				raitui.Text("Cards").TextColor(theme.Gray400).FontSize(11),
				raitui.HStack().Gap("12").Children(
					raitui.Card().FlexGrow(1).Padding("16").Gap("8").Children(
						raitui.Text("Card Title").TextColor(theme.Gray800).FontSize(15),
						raitui.Text("Description text goes here.").TextColor(theme.Gray500).FontSize(13),
					),
					raitui.Card().FlexGrow(1).Padding("16").Gap("8").Children(
						raitui.Text("Stats Card").TextColor(theme.Gray800).FontSize(15),
						raitui.Text("$12,430").TextColor(theme.Blue500).FontSize(24),
					),
				),
			),

			raitui.VStack().Gap("10").Children(
				raitui.Text("Alerts").TextColor(theme.Gray400).FontSize(11),
				raitui.Alert("info").Children(raitui.Text("Information: Update available.").FontSize(13)),
				raitui.Alert("success").Children(raitui.Text("Success: Deploy completed.").FontSize(13)),
				raitui.Alert("warning").Children(raitui.Text("Warning: Disk usage at 85%.").FontSize(13)),
				raitui.Alert("error").Children(raitui.Text("Error: Connection failed.").FontSize(13)),
			),

			raitui.VStack().Gap("10").Children(
				raitui.Text("Avatar & Spinner").TextColor(theme.Gray400).FontSize(11),
				raitui.HStack().Gap("10").AlignItems(props.AlignCenter).Children(
					raitui.Avatar("John Doe"),
					raitui.Avatar("Alice Smith").BackgroundColor(theme.Green500),
					raitui.Avatar("Bob").BackgroundColor(theme.Purple500),
					raitui.Spacer(),
					raitui.Spinner(),
					raitui.Text("Loading...").TextColor(theme.Gray500).FontSize(13),
				),
			),

			raitui.VStack().Gap("10").Children(
				raitui.Text("Progress").TextColor(theme.Gray400).FontSize(11),
				raitui.Progress(25),
				raitui.Progress(60),
				raitui.Progress(90),
			),
		)
}
