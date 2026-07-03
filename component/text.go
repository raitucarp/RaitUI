package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

func Heading(content string, level int) *core.Element {
	sizes := []float32{24, 20, 18, 16, 14, 12}
	wgts := []float32{24, 20, 18, 16, 14, 12}
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	_ = wgts

	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)
	elem.TextColor(theme.Gray800)

	size := sizes[level-1]
	elem.FontSize(size)

	runes := len([]rune(content))
	w := float32(runes) * size * 0.65
	h := size * 1.5
	if w < 50 {
		w = 50
	}

	elem.GNode.SetWidth(w)
	elem.GNode.SetMinWidth(w)
	elem.GNode.SetHeight(h)
	elem.GNode.SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	return elem
}

func Code(content string) *core.Element {
	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)
	elem.FontSize(13)
	elem.TextColor(theme.Gray700)
	elem.BackgroundColor(theme.Gray100)
	elem.BorderColor(theme.Gray200)
	elem.BorderRadius(4)

	runes := len([]rune(content))
	w := float32(runes)*9 + 16
	h := float32(24)

	elem.GNode.SetWidth(w)
	elem.GNode.SetMinWidth(w)
	elem.GNode.SetHeight(h)
	elem.GNode.SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetBorder(goda.EdgeAll, 1)

	return elem
}

func Link(content string) *core.Element {
	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)
	elem.TextColor(theme.Blue500)
	elem.FontSize(14)
	elem.SetUnderline(true)

	runes := len([]rune(content))
	w := float32(runes) * 9
	h := float32(22)

	elem.GNode.SetWidth(w)
	elem.GNode.SetMinWidth(w)
	elem.GNode.SetHeight(h)
	elem.GNode.SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	return elem
}

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
		t := core.NewElement(core.TypeText)
		t.SetTextContent(line)
		t.TextColor(theme.Gray600)
		t.FontSize(14)

		runes := len([]rune(line))
		w := float32(runes) * 9
		h := float32(22)
		t.GNode.SetWidth(w)
		t.GNode.SetMinWidth(w)
		t.GNode.SetHeight(h)
		t.GNode.SetMinHeight(h)
		t.GNode.SetFlexShrink(0)

		children = append(children, t)
	}

	vst := core.NewElement(core.TypeVStack)
	vst.FlexDirection(goda.FlexDirectionColumn)
	vst.GNode.SetAlignItems(goda.AlignFlexStart)
	vst.Gap("6")
	vst.Children(children...)

	box.Children(vst)
	return box
}

func List(items []string, ordered bool) *core.Element {
	vst := core.NewElement(core.TypeVStack)
	vst.FlexDirection(goda.FlexDirectionColumn)
	vst.GNode.SetAlignItems(goda.AlignFlexStart)
	vst.Gap("4")
	vst.GNode.SetFlexShrink(0)

	for i, item := range items {
		var prefix string
		if ordered {
			prefix = itoa(i+1) + ". "
		} else {
			prefix = "\u2022 "
		}
		label := prefix + item

		t := core.NewElement(core.TypeText)
		t.SetTextContent(label)
		t.TextColor(theme.Gray700)
		t.FontSize(14)

		runes := len([]rune(label))
		w := float32(runes) * 9
		h := float32(22)
		t.GNode.SetWidth(w)
		t.GNode.SetMinWidth(w)
		t.GNode.SetHeight(h)
		t.GNode.SetMinHeight(h)
		t.GNode.SetFlexShrink(0)

		vst.AppendChild(t)
	}

	return vst
}

func Kbd(content string) *core.Element {
	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)
	elem.TextColor(theme.Gray700)
	elem.BackgroundColor(theme.Gray50)
	elem.BorderColor(theme.Gray300)
	elem.BorderRadius(4)
	elem.FontSize(12)

	runes := len([]rune(content))
	w := float32(runes)*8 + 12
	h := float32(22)

	elem.GNode.SetWidth(w)
	elem.GNode.SetMinWidth(w)
	elem.GNode.SetHeight(h)
	elem.GNode.SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)
	elem.GNode.SetBorder(goda.EdgeAll, 1)
	elem.BoxShadow(0, 1, 0, 0, theme.Gray300)

	return elem
}

func Mark(content string) *core.Element {
	elem := core.NewElement(core.TypeText)
	elem.SetTextContent(content)
	elem.TextColor(theme.Gray800)
	elem.BackgroundColor(theme.Yellow100)
	elem.BorderRadius(2)
	elem.FontSize(14)

	runes := len([]rune(content))
	w := float32(runes)*9 + 6
	h := float32(22)

	elem.GNode.SetWidth(w)
	elem.GNode.SetMinWidth(w)
	elem.GNode.SetHeight(h)
	elem.GNode.SetMinHeight(h)
	elem.GNode.SetFlexShrink(0)

	return elem
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	return s
}
