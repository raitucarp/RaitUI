package core

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
)

func renderBox(rctx *RenderCtx, elem *Element) {
	if rctx.W <= 0 || rctx.H <= 0 {
		return
	}
	bg := elem.bgColor
	border := elem.borderClr
	alpha := uint8(255 * elem.opacity)
	if alpha == 0 {
		return
	}

	hasBg := bg.A > 0
	hasBorder := border.A > 0 && border != bg

	if hasBg {
		bg.A = alpha
	}
	if hasBorder {
		border.A = alpha
	} else {
		border = color.NRGBA{}
	}

	x, y, w, h := rctx.X, rctx.Y, rctx.W, rctx.H
	r := elem.radius
	tl, tr, bl, br := elem.CornerRadius()
	if tl > 0 || tr > 0 || bl > 0 || br > 0 {
		r = max(max(tl, tr), max(bl, br))
	}

	if !hasBg && !hasBorder {
		return
	}

	if r > 0 {
		drawRoundedBox(screen(rctx), x, y, w, h, r, bg, border, elem.shadow)
	} else {
		if elem.shadow.Color.A > 0 {
			drawShadow(screen(rctx), x, y, w, h, elem.shadow)
		}
		if hasBg {
			drawFilledRect(screen(rctx), x, y, w, h, bg)
		}
		if hasBorder {
			vector.StrokeRect(screen(rctx), x, y, w, h, 1, border, true)
		}
	}
}

func drawRoundedBox(screen *ebiten.Image, x, y, w, h, r float32, fill, stroke color.NRGBA, s Shadow) {
	if s.Color.A > 0 {
		sx := x + s.OffsetX
		sy := y + s.OffsetY
		sw := w + s.Spread*2
		sh := h + s.Spread*2
		sr := r + s.Spread
		if sr < 0 {
			sr = 0
		}
		filledRoundedRect(screen, sx, sy, sw, sh, sr, s.Color)
	}

	hasFill := fill.A > 0
	hasStroke := stroke.A > 0 && stroke != fill

	if hasFill && !hasStroke {
		filledRoundedRect(screen, x, y, w, h, r, fill)
		return
	}

	if hasStroke && !hasFill {
		drawRoundedBorder(screen, x, y, w, h, r, 1.5, stroke)
		return
	}

	if hasFill && hasStroke {
		filledRoundedRect(screen, x, y, w, h, r, stroke)
		r2 := r - 1.5
		if r2 < 0 {
			r2 = 0
		}
		filledRoundedRect(screen, x+1.5, y+1.5, w-3, h-3, r2, fill)
	}
}

func drawRoundedBorder(screen *ebiten.Image, x, y, w, h, r, thickness float32, clr color.NRGBA) {
	if r <= 0 || r*2 > w || r*2 > h {
		vector.StrokeRect(screen, x, y, w, h, thickness, clr, true)
		return
	}
	segments := 8
	for i := 0; i < segments; i++ {
		a0 := float64(i) * (math.Pi / 2) / float64(segments)
		a1 := float64(i+1) * (math.Pi / 2) / float64(segments)

		for _, corner := range []struct{ cx, cy, sx, sy float32 }{
			{x + r, y + r, -1, -1},
			{x + w - r, y + r, 1, -1},
			{x + r, y + h - r, -1, 1},
			{x + w - r, y + h - r, 1, 1},
		} {
			cx := corner.cx
			cy := corner.cy
			x0 := cx + float32(math.Cos(a0))*float32(corner.sx)*r
			y0 := cy + float32(math.Sin(a0))*float32(corner.sy)*r
			x1 := cx + float32(math.Cos(a1))*float32(corner.sx)*r
			y1 := cy + float32(math.Sin(a1))*float32(corner.sy)*r
			vector.StrokeLine(screen, x0, y0, x1, y1, thickness, clr, true)
		}
	}

	vector.StrokeLine(screen, x+r, y, x+w-r, y, thickness, clr, true)
	vector.StrokeLine(screen, x+r, y+h, x+w-r, y+h, thickness, clr, true)
	vector.StrokeLine(screen, x, y+r, x, y+h-r, thickness, clr, true)
	vector.StrokeLine(screen, x+w, y+r, x+w, y+h-r, thickness, clr, true)
}

