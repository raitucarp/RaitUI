package component

import (
	"raitui/core"
	"raitui/props"
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
	w.root = VStack().
		AlignItems(props.AlignStretch).
		Width("100%").Height("100%").
		BackgroundColor(w.bgColor)
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
	return VStack()
}

func Portals() *core.Element {
	return Center().
		Width("100%").Height("100%")
}
