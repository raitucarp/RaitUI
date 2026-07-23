package core

import (
	"image/color"
	"strconv"
	"strings"

	goda "goda"
	"github.com/hajimehoshi/ebiten/v2"
)

type Element struct {
	GNode    *goda.Node
	children []*Element
	parent   *Element

	ElemType ElementType

	textContent string

	bgColor   color.NRGBA
	borderClr color.NRGBA
	textClr   color.NRGBA
	radius    float32
	fontSize  float32
	opacity   float32

	visible bool
	hasInit bool

	tooltip  string
	lang     string
	dir      string
	tabIndex int
	zIndex   int
	elemID   string

	shadow      Shadow
	textAlign   TextAlign
	lineHeight  float32
	cursor      CursorType
	borderStyle BorderStyle

	radiusTL float32
	radiusTR float32
	radiusBL float32
	radiusBR float32

	onClick   func()
	onHover   func(entered bool)
	underline bool
	checked     bool
	placeholder string
	textBind    func() string
	_bsStart    int
	_cursorPos  int
	_progress   float32
	avatarImg   any
	hsh         string
	DialogCard  *Element
	DialogFooter *Element
	_layoutCtx  *Context
	_tooltipPlace int
	_triggerRef  *Element
	_scrollY     float32
	_scrollMax   float32
	drawFunc     func(ctx CanvasContext)
}

type CanvasContext struct {
	Screen *ebiten.Image
	X      float32
	Y      float32
	W      float32
	H      float32
	Frame  int
}

func NewElement(typ ElementType) *Element {
	gNode := goda.New()
	gNode.SetBoxSizing(goda.BoxSizingBorderBox)
	return &Element{
		GNode:    gNode,
		ElemType: typ,
		visible:  true,
		opacity:  1.0,
		hasInit:  false,
	}
}

func (e *Element) Init() {
	if e.hasInit {
		return
	}
	e.hasInit = true
}

func (e *Element) Children(elems ...*Element) *Element {
	e.children = elems
	for i, child := range elems {
		child.parent = e
		e.GNode.InsertChildNode(child.GNode, i)
	}
	return e
}

func (e *Element) AppendChild(child *Element) *Element {
	idx := e.GNode.GetChildCount()
	child.parent = e
	e.children = append(e.children, child)
	e.GNode.InsertChildNode(child.GNode, idx)
	return e
}

func (e *Element) ChildrenCount() int {
	return len(e.children)
}

func (e *Element) ChildAt(i int) *Element {
	if i < 0 || i >= len(e.children) {
		return nil
	}
	return e.children[i]
}

func (e *Element) Width(v string) *Element {
	e.GNode.SetWidthAuto()
	val, unit := parseValue(v)
	switch unit {
	case "px", "":
		e.GNode.SetWidth(val)
	case "%":
		e.GNode.SetWidthPercent(val)
	case "auto":
		e.GNode.SetWidthAuto()
	}
	return e
}

func (e *Element) Height(v string) *Element {
	e.GNode.SetHeightAuto()
	val, unit := parseValue(v)
	switch unit {
	case "px", "":
		e.GNode.SetHeight(val)
	case "%":
		e.GNode.SetHeightPercent(val)
	case "auto":
		e.GNode.SetHeightAuto()
	}
	return e
}

func (e *Element) MinWidth(v string) *Element {
	val, unit := parseValue(v)
	switch unit {
	case "px", "":
		e.GNode.SetMinWidth(val)
	case "%":
		e.GNode.SetMinWidthPercent(val)
	}
	return e
}

func (e *Element) MinHeight(v string) *Element {
	val, unit := parseValue(v)
	switch unit {
	case "px", "":
		e.GNode.SetMinHeight(val)
	case "%":
		e.GNode.SetMinHeightPercent(val)
	}
	return e
}