func filledRoundedRect(screen *ebiten.Image, x, y, w, h, r float32, fill color.NRGBA) {
	if w <= 0 || h <= 0 || fill.A == 0 {
		return
	}
	if r <= 0 || r*2 > w || r*2 > h {
		drawFilledRect(screen, x, y, w, h, fill)
		return
	}

	vector.DrawFilledRect(screen, x+r, y, w-2*r, h, fill, true)
	vector.DrawFilledRect(screen, x, y+r, r, h-2*r, fill, true)
	vector.DrawFilledRect(screen, x+w-r, y+r, r, h-2*r, fill, true)

	vector.DrawFilledCircle(screen, x+r, y+r, r, fill, true)
	vector.DrawFilledCircle(screen, x+w-r, y+r, r, fill, true)
	vector.DrawFilledCircle(screen, x+r, y+h-r, r, fill, true)
	vector.DrawFilledCircle(screen, x+w-r, y+h-r, r, fill, true)
}

func drawShadow(screen *ebiten.Image, x, y, w, h float32, s Shadow) {
	steps := 3
	for i := 0; i < steps; i++ {
		offX := s.OffsetX * float32(i+1) / float32(steps)
		offY := s.OffsetY * float32(i+1) / float32(steps)
		spread := s.Spread * float32(i+1) / float32(steps)
		a := uint8(int(s.Color.A) / steps)
		c := color.NRGBA{R: s.Color.R, G: s.Color.G, B: s.Color.B, A: a}
		drawFilledRect(screen, x+offX-spread, y+offY-spread, w+spread*2, h+spread*2, c)
	}
}

func drawFilledRect(screen *ebiten.Image, x, y, w, h float32, fill color.NRGBA) {
	if w <= 0 || h <= 0 || fill.A == 0 {
		return
	}
	vector.DrawFilledRect(screen, x, y, w, h, fill, true)
}

func renderText(rctx *RenderCtx, elem *Element) {
	if elem.textContent == "" {
		return
	}

	w, h := rctx.W, rctx.H
	if w < 1 {
		w = 1
	}
	if h < 18 {
		h = 18
	}

	hasBg := elem.bgColor.A > 0
	hasBorder := elem.borderClr.A > 0 && elem.borderClr != elem.bgColor

	if hasBg && w > 0 && h > 0 {
		bg := elem.bgColor
		bg.A = uint8(255 * elem.opacity)
		r := elem.radius
		if r > 0 {
			filledRoundedRect(screen(rctx), rctx.X, rctx.Y, w, h, r, bg)
		} else {
			drawFilledRect(screen(rctx), rctx.X, rctx.Y, w, h, bg)
		}
		if hasBorder {
			border := elem.borderClr
			border.A = uint8(255 * elem.opacity)
			vector.StrokeRect(screen(rctx), rctx.X, rctx.Y, w, h, 1, border, true)
		}
	}

	txtClr := elem.textClr
	if txtClr.A == 0 {
		txtClr = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	}
	txtClr.A = uint8(255 * elem.opacity)

	drawTextAligned(screen(rctx), elem.textContent, rctx.Font, rctx.X, rctx.Y, w, h, txtClr, elem.textAlign, elem.lineHeight)

	if elem.underline {
		tw := textLen(elem.textContent, rctx.Font)
		tx := rctx.X + 4.0
		if elem.textAlign == AlignCenter {
			tx = rctx.X + (w-tw)/2
		} else if elem.textAlign == AlignRight {
			tx = rctx.X + w - tw - 4
		}
		uy := rctx.Y + h - 4
		vector.StrokeLine(screen(rctx), tx, uy, tx+tw, uy, 1, txtClr, true)
	}
}

func renderStack(rctx *RenderCtx, elem *Element) {
	renderBox(rctx, elem)
}

func renderSeparator(rctx *RenderCtx, elem *Element) {
	if rctx.W <= 0 || rctx.H <= 0 {
		return
	}
	clr := elem.borderClr
	if clr.A == 0 {
		clr = color.NRGBA{R: 203, G: 213, B: 225, A: 255}
	}
	clr.A = uint8(255 * elem.opacity)
	vector.StrokeLine(screen(rctx), rctx.X+4, rctx.Y+rctx.H/2, rctx.X+rctx.W-4, rctx.Y+rctx.H/2, 1, clr, true)
}

