package main

import (
	"fmt"

	"raitui"
	"raitui/core"
	"raitui/theme"
)

func main() {
	raitui.SetState("solidClicks", 0)
	raitui.SetState("outlineClicks", 0)

	root := buildUI()
	ctx := core.NewContext(theme.Gray50)
	ctx.SetMinWindowSize(400, 350)
	ctx.Run(root, "RaitUI - Form Components", 650, 600)
}

func buildUI() *core.Element {
	solidClicks := raitui.StateInt("solidClicks", 0)
	setSolid := raitui.Setter[int]("solidClicks")
	outlineClicks := raitui.StateInt("outlineClicks", 0)
	setOutline := raitui.Setter[int]("outlineClicks")

	solidBtn := raitui.Button(fmt.Sprintf("Solid: %d", solidClicks()))
	solidBtn.OnClick(func() {
		setSolid(solidClicks() + 1)
		solidBtn.SetTextContent(fmt.Sprintf("Solid: %d", solidClicks()))
	})

	outBtn := raitui.OutlineButton(fmt.Sprintf("Outline: %d", outlineClicks()))
	outBtn.OnClick(func() {
		setOutline(outlineClicks() + 1)
		outBtn.SetTextContent(fmt.Sprintf("Outline: %d", outlineClicks()))
	})

	var _ = solidBtn
	var _ = outBtn

	return raitui.VStack().
		Width("100%").Height("100%").
		MinWidth("400").MinHeight("350").
		Padding("24").Gap("16").
		BackgroundColor(theme.Gray50).
		Children(
			raitui.Text("Form Components").TextColor(theme.Gray800).FontSize(18),
			raitui.Separator().Width("100%"),

			raitui.VStack().Gap("10").Children(
				raitui.Text("Buttons").TextColor(theme.Gray400).FontSize(11),
				raitui.HStack().Gap("10").Children(solidBtn, outBtn),
			),

			raitui.VStack().Gap("10").Children(
				raitui.Text("Checkbox & Switch").TextColor(theme.Gray400).FontSize(11),
				raitui.Checkbox("Accept terms", false),
				raitui.Checkbox("Subscribe newsletter", true),
				raitui.Switch("Dark mode", false),
				raitui.Switch("Notifications", true),
			),

			raitui.VStack().Gap("10").Children(
				raitui.Text("Input (click then type)").TextColor(theme.Gray400).FontSize(11),
				raitui.Input("Enter your name"),
				raitui.Input("email@example.com"),
			),

			raitui.VStack().Gap("10").Children(
				raitui.Text("TextArea (click then type)").TextColor(theme.Gray400).FontSize(11),
				raitui.TextArea("Write your message..."),
			),
		)
}
