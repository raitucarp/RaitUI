package component

import (
	"raitui/core"
	"raitui/theme"
)

func Blockquote(content ...string) *core.Element {
	box := Box().
		BackgroundColor(theme.Gray50).
		BorderColor(theme.Blue500).
		BorderLeft("3").
		Padding("12").
		PaddingLeft("16").
		BorderRadius(4)

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

func BlockquoteCaption(cite string) *core.Element {
	return Text("\u2014 " + cite).
		TextColor(theme.Gray500).
		FontSize(13)
}