func renderCheckbox(rctx *RenderCtx, elem *Element) {
	x, y, w, h := rctx.X, rctx.Y, rctx.W, rctx.H
	boxSize := float32(18)
	if h < boxSize {
		boxSize = h
	}
	bx := x + 2
	by := y + h/2 - boxSize/2

	bg := themeColor(50)
	borderClr := color.NRGBA{R: 203, G: 213, B: 225, A: 255}
	if elem.checked {
		bg = color.NRGBA{R: 49, G: 130, B: 206, A: 255}
		borderClr = bg
	}

	filledRoundedRect(rctx.Screen, bx, by, boxSize, boxSize, 4, bg)
	vector.StrokeRect(rctx.Screen, bx, by, boxSize, boxSize, 1.5, borderClr, true)

	if elem.checked {
		checkClr := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		vector.StrokeLine(rctx.Screen, bx+4, by+boxSize/2, bx+boxSize/2, by+boxSize-5, 2, checkClr, true)
		vector.StrokeLine(rctx.Screen, bx+boxSize/2, by+boxSize-5, bx+boxSize-3, by+4, 2, checkClr, true)
	}

	label := elem.TextContent()
	if label != "" {
		txtClr := color.NRGBA{R: 55, G: 65, B: 81, A: 255}
		drawTextAligned(rctx.Screen, label, rctx.Font, bx+boxSize+8, y, w-boxSize-8, h, txtClr, AlignLeft, 1.2)
	}
}

func renderSwitch(rctx *RenderCtx, elem *Element) {
	x, y, w, h := rctx.X, rctx.Y, rctx.W, rctx.H
	tw := float32(40)
	th := float32(22)
	tx := x + 2
	ty := y + h/2 - th/2

	bg := color.NRGBA{R: 203, G: 213, B: 225, A: 255}
	knobX := tx + 2
	if elem.checked {
		bg = color.NRGBA{R: 49, G: 130, B: 206, A: 255}
		knobX = tx + tw - th + 2
	}

	filledRoundedRect(rctx.Screen, tx, ty, tw, th, th/2, bg)
	vector.DrawFilledCircle(rctx.Screen, knobX+th/2-2, ty+th/2, th/2-3, color.NRGBA{R: 255, G: 255, B: 255, A: 255}, true)

	label := elem.TextContent()
	if label != "" {
		txtClr := color.NRGBA{R: 55, G: 65, B: 81, A: 255}
		drawTextAligned(rctx.Screen, label, rctx.Font, tx+tw+8, y, w-tw-8, h, txtClr, AlignLeft, 1.2)
	}
}

