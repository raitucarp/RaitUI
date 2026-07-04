package main

import (
	"fmt"

	"raitui/core"
	"raitui/theme"
)

func main() {
	root := core.NewElement(core.TypeVStack)
	root.FlexDirection(0).Width("100%").Height("100%")
	root.Padding("24").Gap("16").BackgroundColor(theme.Gray50)

	btn := core.NewElement(core.TypeButton)
	btn.BackgroundColor(theme.Blue500).BorderRadius(6)
	btn.GNode.SetWidth(120).SetMinWidth(120)
	btn.GNode.SetHeight(36).SetMinHeight(36)
	btn.GNode.SetFlexShrink(0)
	btn.TextColor(theme.White).FontSize(14)
	btn.SetTextContent("Click Me")

	st := btn.State()
	st.Set("count", 0)

	btn.OnClick(func() {
		c := st.Int("count", 0) + 1
		st.Set("count", c)
		btn.SetTextContent(fmt.Sprintf("%d clicks", c))
	})

	root.Children(btn)

	ctx := core.NewContext(theme.Gray50)
	ctx.Debug = true
	ctx.Run(root, "Click Test", 400, 300)
}
