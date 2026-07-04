package main

import "raitui/theme"

func main() {
	dlg := DialogRoot()
	dlg.Title("Confirm").Body(Text("Are you sure?").TextColor(theme.Gray600).FontSize(14)).
		Footer(Spacer(), dlg.ActionTrigger("Cancel"),
			Button("OK").BackgroundColor(theme.Red500).OnClick(func() { dlg.Close() }))
	openBtn := dlg.Trigger(Button("Open Dialog"))

	tooltipTop := Tooltip("Top!").TooltipPlacement(PlaceTop)
	tooltipBottom := Tooltip("Bottom!").TooltipPlacement(PlaceBottom)
	tooltipLeft := Tooltip("Left!").TooltipPlacement(PlaceLeft)
	tooltipRight := Tooltip("Right!").TooltipPlacement(PlaceRight)

	popBottom := Popper().Padding("12").Children(Text("Bottom").TextColor(theme.Gray700).FontSize(13))
	popLeft := Popper().Padding("12").Children(Text("Left").TextColor(theme.Gray700).FontSize(13))
	popRight := Popper().Padding("12").Children(Text("Right").TextColor(theme.Gray700).FontSize(13))

	portals := Portals().Children(dlg.Element())

	app := App().
		Padding("24").Gap("16").BackgroundColor(theme.Gray50).
		Children(
			Text("Overlay Components").TextColor(theme.Gray800).FontSize(18),
			Separator().Width("100%"),
			HStack().Gap("10").Children(openBtn, Text("Dialog").TextColor(theme.Gray500).FontSize(13)),
			Separator().Width("100%"),
			Text("Tooltip Placements").TextColor(theme.Gray400).FontSize(11),
			HStack().Gap("10").Children(
				WithTooltip(Button("Top"), tooltipTop),
				WithTooltip(Button("Bottom"), tooltipBottom),
				WithTooltip(Button("Left"), tooltipLeft),
				WithTooltip(Button("Right"), tooltipRight),
			),
			Separator().Width("100%"),
			Text("Popover Placements").TextColor(theme.Gray400).FontSize(11),
			HStack().Gap("10").Children(
				PopoverAt(Button("Bottom"), popBottom, PlaceBottom),
				PopoverAt(Button("Left"), popLeft, PlaceLeft),
				PopoverAt(Button("Right"), popRight, PlaceRight),
			),
		)

	Window().Title("RaitUI - Overlays").Width(700).Height(560).MinSize(400, 300).Children(app, portals).Run()
}
