package core

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"

	goda "goda"

	"raitui/core/keyboard"
)

type Context struct {
	Root    *Element
	Game    *game
	font    font.Face
	fontSm  font.Face
	fontLg  font.Face
	winW    int
	winH    int
	minW    int
	minH    int
	bgColor color.NRGBA
	mouseX  int
	mouseY  int
	Debug   bool
	focused *Element
	frame   int
	kb      *keyboard.Handler
	layers  []*layer
	modal   *Element
}

type layer struct {
	elem   *Element
	gNode  *goda.Node
}

func (c *Context) SetModal(elem *Element) { c.modal = elem }
func (c *Context) IsModal() bool          { return c.modal != nil }

func (c *Context) Keyboard() *keyboard.Handler { return c.kb }

func (c *Context) PushLayer(elem *Element) {
	if elem.GNode.GetParent() != nil {
		elem.GNode.GetParent().RemoveChildNode(elem.GNode)
	}
	gn := goda.New()
	gn.InsertChildNode(elem.GNode, 0)
	c.layers = append(c.layers, &layer{elem: elem, gNode: gn})
	if c.winW > 0 {
		c.rebuildLayers()
	}
}

func (c *Context) PopLayer(elem *Element) {
	for i, l := range c.layers {
		if l.elem == elem {
			c.layers = append(c.layers[:i], c.layers[i+1:]...)
			c.rebuildLayers()
			return
		}
	}
}

func (c *Context) rebuildLayers() {
	for _, l := range c.layers {
		goda.CalculateNodeLayout(l.gNode, float32(c.winW), float32(c.winH), goda.DirectionLTR)
	}
}

func (c *Context) TopLayer() *Element {
	if len(c.layers) == 0 {
		return nil
	}
	return c.layers[len(c.layers)-1].elem
}

func (c *Context) HasLayer() bool { return len(c.layers) > 1 }
type RenderFunc func(ctx *RenderCtx)

type RenderCtx struct {
	Screen  *ebiten.Image
	Elem    *Element
	X       float32
	Y       float32
	W       float32
	H       float32
	Font    font.Face
	Hovered bool
	Focused bool
	Frame   int
}

type game struct {
	ctx      *Context
	winW     int
	winH     int
	prevW    int
	prevH    int
	dirty    bool
}

func NewContext(bgColor color.NRGBA) *Context {
	return &Context{bgColor: bgColor}
}

func (c *Context) SetMinWindowSize(minW, minH int) {
	c.minW = minW
	c.minH = minH
}

func (c *Context) initFont() {
	tt, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatalf("raitui: font parse: %v", err)
	}
	c.font, _ = opentype.NewFace(tt, &opentype.FaceOptions{Size: 14, DPI: 72})
	c.fontSm, _ = opentype.NewFace(tt, &opentype.FaceOptions{Size: 11, DPI: 72})
	c.fontLg, _ = opentype.NewFace(tt, &opentype.FaceOptions{Size: 20, DPI: 72})
}

func (c *Context) Font() font.Face   { return c.font }
func (c *Context) FontSm() font.Face { return c.fontSm }
func (c *Context) FontLg() font.Face { return c.fontLg }

func (c *Context) Run(root *Element, title string, width, height int) {
	c.Root = root
	c.winW = width
	c.winH = height
	c.kb = keyboard.New()
	root.setLayoutCtx(c)
	c.initFont()

	c.Game = &game{
		ctx:   c,
		winW:  width,
		winH:  height,
		dirty: true,
	}

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if c.minW > 0 && c.minH > 0 {
		ebiten.SetWindowSizeLimits(c.minW, c.minH, -1, -1)
	}

	if err := ebiten.RunGame(c.Game); err != nil {
		log.Fatal(err)
	}
}

func (c *Context) RebuildLayout() {
	if c.Root != nil {
		c.ForceLayout()
	}
}

func (c *Context) ForceLayout() {
	if c.Root == nil {
		return
	}
	goda.CalculateNodeLayout(c.Root.GNode, float32(c.winW), float32(c.winH), goda.DirectionLTR)
}

func selectFont(ctx *Context, elem *Element) font.Face {
	fs := elem.FontSizeValue()
	if fs <= 0 {
		return ctx.font
	}
	if fs <= 11 {
		return ctx.fontSm
	}
	if fs >= 18 {
		return ctx.fontLg
	}
	return ctx.font
}

