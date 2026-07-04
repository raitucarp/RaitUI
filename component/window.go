package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

type WindowConfig struct {
	root    *core.Element
	ctx     *core.Context
	title   string
	width   int
	height  int
	minW    int
	minH    int
	bgColor theme.Color
}

func Window() *WindowConfig {
	return &WindowConfig{
		bgColor: theme.Gray50,
		width:   800,
		height:  600,
	}
}

func (w *WindowConfig) Title(t string) *WindowConfig { w.title = t; return w }
func (w *WindowConfig) Width(pixels int) *WindowConfig { w.width = pixels; return w }
func (w *WindowConfig) Height(pixels int) *WindowConfig { w.height = pixels; return w }
func (w *WindowConfig) MinSize(w2, h int) *WindowConfig { w.minW = w2; w.minH = h; return w }
func (w *WindowConfig) Theme(bg theme.Color) *WindowConfig { w.bgColor = bg; return w }

func (w *WindowConfig) Children(children ...*core.Element) *WindowConfig {
	w.root = core.NewElement(core.TypeVStack)
	w.root.FlexDirection(goda.FlexDirectionColumn)
	w.root.GNode.SetAlignItems(goda.AlignStretch)
	w.root.Width("100%").Height("100%")
	w.root.BackgroundColor(w.bgColor)
	for _, c := range children {
		w.root.AppendChild(c)
	}
	return w
}

func (w *WindowConfig) Run() {
	ctx := core.NewContext(w.bgColor)
	if w.minW > 0 && w.minH > 0 {
		ctx.SetMinWindowSize(w.minW, w.minH)
	}
	w.ctx = ctx
	ctx.Run(w.root, w.title, w.width, w.height)
}

func (w *WindowConfig) Ctx() *core.Context { return w.ctx }

func App() *core.Element {
	elem := core.NewElement(core.TypeVStack)
	elem.FlexDirection(goda.FlexDirectionColumn)
	return elem
}

func Portals() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetPositionType(goda.PositionTypeAbsolute)
	elem.GNode.SetFlexDirection(goda.FlexDirectionColumn)
	elem.GNode.SetJustifyContent(goda.JustifyCenter)
	elem.GNode.SetAlignItems(goda.AlignCenter)
	elem.Width("100%").Height("100%")
	return elem
}
