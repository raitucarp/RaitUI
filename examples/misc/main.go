package main

import "raitui/theme"

func main() {
	tg := Tag()
	tg.Children(TagLabel("Closable"), TagCloseButton(func() { tg.Visible(false) }))

	app := App().
		Padding("24").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			Text("Miscellaneous Components").TextColor(theme.Gray800).FontSize(18),
			Separator().Width("100%"),

			Text("Skeleton").TextColor(theme.Gray400).FontSize(11),
			VStack().Gap("8").Children(
				Skeleton().Width("100%"),
				Skeleton().Width("80%"),
				Skeleton().Width("60%"),
			),
			Separator().Width("100%"),

			Text("Tags").TextColor(theme.Gray400).FontSize(11),
			VStack().Gap("10").Children(
				Wrap().Gap("8").Children(
					Tag().BackgroundColor(theme.Gray100).Children(TagLabel("React").TextColor(theme.Gray700)),
					Tag().BackgroundColor(theme.Green100).Children(TagLabel("Vue").TextColor(theme.Green700)),
					Tag().BackgroundColor(theme.Blue100).Children(TagLabel("Go").TextColor(theme.Blue700)),
				),
				Wrap().Gap("8").Children(
					Tag().BackgroundColor(theme.Blue500).Children(TagLabel("React").TextColor(theme.White)),
					Tag().BackgroundColor(theme.Green500).Children(TagLabel("Vue").TextColor(theme.White)),
					Tag().BackgroundColor(theme.Cyan500).Children(TagLabel("Go").TextColor(theme.White)),
				),
				Wrap().Gap("8").Children(tg),
			),
			Separator().Width("100%"),

			Text("ScrollArea").TextColor(theme.Gray400).FontSize(11),
			ScrollAreaRoot().Width("100%").Height("140").Padding("12").Children(
				ScrollAreaViewport().Children(
					ScrollAreaContent().Gap("8").Children(
						Text("Item 1").TextColor(theme.Gray600).FontSize(14),
						Text("Item 2").TextColor(theme.Gray600).FontSize(14),
						Text("Item 3").TextColor(theme.Gray600).FontSize(14),
						Text("Item 4").TextColor(theme.Gray600).FontSize(14),
						Text("Item 5").TextColor(theme.Gray600).FontSize(14),
						Text("Item 6").TextColor(theme.Gray600).FontSize(14),
						Text("Item 7").TextColor(theme.Gray600).FontSize(14),
						Text("Item 8").TextColor(theme.Gray600).FontSize(14),
					),
				),
			),
			Separator().Width("100%"),

			Text("AspectRatio (16:9)").TextColor(theme.Gray400).FontSize(11),
			AspectRatio(16.0/9.0).Width("200").BackgroundColor(theme.Blue100).BorderRadius(8).Children(
				Text("16:9").TextColor(theme.Blue700).FontSize(14),
			),
		)

	Window().Title("RaitUI - Misc").Width(700).Height(620).MinSize(400, 300).Children(app).Run()
}