func (e *Element) MaxWidth(v string) *Element {
	val, unit := parseValue(v)
	switch unit {
	case "px", "":
		e.GNode.SetMaxWidth(val)
	case "%":
		e.GNode.SetMaxWidthPercent(val)
	}
	return e
}

func (e *Element) MaxHeight(v string) *Element {
	val, unit := parseValue(v)
	switch unit {
	case "px", "":
		e.GNode.SetMaxHeight(val)
	case "%":
		e.GNode.SetMaxHeightPercent(val)
	}
	return e
}

func (e *Element) Padding(v string) *Element {
	val := parseFloat(v)
	e.GNode.SetPadding(goda.EdgeAll, val)
	return e
}

func (e *Element) PaddingTop(v string) *Element {
	e.GNode.SetPadding(goda.EdgeTop, parseFloat(v))
	return e
}

func (e *Element) PaddingRight(v string) *Element {
	e.GNode.SetPadding(goda.EdgeRight, parseFloat(v))
	return e
}

func (e *Element) PaddingBottom(v string) *Element {
	e.GNode.SetPadding(goda.EdgeBottom, parseFloat(v))
	return e
}

func (e *Element) PaddingLeft(v string) *Element {
	e.GNode.SetPadding(goda.EdgeLeft, parseFloat(v))
	return e
}

func (e *Element) PaddingX(v string) *Element {
	val := parseFloat(v)
	e.GNode.SetPadding(goda.EdgeLeft, val)
	e.GNode.SetPadding(goda.EdgeRight, val)
	return e
}

func (e *Element) PaddingY(v string) *Element {
	val := parseFloat(v)
	e.GNode.SetPadding(goda.EdgeTop, val)
	e.GNode.SetPadding(goda.EdgeBottom, val)
	return e
}

func (e *Element) Margin(v string) *Element {
	val := parseFloat(v)
	e.GNode.SetMargin(goda.EdgeAll, val)
	return e
}

func (e *Element) MarginTop(v string) *Element {
	e.GNode.SetMargin(goda.EdgeTop, parseFloat(v))
	return e
}

func (e *Element) MarginRight(v string) *Element {
	e.GNode.SetMargin(goda.EdgeRight, parseFloat(v))
	return e
}

func (e *Element) MarginBottom(v string) *Element {
	e.GNode.SetMargin(goda.EdgeBottom, parseFloat(v))
	return e
}

func (e *Element) MarginLeft(v string) *Element {
	e.GNode.SetMargin(goda.EdgeLeft, parseFloat(v))
	return e
}

func (e *Element) MarginX(v string) *Element {
	val := parseFloat(v)
	e.GNode.SetMargin(goda.EdgeLeft, val)
	e.GNode.SetMargin(goda.EdgeRight, val)
	return e
}

func (e *Element) MarginY(v string) *Element {
	val := parseFloat(v)
	e.GNode.SetMargin(goda.EdgeTop, val)
	e.GNode.SetMargin(goda.EdgeBottom, val)
	return e
}

func (e *Element) FlexDirection(d goda.FlexDirection) *Element {
	e.GNode.SetFlexDirection(d)
	return e
}

func (e *Element) JustifyContent(j goda.Justify) *Element {
	e.GNode.SetJustifyContent(j)
	return e
}

func (e *Element) AlignItems(a goda.Align) *Element {
	e.GNode.SetAlignItems(a)
	return e
}

func (e *Element) AlignContent(a goda.Align) *Element {
	e.GNode.SetAlignContent(a)
	return e
}

func (e *Element) AlignSelf(a goda.Align) *Element {
	e.GNode.SetAlignSelf(a)
	return e
}

func (e *Element) FlexWrap(w goda.Wrap) *Element {
	e.GNode.SetFlexWrap(w)
	return e
}

func (e *Element) FlexGrow(v float32) *Element {
	e.GNode.SetFlexGrow(v)
	return e
}

func (e *Element) FlexShrink(v float32) *Element {
	e.GNode.SetFlexShrink(v)
	return e
}