func renderInput(rctx *RenderCtx, elem *Element) {
	x, y, w, h := rctx.X, rctx.Y, rctx.W, rctx.H
	if w <= 0 || h <= 0 {
		return
	}

	bg := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	borderClr := color.NRGBA{R: 203, G: 213, B: 225, A: 255}
	if rctx.Focused || rctx.Hovered {
		borderClr = color.NRGBA{R: 49, G: 130, B: 206, A: 255}
	}

	drawRoundedBox(rctx.Screen, x, y, w, h, 4, bg, borderClr, Shadow{})

	m := rctx.Font.Metrics()
	asc := float32(m.Ascent) / 64
	desc := float32(m.Descent) / 64
	lineH := (asc + desc) * 1.4
	padX := float32(6)
	padY := float32(5)
	availW := w - padX*2

	txt := elem.TextContent()
	isPlaceholder := txt == ""
	if isPlaceholder {
		txt = elem.Placeholder()
	}

	if !isPlaceholder {
		lines := splitLines(txt)
		wrapped := make([]string, 0)
		for _, line := range lines {
			wrapped = append(wrapped, wrapLine(line, availW, rctx.Font)...)
		}
		if len(wrapped) == 0 {
			wrapped = []string{""}
		}

		maxLines := int((h - padY*2) / lineH)
		if maxLines < 1 {
			maxLines = 1
		}
		if len(wrapped) > maxLines {
			wrapped = wrapped[len(wrapped)-maxLines:]
		}

		for i, line := range wrapped {
			ly := y + padY + float32(i)*lineH + asc
			if ly > y+h-padY {
				break
			}

			displayLine := line
			tw := textLen(line, rctx.Font)

			if i == len(wrapped)-1 && tw > availW && rctx.Focused {
				rs := []rune(line)
				for textLen(string(rs), rctx.Font) > availW-4 && len(rs) > 1 {
					rs = rs[1:]
				}
				displayLine = string(rs)
				tw = textLen(displayLine, rctx.Font)
			}

			if i == len(wrapped)-1 {
			}

			text.Draw(rctx.Screen, displayLine, rctx.Font, int(x+padX), int(ly), color.NRGBA{R: 55, G: 65, B: 81, A: 255})
		}

		if rctx.Focused {
			pos := elem._cursorPos
			txtRunes := []rune(txt)
			if pos < 0 {
				pos = 0
			}
			if pos > len(txtRunes) {
				pos = len(txtRunes)
			}

			var caretX, caretY float32
			caretY = y + padY + float32(len(wrapped)-1)*lineH + asc
			caretLine := len(wrapped) - 1

			cumulative := 0
			for i, line := range wrapped {
				lineRunes := len([]rune(line))
				if cumulative+lineRunes >= pos && i < len(wrapped)-1 {
					if cumulative+lineRunes == pos && i+1 < len(wrapped) {
						continue
					}
					caretLine = i
					break
				}
				cumulative += lineRunes
				if i < len(wrapped)-1 && cumulative >= pos {
					caretLine = i
					break
				}
			}

			if caretLine >= 0 && caretLine < len(wrapped) {
				caretY = y + padY + float32(caretLine)*lineH + asc
				charsBefore := pos
				for i := 0; i < caretLine; i++ {
					charsBefore -= len([]rune(wrapped[i]))
				}
				if charsBefore < 0 {
					charsBefore = 0
				}
				if charsBefore > len([]rune(wrapped[caretLine])) {
					charsBefore = len([]rune(wrapped[caretLine]))
				}
				lineRunes := []rune(wrapped[caretLine])
				if charsBefore <= len(lineRunes) {
					caretX = textLen(string(lineRunes[:charsBefore]), rctx.Font)
				}
			}

			if caretX > availW {
				caretX = availW - 4
			}
			cx := x + padX + caretX
			if cx > x+w-padX {
				cx = x + w - padX
			}

			cursorClr := color.NRGBA{R: 49, G: 130, B: 206, A: 255}
			showCursor := (rctx.Frame/30)%2 == 0
			if showCursor {
				vector.StrokeLine(rctx.Screen, cx, caretY-asc+2, cx, caretY+desc-1, 1.5, cursorClr, true)
			}
		}
	} else if isPlaceholder && txt != "" {
		if textLen(txt, rctx.Font) > availW {
			rs := []rune(txt)
			for textLen(string(rs), rctx.Font) > availW-12 && len(rs) > 1 {
				rs = rs[:len(rs)-1]
			}
			txt = string(rs) + "..."
		}
		txtClr := color.NRGBA{R: 160, G: 174, B: 192, A: 255}
		text.Draw(rctx.Screen, txt, rctx.Font, int(x+padX), int(y+padY+asc), txtClr)
	}
}

func wrapLine(line string, maxW float32, face font.Face) []string {
	if line == "" {
		return []string{""}
	}
	var result []string
	rs := []rune(line)
	start := 0
	lastSpace := -1

	for i := 0; i <= len(rs); i++ {
		if i < len(rs) && rs[i] == ' ' {
			lastSpace = i
		}

		endIdx := i
		if i < len(rs) {
			endIdx = i + 1
		} else {
			endIdx = len(rs)
		}

		if endIdx > start {
			seg := string(rs[start:endIdx])
			if textLen(seg, face) > maxW {
				breakAt := i - 1
				if lastSpace > start {
					breakAt = lastSpace
				}
				if breakAt <= start {
					breakAt = start + 1
				}
				if breakAt > len(rs) {
					breakAt = len(rs)
				}
				result = append(result, string(rs[start:breakAt]))
				start = breakAt
				i = start - 1
				lastSpace = -1
			}
		}
	}

	if start < len(rs) {
		result = append(result, string(rs[start:]))
	}
	if len(result) == 0 {
		result = append(result, line)
	}
	return result
}

func splitWords(s string) []string {
	var words []string
	current := ""
	for _, r := range s {
		if r == ' ' {
			if current != "" {
				words = append(words, current)
				current = ""
			}
		} else {
			current += string(r)
		}
	}
	if current != "" {
		words = append(words, current)
	}
	return words
}

