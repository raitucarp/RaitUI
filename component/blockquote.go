package component

import (
	goda "goda"
	"raitui/core"
	"raitui/theme"
)

func Blockquote(content ...string) *core.Element {
	box := core.NewElement(core.TypeBox)
	box.BackgroundColor(theme.Gray50)
	box.BorderColor(theme.Blue500)
	box.GNode.SetBorder(goda.EdgeLeft, 3)
	box.Padding("12").PaddingLeft("16")
	box.BorderRadius(4)
	box.GNode.SetFlexShrink(0)

	var children []*core.Element
	for _, line := range content {
		t := Text(line)
		t.TextColor(theme.Gray600)
		t.FontSize(14)
		children = append(children, t)
	}

	vst := VStack()
	vst.Gap("6")
	vst.Children(children...)
	box.Children(vst)

	return box
}