func (e *Element) FlexBasis(v string) *Element {
	val, unit := parseValue(v)
	switch unit {
	case "px", "":
		e.GNode.SetFlexBasis(val)
	case "%":
		e.GNode.SetFlexBasisPercent(val)
	case "auto":
		e.GNode.SetFlexBasisAuto()
	}
	return e
}

func (e *Element) Gap(v string) *Element {
	e.GNode.SetGap(goda.GutterAll, parseFloat(v))
	return e
}

func (e *Element) ColumnGap(v string) *Element {
	e.GNode.SetGap(goda.GutterColumn, parseFloat(v))
	return e
}

func (e *Element) RowGap(v string) *Element {
	e.GNode.SetGap(goda.GutterRow, parseFloat(v))
	return e
}

func (e *Element) BorderWidth(v string) *Element {
	e.GNode.SetBorder(goda.EdgeAll, parseFloat(v))
	return e
}

func (e *Element) AspectRatio(v float32) *Element {
	e.GNode.SetAspectRatio(v)
	return e
}

func (e *Element) Position(v goda.PositionType) *Element {
	e.GNode.SetPositionType(v)
	return e
}

func (e *Element) Display(v goda.Display) *Element {
	e.GNode.SetDisplay(v)
	return e
}

func (e *Element) Overflow(v goda.Overflow) *Element {
	e.GNode.SetOverflow(v)
	return e
}

func (e *Element) BoxSizing(v goda.BoxSizing) *Element {
	e.GNode.SetBoxSizing(v)
	return e
}

func (e *Element) BackgroundColor(c color.NRGBA) *Element {
	e.bgColor = c
	return e
}

func (e *Element) BorderColor(c color.NRGBA) *Element {
	e.borderClr = c
	return e
}

func (e *Element) TextColor(c color.NRGBA) *Element {
	e.textClr = c
	return e
}

func (e *Element) BorderRadius(v float32) *Element {
	e.radius = v
	return e
}

func (e *Element) FontSize(v float32) *Element {
	e.fontSize = v
	return e
}

func (e *Element) Opacity(v float32) *Element {
	e.opacity = v
	return e
}

func (e *Element) Visible(v bool) *Element {
	e.visible = v
	if v {
		e.GNode.SetDisplay(goda.DisplayFlex)
	} else {
		e.GNode.SetDisplay(goda.DisplayNone)
	}
	if e._layoutCtx != nil {
		e._layoutCtx.ForceLayout()
	}
	return e
}

func (e *Element) ToggleVisible() *Element {
	return e.Visible(!e.visible)
}

func (e *Element) Open() *Element {
	e.Visible(true)
	if e._layoutCtx != nil {
		e._layoutCtx.modal = e
	}
	return e
}

func (e *Element) Close() *Element {
	e.Visible(false)
	if e._layoutCtx != nil {
		e._layoutCtx.modal = nil
	}
	return e
}

func (e *Element) IsOpen() bool { return e.visible }

func (e *Element) TooltipPlacement(p Placement) *Element {
	e._tooltipPlace = int(p)
	e.GNode.SetEdgePositionAuto(goda.EdgeTop)
	e.GNode.SetEdgePositionAuto(goda.EdgeLeft)
	e.GNode.SetEdgePositionAuto(goda.EdgeRight)
	e.GNode.SetEdgePositionAuto(goda.EdgeBottom)

	switch p {
	case PlaceTop:
		e.GNode.SetEdgePosition(goda.EdgeTop, float32(-36))
	case PlaceBottom:
		e.GNode.SetEdgePosition(goda.EdgeTop, float32(44))
	case PlaceLeft:
		e.GNode.SetEdgePosition(goda.EdgeRight, float32(128))
	case PlaceRight:
		e.GNode.SetEdgePosition(goda.EdgeLeft, float32(128))
	}
	return e
}