func renderButton(rctx *RenderCtx, elem *Element) {
	x, y, w, h := rctx.X, rctx.Y, rctx.W, rctx.H
	if w <= 0 || h <= 0 {
		return
	}

	bg := elem.bgColor
	if bg.A > 0 {
		bg.A = uint8(255 * elem.opacity)
	}
	border := elem.borderClr
	if border.A > 0 {
		border.A = uint8(255 * elem.opacity)
	}

	if elem.radius > 0 {
		drawRoundedBox(rctx.Screen, x, y, w, h, elem.radius, bg, border, elem.shadow)
	} else {
		if bg.A > 0 {
			drawFilledRect(rctx.Screen, x, y, w, h, bg)
		}
		if border.A > 0 && border != bg {
			vector.StrokeRect(rctx.Screen, x, y, w, h, 1.5, border, true)
		}
	}

	label := elem.TextContent()
	if label != "" {
		txtClr := elem.textClr
		if txtClr.A == 0 {
			txtClr = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		}
		txtClr.A = uint8(255 * elem.opacity)
		drawTextAligned(rctx.Screen, label, rctx.Font, x, y, w, h, txtClr, AlignCenter, 1.2)
	}
}

func splitLines(s string) []string {
	if s == "" {
		return []string{""}
	}
	var lines []string
	current := ""
	for _, r := range s {
		if r == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(r)
		}
	}
	lines = append(lines, current)
	return lines
}

func themeColor(shade uint8) color.NRGBA {
	return color.NRGBA{R: 255, G: 255, B: 255, A: 255}
}

func drawTextAligned(screen *ebiten.Image, str string, face font.Face, x, y, w, h float32, clr color.NRGBA, align TextAlign, lh float32) {
	if lh <= 0 {
		lh = 1.2
	}
	m := face.Metrics()
	ascent := float32(m.Ascent) / 64
	descent := float32(m.Descent) / 64
	tw := textLen(str, face)

	var tx float32
	switch align {
	case AlignCenter:
		tx = x + (w-tw)/2
	case AlignRight:
		tx = x + w - tw - 4
	default:
		tx = x + 4
	}

	th := (ascent + descent) * lh
	baseline := y + (h-th)/2 + ascent
	text.Draw(screen, str, face, int(tx), int(baseline), clr)
}

func textLen(str string, face font.Face) float32 {
	var tw float32
	for _, r := range str {
		a, ok := face.GlyphAdvance(r)
		if ok {
			tw += float32(a) / 64
		} else {
			tw += 4
		}
	}
	return tw
}

func screen(rctx *RenderCtx) *ebiten.Image { return rctx.Screen }

func RGBA(r, g, b uint8) color.NRGBA      { return color.NRGBA{R: r, G: g, B: b, A: 255} }
func RGBAHex(s string) color.NRGBA         { return hexColor(s) }

func hexColor(hex string) color.NRGBA {
	if len(hex) == 0 || hex[0] != '#' {
		return color.NRGBA{R: 128, G: 128, B: 128, A: 255}
	}
	hex = hex[1:]
	var r, g, b uint8
	switch len(hex) {
	case 6:
		v := uint32(0)
		for _, c := range hex {
			v = v*16 + uint32(hd(byte(c)))
		}
		r = uint8(v >> 16)
		g = uint8(v >> 8)
		b = uint8(v)
	case 3:
		v := uint32(0)
		for _, c := range hex {
			v = v*16 + uint32(hd(byte(c)))
		}
		r = uint8((v>>8)&0xF) * 17
		g = uint8((v>>4)&0xF) * 17
		b = uint8(v&0xF) * 17
	default:
		return color.NRGBA{R: 128, G: 128, B: 128, A: 255}
	}
	return color.NRGBA{R: r, G: g, B: b, A: 255}
}

func hd(c byte) uint8 {
	switch {
	case c >= '0' && c <= '9':
		return c - '0'
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func darkenColor(c color.NRGBA, amount float64) color.NRGBA {
	return color.NRGBA{
		R: uint8(math.Max(0, float64(c.R)*(1-amount))),
		G: uint8(math.Max(0, float64(c.G)*(1-amount))),
		B: uint8(math.Max(0, float64(c.B)*(1-amount))),
		A: c.A,
	}
}
