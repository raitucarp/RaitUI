package keyboard

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Handler struct {
	keys      map[ebiten.Key]int
	shortcuts []Shortcut
}

type Shortcut struct {
	Keys    []ebiten.Key
	Ctrl    bool
	Shift   bool
	Handler func()
	Global  bool
}

func New() *Handler {
	return &Handler{
		keys: make(map[ebiten.Key]int),
	}
}

func (h *Handler) Update() {
	for k := range h.keys {
		if ebiten.IsKeyPressed(k) {
			h.keys[k]++
		} else {
			delete(h.keys, k)
		}
	}
}

func (h *Handler) KeyPressed(key ebiten.Key) bool {
	_, held := h.keys[key]
	if !held && inpututil.IsKeyJustPressed(key) {
		h.keys[key] = 1
		return true
	}
	return false
}

func (h *Handler) KeyHeld(key ebiten.Key, initialDelay, repeatRate int) bool {
	dur, held := h.keys[key]
	if !held {
		if inpututil.IsKeyJustPressed(key) {
			h.keys[key] = 1
			return true
		}
		return false
	}
	if dur <= 1 {
		return true
	}
	elapsed := dur - initialDelay
	if elapsed > 0 && elapsed%repeatRate == 0 {
		return true
	}
	return false
}

func (h *Handler) IsHeld(key ebiten.Key) bool {
	_, ok := h.keys[key]
	return ok
}

func (h *Handler) InputChars() []rune {
	return ebiten.AppendInputChars(nil)
}

func (h *Handler) ModCtrl() bool {
	return ebiten.IsKeyPressed(ebiten.KeyControl) || ebiten.IsKeyPressed(ebiten.KeyControlLeft) || ebiten.IsKeyPressed(ebiten.KeyControlRight)
}

func (h *Handler) ModShift() bool {
	return ebiten.IsKeyPressed(ebiten.KeyShift) || ebiten.IsKeyPressed(ebiten.KeyShiftLeft) || ebiten.IsKeyPressed(ebiten.KeyShiftRight)
}

func (h *Handler) RegisterShortcut(s Shortcut) {
	h.shortcuts = append(h.shortcuts, s)
}

func (h *Handler) Clear() {
	h.keys = make(map[ebiten.Key]int)
}

func (h *Handler) CheckShortcuts(focused bool) {
	for _, s := range h.shortcuts {
		if !s.Global && !focused {
			continue
		}
		allPressed := true
		for _, k := range s.Keys {
			if !inpututil.IsKeyJustPressed(k) {
				allPressed = false
				break
			}
		}
		if !allPressed {
			continue
		}
		if s.Ctrl && !h.ModCtrl() {
			continue
		}
		if s.Shift && !h.ModShift() {
			continue
		}
		s.Handler()
	}
}

type KeyRune struct {
	Key   ebiten.Key
	Lower rune
	Upper rune
}

var keyRunes = []KeyRune{
	{ebiten.KeyA, 'a', 'A'}, {ebiten.KeyB, 'b', 'B'}, {ebiten.KeyC, 'c', 'C'},
	{ebiten.KeyD, 'd', 'D'}, {ebiten.KeyE, 'e', 'E'}, {ebiten.KeyF, 'f', 'F'},
	{ebiten.KeyG, 'g', 'G'}, {ebiten.KeyH, 'h', 'H'}, {ebiten.KeyI, 'i', 'I'},
	{ebiten.KeyJ, 'j', 'J'}, {ebiten.KeyK, 'k', 'K'}, {ebiten.KeyL, 'l', 'L'},
	{ebiten.KeyM, 'm', 'M'}, {ebiten.KeyN, 'n', 'N'}, {ebiten.KeyO, 'o', 'O'},
	{ebiten.KeyP, 'p', 'P'}, {ebiten.KeyQ, 'q', 'Q'}, {ebiten.KeyR, 'r', 'R'},
	{ebiten.KeyS, 's', 'S'}, {ebiten.KeyT, 't', 'T'}, {ebiten.KeyU, 'u', 'U'},
	{ebiten.KeyV, 'v', 'V'}, {ebiten.KeyW, 'w', 'W'}, {ebiten.KeyX, 'x', 'X'},
	{ebiten.KeyY, 'y', 'Y'}, {ebiten.KeyZ, 'z', 'Z'},
	{ebiten.Key0, '0', ')'}, {ebiten.Key1, '1', '!'}, {ebiten.Key2, '2', '@'},
	{ebiten.Key3, '3', '#'}, {ebiten.Key4, '4', '$'}, {ebiten.Key5, '5', '%'},
	{ebiten.Key6, '6', '^'}, {ebiten.Key7, '7', '&'}, {ebiten.Key8, '8', '*'},
	{ebiten.Key9, '9', '('},
	{ebiten.KeyPeriod, '.', '>'}, {ebiten.KeyComma, ',', '<'},
	{ebiten.KeySlash, '/', '?'}, {ebiten.KeySemicolon, ';', ':'},
	{ebiten.KeyQuote, '\'', '"'}, {ebiten.KeyBracketLeft, '[', '{'},
	{ebiten.KeyBracketRight, ']', '}'}, {ebiten.KeyBackslash, '\\', '|'},
	{ebiten.KeyMinus, '-', '_'}, {ebiten.KeyEqual, '=', '+'},
	{ebiten.KeyGraveAccent, '`', '~'},
}

func KeyRunes() []KeyRune {
	return keyRunes
}

func InsertRune(s string, pos int, ch rune) (string, int) {
	rs := []rune(s)
	if pos > len(rs) {
		pos = len(rs)
	}
	if pos < 0 {
		pos = 0
	}
	rs = append(rs[:pos], append([]rune{ch}, rs[pos:]...)...)
	return string(rs), pos + 1
}

func DeleteAt(s string, pos int) (string, int) {
	rs := []rune(s)
	if pos > len(rs) {
		pos = len(rs)
	}
	if pos > 0 && len(rs) > 0 {
		rs = append(rs[:pos-1], rs[pos:]...)
		return string(rs), pos - 1
	}
	return s, pos
}

func DeleteForward(s string, pos int) (string, int) {
	rs := []rune(s)
	if pos > len(rs) {
		pos = len(rs)
	}
	if pos < len(rs) {
		rs = append(rs[:pos], rs[pos+1:]...)
	}
	return string(rs), pos
}