func (e *Element) hasVisibleOverlay() bool {
	for _, child := range e.children {
		if child.visible && child.ElemType == TypeBox && child.Width("") != nil && child.Height("") != nil {
			// Check if child is a full-screen overlay (like dialog backdrop)
			lo := child.GNode.LayoutOut()
			if lo.Width > 100 && lo.Height > 100 {
				return true
			}
		}
	}
	return false
}

func (e *Element) setLayoutCtx(ctx *Context) {
	e._layoutCtx = ctx
	for _, child := range e.children {
		child.setLayoutCtx(ctx)
	}
}

func (e *Element) TextContent() string {
	return e.textContent
}

func (e *Element) SetTextContent(s string) {
	e.textContent = s
}

func (e *Element) BackgroundColorValue() color.NRGBA {
	return e.bgColor
}

func (e *Element) BorderColorValue() color.NRGBA {
	return e.borderClr
}

func (e *Element) TextColorValue() color.NRGBA {
	return e.textClr
}

func (e *Element) BorderRadiusValue() float32 {
	return e.radius
}

func (e *Element) FontSizeValue() float32 {
	return e.fontSize
}

func (e *Element) OpacityValue() float32 {
	return e.opacity
}

func (e *Element) IsVisible() bool {
	return e.visible
}

func (e *Element) Walk(fn func(elem *Element)) {
	fn(e)
	for _, child := range e.children {
		child.Walk(fn)
	}
}

// ============================================
// Additional CSS Props
// ============================================

func (e *Element) BorderTop(v string) *Element {
	e.GNode.SetBorder(goda.EdgeTop, parseFloat(v))
	return e
}

func (e *Element) BorderRight(v string) *Element {
	e.GNode.SetBorder(goda.EdgeRight, parseFloat(v))
	return e
}

func (e *Element) BorderBottom(v string) *Element {
	e.GNode.SetBorder(goda.EdgeBottom, parseFloat(v))
	return e
}

func (e *Element) BorderLeft(v string) *Element {
	e.GNode.SetBorder(goda.EdgeLeft, parseFloat(v))
	return e
}

func (e *Element) BorderStyle(style BorderStyle) *Element {
	e.borderStyle = style
	return e
}

func (e *Element) BorderTopLeftRadius(v float32) *Element {
	e.radiusTL = v
	return e
}

func (e *Element) BorderTopRightRadius(v float32) *Element {
	e.radiusTR = v
	return e
}

func (e *Element) BorderBottomLeftRadius(v float32) *Element {
	e.radiusBL = v
	return e
}

func (e *Element) BorderBottomRightRadius(v float32) *Element {
	e.radiusBR = v
	return e
}

func (e *Element) PositionTop(v string) *Element {
	val, unit := parseValue(v)
	switch unit {
	case "%":
		e.GNode.SetEdgePositionPercent(goda.EdgeTop, val)
	default:
		e.GNode.SetEdgePosition(goda.EdgeTop, val)
	}
	return e
}

func (e *Element) PositionRight(v string) *Element {
	val, unit := parseValue(v)
	switch unit {
	case "%":
		e.GNode.SetEdgePositionPercent(goda.EdgeRight, val)
	default:
		e.GNode.SetEdgePosition(goda.EdgeRight, val)
	}
	return e
}

func (e *Element) PositionBottom(v string) *Element {
	val, unit := parseValue(v)
	switch unit {
	case "%":
		e.GNode.SetEdgePositionPercent(goda.EdgeBottom, val)
	default:
		e.GNode.SetEdgePosition(goda.EdgeBottom, val)
	}
	return e
}

func (e *Element) PositionLeft(v string) *Element {
	val, unit := parseValue(v)
	switch unit {
	case "%":
		e.GNode.SetEdgePositionPercent(goda.EdgeLeft, val)
	default:
		e.GNode.SetEdgePosition(goda.EdgeLeft, val)
	}
	return e
}