func (g *game) Update() error {
	g.ctx.frame++

	mx, my := ebiten.CursorPosition()
	g.ctx.mouseX = mx
	g.ctx.mouseY = my

	if g.dirty {
		g.ctx.RebuildLayout()
		g.dirty = false
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && g.ctx.Root != nil {
		dispatchClick(g.ctx, g.ctx.Root, 0, 0, float32(mx), float32(my))
	}

	if g.ctx.Root != nil {
		dispatchHover(g.ctx.Root, 0, 0, float32(mx), float32(my), g.ctx.modal)
	}

	if g.ctx.focused != nil && g.ctx.focused.ElemType == TypeInput {
		handleKeyboardInput(g.ctx.focused, g.ctx.kb, g.ctx.frame)
	}

	g.ctx.kb.CheckShortcuts(g.ctx.focused != nil)

	return nil
}

func dispatchHover(elem *Element, absLeft, absTop, mx, my float32, modal *Element) {
	if !elem.visible {
		return
	}

	lo := elem.GNode.LayoutOut()
	x := absLeft + float32(lo.Left)
	y := absTop + float32(lo.Top)
	w := float32(lo.Width)
	h := float32(lo.Height)

	hovered := w > 0 && h > 0 && mx >= x && mx < x+w && my >= y && my < y+h

	if elem.onHover != nil {
		blocked := modal != nil && !inModalTree(elem, modal)
		if !blocked {
			elem.onHover(hovered)
		}
	}

	for _, child := range elem.children {
		dispatchHover(child, x, y, mx, my, modal)
	}
}

func inModalTree(elem, modal *Element) bool {
	if elem == modal {
		return true
	}
	for p := elem.parent; p != nil; p = p.parent {
		if p == modal {
			return true
		}
	}
	for p := modal.parent; p != nil; p = p.parent {
		if p == elem {
			return true
		}
	}
	return false
}

func (e *Element) IsChildOf(ancestor *Element) bool {
	if e == ancestor {
		return true
	}
	for p := e.parent; p != nil; p = p.parent {
		if p == ancestor {
			return true
		}
	}
	return false
}

func dispatchClick(ctx *Context, elem *Element, absLeft, absTop, mx, my float32) bool {
	if !elem.visible {
		return false
	}

	if ctx.modal != nil && !inModalTree(elem, ctx.modal) {
		return false
	}

	lo := elem.GNode.LayoutOut()
	x := absLeft + float32(lo.Left)
	y := absTop + float32(lo.Top)
	w := float32(lo.Width)
	h := float32(lo.Height)

	hit := w > 0 && h > 0 && mx >= x && mx < x+w && my >= y && my < y+h

	if hit {
		oldFocus := ctx.focused
		if elem.ElemType == TypeInput {
			ctx.focused = elem
		} else {
			ctx.focused = nil
		}
		if oldFocus != ctx.focused {
			ctx.kb.Clear()
		}
		if elem.onClick != nil {
			elem.onClick()
			return true
		}
		if elem.ElemType == TypeInput || elem.ElemType == TypeCheckbox || elem.ElemType == TypeSwitch {
			return true
		}
	}

	for _, child := range elem.children {
		if dispatchClick(ctx, child, x, y, mx, my) {
			return true
		}
	}
	return false
}
func handleKeyboardInput(elem *Element, kb *keyboard.Handler, frame int) {
	insert := func(ch rune) {
		s, pos := keyboard.InsertRune(elem.TextContent(), elem._cursorPos, ch)
		elem.SetTextContent(s)
		if pos < 0 {
			pos = 0
		}
		rs := []rune(s)
		if pos > len(rs) {
			pos = len(rs)
		}
		elem._cursorPos = pos
	}

	chars := kb.InputChars()
	for _, r := range chars {
		if r >= ' ' || r == '\n' || r == '\t' {
			insert(r)
		}
	}

	if kb.KeyHeld(ebiten.KeyLeft, 15, 2) {
		if elem._cursorPos > 0 {
			elem._cursorPos--
		}
	}
	if kb.KeyHeld(ebiten.KeyRight, 15, 2) {
		rs := []rune(elem.TextContent())
		if elem._cursorPos < len(rs) {
			elem._cursorPos++
		}
	}
	if kb.KeyPressed(ebiten.KeyHome) || kb.KeyPressed(ebiten.KeyUp) {
		elem._cursorPos = 0
	}
	if kb.KeyPressed(ebiten.KeyEnd) || kb.KeyPressed(ebiten.KeyDown) {
		rs := []rune(elem.TextContent())
		elem._cursorPos = len(rs)
	}

	delOne := false
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		delOne = true
		elem._bsStart = frame
	} else if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		held := frame - elem._bsStart
		if held > 1 && held > 15 && (held-15)%2 == 0 {
			delOne = true
		}
	}
	if delOne {
		rs := []rune(elem.TextContent())
		pos := elem._cursorPos
		if pos > len(rs) {
			pos = len(rs)
		}
		if pos > 0 && len(rs) > 0 {
			rs = append(rs[:pos-1], rs[pos:]...)
			elem.SetTextContent(string(rs))
			pos--
			if pos < 0 {
				pos = 0
			}
			elem._cursorPos = pos
		}
	}

	if kb.KeyPressed(ebiten.KeyDelete) {
		s, pos := keyboard.DeleteForward(elem.TextContent(), elem._cursorPos)
		elem.SetTextContent(s)
		rs := []rune(s)
		if pos > len(rs) {
			pos = len(rs)
		}
		elem._cursorPos = pos
	}

	if kb.KeyPressed(ebiten.KeyEnter) || kb.KeyPressed(ebiten.KeyKPEnter) {
		insert('\n')
	}
	if kb.KeyPressed(ebiten.KeyTab) {
		insert('\t')
	}
}

