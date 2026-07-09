package main

import "raitui/theme"

func main() {
	app := App().
		Padding("24").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			Text("Data View Components").TextColor(theme.Gray800).FontSize(18),
			Separator().Width("100%"),

			Text("Stats").TextColor(theme.Gray400).FontSize(11),
			HStack().Gap("16").Children(
				Stat().Children(
					StatLabel("Revenue"),
					StatNumber("$45,200"),
					StatHelpText("+20.1% from last month"),
				),
				Stat().Children(
					StatLabel("Users"),
					StatNumber("2,840"),
					StatHelpText("+180 new today"),
				),
				Stat().Children(
					StatLabel("Orders"),
					StatNumber("1,204"),
					StatHelpText("-3.1% from yesterday"),
				),
			),
			Separator().Width("100%"),

			Text("Table").TextColor(theme.Gray400).FontSize(11),
			Table().Children(
				TableHead().Children(
					TableRow().Children(TableHeaderCell("Name"), TableHeaderCell("Role"), TableHeaderCell("Status")),
				),
				TableBody().Children(
					TableRow().Children(TableCell("Alice"), TableCell("Admin"), TableCell("Active")),
					TableRow().Children(TableCell("Bob"), TableCell("Editor"), TableCell("Active")),
					TableRow().Children(TableCell("Charlie"), TableCell("Viewer"), TableCell("Inactive")),
				),
			),
			Separator().Width("100%"),

			Text("Timeline").TextColor(theme.Gray400).FontSize(11),
			Timeline().Children(
				TimelineItem().Children(
					TimelineDot(),
					TimelineContent().Children(
						Text("Project started").TextColor(theme.Gray800).FontSize(14),
						Text("January 2025").TextColor(theme.Gray500).FontSize(12),
					),
				),
				TimelineItem().Children(
					TimelineDot(),
					TimelineContent().Children(
						Text("First release").TextColor(theme.Gray800).FontSize(14),
						Text("March 2025").TextColor(theme.Gray500).FontSize(12),
					),
				),
				TimelineItem().Children(
					TimelineDot(),
					TimelineContent().Children(
						Text("Version 2.0 launched").TextColor(theme.Gray800).FontSize(14),
						Text("June 2025").TextColor(theme.Gray500).FontSize(12),
					),
				),
			),
		)

	Window().Title("RaitUI - Data View").Width(700).Height(580).MinSize(400, 300).Children(app).Run()
}