func (e *Element) ZIndex(v int) *Element {
	e.zIndex = v
	return e
}

func (e *Element) BoxShadow(offsetX, offsetY, blur, spread float32, c color.NRGBA) *Element {
	e.shadow = Shadow{
		OffsetX: offsetX,
		OffsetY: offsetY,
		Blur:    blur,
		Spread:  spread,
		Color:   c,
	}
	return e
}

func (e *Element) TextAlign(align TextAlign) *Element {
	e.textAlign = align
	return e
}

func (e *Element) LineHeight(v float32) *Element {
	e.lineHeight = v
	return e
}

func (e *Element) Cursor(cur CursorType) *Element {
	e.cursor = cur
	return e
}

func (e *Element) Scrollable(v bool) *Element {
	if v {
		e.GNode.SetOverflow(goda.OverflowScroll)
	} else {
		e.GNode.SetOverflow(goda.OverflowVisible)
	}
	return e
}

func (e *Element) Use(getter func() string) *Element {
	e.textBind = getter
	return e
}

func (e *Element) TextBind() func() string { return e.textBind }

func (e *Element) ProgressValue() float32 { return e._progress }
func (e *Element) SetProgressValue(v float32) { e._progress = v }

func (e *Element) SetDrawFunc(fn func(CanvasContext)) *Element {
	e.drawFunc = fn
	return e
}

func (e *Element) DrawFunc() func(CanvasContext) { return e.drawFunc }

func (e *Element) SetTrigger(ref *Element) *Element {
	e._triggerRef = ref
	return e
}

func (e *Element) TriggerRef() *Element { return e._triggerRef }

func (e *Element) ScrollY() float32     { return e._scrollY }
func (e *Element) SetScrollY(v float32) { e._scrollY = v; if e._scrollY < 0 { e._scrollY = 0 }; if e._scrollMax > 0 && e._scrollY > e._scrollMax { e._scrollY = e._scrollMax } }
func (e *Element) ScrollMax() float32   { return e._scrollMax }
func (e *Element) SetScrollMax(v float32) { e._scrollMax = v }
func (e *Element) Scroll(dy float32)    { e.SetScrollY(e._scrollY + dy) }

func (e *Element) Image(img any) *Element {
	e.avatarImg = img
	return e
}

func (e *Element) AvatarImage() any { return e.avatarImg }

// ============================================
// Getters for new fields
// ============================================

func (e *Element) Tooltip() string         { return e.tooltip }
func (e *Element) LangValue() string       { return e.lang }
func (e *Element) DirValue() string        { return e.dir }
func (e *Element) TabIndexValue() int      { return e.tabIndex }
func (e *Element) ZIndexValue() int        { return e.zIndex }
func (e *Element) ShadowValue() Shadow     { return e.shadow }
func (e *Element) TextAlignValue() TextAlign { return e.textAlign }
func (e *Element) LineHeightValue() float32  { return e.lineHeight }
func (e *Element) CursorValue() CursorType   { return e.cursor }
func (e *Element) BorderStyleValue() BorderStyle { return e.borderStyle }
func (e *Element) CornerRadius() (tl, tr, bl, br float32) {
	return e.radiusTL, e.radiusTR, e.radiusBL, e.radiusBR
}

func parseFloat(v string) float32 {
	v = strings.TrimSpace(v)
	v = strings.TrimSuffix(v, "px")
	v = strings.TrimSuffix(v, "PX")
	val, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return 0
	}
	return float32(val)
}

func parseValue(v string) (float32, string) {
	v = strings.TrimSpace(strings.ToLower(v))
	if v == "auto" {
		return 0, "auto"
	}
	if strings.HasSuffix(v, "%") {
		f := parseFloat(strings.TrimSuffix(v, "%"))
		return f, "%"
	}
	if strings.HasSuffix(v, "px") {
		return parseFloat(v), "px"
	}
	return parseFloat(v), ""
}