func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(g.ctx.bgColor)
	if g.ctx.Root != nil {
		renderElement(screen, g.ctx.Root, 0, 0, g.ctx)
	}

	var floating []*Element
	collectFloating(g.ctx.Root, &floating)
	for _, el := range floating {
		if el._triggerRef != nil {
			renderWithTrigger(screen, el, g.ctx)
		}
	}

	if g.ctx.Debug {
		b := ColorDebugRed
		vector.StrokeRect(screen, 0, 0, float32(g.winW)-1, float32(g.winH)-1, 2, b, true)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	if outsideWidth <= 0 {
		outsideWidth = 800
	}
	if outsideHeight <= 0 {
		outsideHeight = 600
	}
	if outsideWidth != g.prevW || outsideHeight != g.prevH {
		g.prevW = outsideWidth
		g.prevH = outsideHeight
		g.winW = outsideWidth
		g.winH = outsideHeight
		g.ctx.winW = outsideWidth
		g.ctx.winH = outsideHeight
		g.dirty = true
	}
	return outsideWidth, outsideHeight
}

func collectFloating(elem *Element, out *[]*Element) {
	if elem._triggerRef != nil && elem.visible {
		*out = append(*out, elem)
	}
	for _, child := range elem.children {
		collectFloating(child, out)
	}
}

func renderWithTrigger(screen *ebiten.Image, elem *Element, ctx *Context) {
	ref := elem._triggerRef
	if ref == nil {
		return
	}
	refLo := ref.GNode.LayoutOut()
	refAbsLeft := refLo.Left
	refAbsTop := refLo.Top
	cur := ref.parent
	for cur != nil {
		curLo := cur.GNode.LayoutOut()
		refAbsLeft += curLo.Left
		refAbsTop += curLo.Top
		cur = cur.parent
	}

	toolW := float32(elem.GNode.LayoutOut().Width)
	toolH := float32(elem.GNode.LayoutOut().Height)
	refW := float32(refLo.Width)
	refH := float32(refLo.Height)
	gap := float32(6)

	var atX, atY float32
	switch Placement(elem._tooltipPlace) {
	case PlaceTop:
		atY = refAbsTop - toolH - gap
		atX = refAbsLeft
	case PlaceBottom:
		atY = refAbsTop + refH + gap
		atX = refAbsLeft
	case PlaceLeft:
		atX = refAbsLeft - toolW - gap
		atY = refAbsTop
	case PlaceRight:
		atX = refAbsLeft + refW + gap
		atY = refAbsTop
	}

	renderElementAt(screen, elem, atX, atY, ctx)
}

func renderElementAt(screen *ebiten.Image, elem *Element, absX, absY float32, ctx *Context) {
	lo := elem.GNode.LayoutOut()
	x := absX
	y := absY
	w := float32(lo.Width)
	h := float32(lo.Height)

	mx := float32(ctx.mouseX)
	my := float32(ctx.mouseY)
	hovered := w > 0 && h > 0 && mx >= x && mx < x+w && my >= y && my < y+h

	rctx := &RenderCtx{
		Screen:  screen,
		Elem:    elem,
		X:       x,
		Y:       y,
		W:       w,
		H:       h,
		Font:    selectFont(ctx, elem),
		Hovered: hovered,
		Focused: ctx.focused == elem,
		Frame:   ctx.frame,
	}

	switch elem.ElemType {
	case TypeBox:
		renderBox(rctx, elem)
	case TypeText:
		renderText(rctx, elem)
	case TypeVStack, TypeHStack, TypeFlex, TypeCenter:
		renderStack(rctx, elem)
	case TypeSeparator:
		renderSeparator(rctx, elem)
	case TypeCheckbox:
		renderCheckbox(rctx, elem)
	case TypeSwitch:
		renderSwitch(rctx, elem)
	case TypeInput:
		renderInput(rctx, elem)
	case TypeButton:
		renderButton(rctx, elem)
	case TypeSpinner:
		renderSpinner(rctx, elem)
	case TypeProgress:
		renderProgress(rctx, elem)
	case TypeAvatar:
		renderAvatar(rctx, elem)
	case TypeMenu, TypeDialog, TypeTooltip:
		renderStack(rctx, elem)
	}

	if elem.ElemType != TypeAvatar {
		for _, child := range elem.children {
			childLo := child.GNode.LayoutOut()
			cx := x + float32(childLo.Left)
			cy := y + float32(childLo.Top)
			renderElementAt(screen, child, cx, cy, ctx)
		}
	}
}

func renderElement(screen *ebiten.Image, elem *Element, absLeft, absTop float32, ctx *Context) {
	if !elem.visible {
		return
	}

	if elem._triggerRef != nil {
		return
	}

	lo := elem.GNode.LayoutOut()
	x := absLeft + float32(lo.Left)
	y := absTop + float32(lo.Top)
	w := float32(lo.Width)
	h := float32(lo.Height)

	mx := float32(ctx.mouseX)
	my := float32(ctx.mouseY)
	hovered := w > 0 && h > 0 && mx >= x && mx < x+w && my >= y && my < y+h

	rctx := &RenderCtx{
		Screen:  screen,
		Elem:    elem,
		X:       x,
		Y:       y,
		W:       w,
		H:       h,
		Font:    selectFont(ctx, elem),
		Hovered: hovered,
		Focused: ctx.focused == elem,
		Frame:   ctx.frame,
	}

	switch elem.ElemType {
	case TypeBox:
		renderBox(rctx, elem)
	case TypeText:
		renderText(rctx, elem)
	case TypeVStack, TypeHStack, TypeFlex, TypeCenter:
		renderStack(rctx, elem)
	case TypeSeparator:
		renderSeparator(rctx, elem)
	case TypeCheckbox:
		renderCheckbox(rctx, elem)
	case TypeSwitch:
		renderSwitch(rctx, elem)
	case TypeInput:
		renderInput(rctx, elem)
	case TypeButton:
		renderButton(rctx, elem)
	case TypeSpinner:
		renderSpinner(rctx, elem)
	case TypeProgress:
		renderProgress(rctx, elem)
	case TypeAvatar:
		renderAvatar(rctx, elem)
	case TypeMenu, TypeDialog, TypeTooltip:
		renderStack(rctx, elem)
	}

	if ctx.Debug && w > 0 && h > 0 {
		debugColors := []color.NRGBA{
			{255, 0, 0, 255}, {0, 200, 0, 255}, {0, 0, 255, 255},
			{255, 200, 0, 255}, {255, 0, 255, 255}, {0, 200, 200, 255},
			{255, 100, 0, 255}, {150, 0, 255, 255},
		}
		depth := elemDebugDepth(elem)
		clr := debugColors[depth%len(debugColors)]
		vector.StrokeRect(screen, x, y, w, h, 1, clr, true)

		var label string
		switch elem.ElemType {
		case TypeBox:
			label = "B"
		case TypeText:
			label = "T"
		case TypeVStack:
			label = "V"
		case TypeHStack:
			label = "H"
		case TypeFlex:
			label = "F"
		case TypeCenter:
			label = "C"
		case TypeSeparator:
			label = "S"
		case TypeCheckbox:
			label = "Ck"
		case TypeSwitch:
			label = "Sw"
		case TypeInput:
			label = "In"
		case TypeButton:
			label = "Bt"
		case TypeSpinner:
			label = "Sp"
		case TypeProgress:
			label = "Pr"
		case TypeAvatar:
			label = "Av"
		case TypeDialog:
			label = "Dg"
		case TypeMenu:
			label = "Mn"
		case TypeTooltip:
			label = "Tt"
		}
		txt := fmt.Sprintf("%s %.0fx%.0f", label, w, h)
		text.Draw(screen, txt, ctx.fontSm, int(x+2), int(y+10), clr)
	}

	if elem.ElemType != TypeAvatar {
		var absoluteChildren []*Element
		for _, child := range elem.children {
			if child.GNode.GetPositionType() == goda.PositionTypeAbsolute {
				absoluteChildren = append(absoluteChildren, child)
			} else {
				renderElement(screen, child, x, y, ctx)
			}
		}
		for _, child := range absoluteChildren {
			renderElement(screen, child, x, y, ctx)
		}
	}
}

func elemDebugDepth(elem *Element) int {
	d := 0
	for p := elem.parent; p != nil; p = p.parent {
		d++
	}
	return d
}
