package core

import (
	"testing"
)

func TestElement_BuilderChain(t *testing.T) {
	e := NewElement(TypeBox)
	e.Width("100").Height("50").Padding("8").Gap("4").
		BackgroundColor(RGBA(255, 0, 0)).
		BorderRadius(8).Opacity(0.9).
		Class("card", "active").
		Title("Hello").Lang("en")

	if !e.IsVisible() {
		t.Error("new element should be visible")
	}
	if e.TextContent() != "" {
		t.Error("box should have no text content")
	}
	if !e.HasClass("card") {
		t.Error("should have class card")
	}
	if !e.HasClass("active") {
		t.Error("should have class active")
	}
}

func TestText_Content(t *testing.T) {
	e := NewElement(TypeText)
	e.SetTextContent("Hello")
	e.FontSize(16).TextColor(RGBA(0, 0, 0))

	if e.TextContent() != "Hello" {
		t.Errorf("expected 'Hello', got %q", e.TextContent())
	}
	if e.FontSizeValue() != 16 {
		t.Errorf("expected fontSize 16, got %v", e.FontSizeValue())
	}
}

func TestChildren_Parent(t *testing.T) {
	parent := NewElement(TypeVStack)
	child1 := NewElement(TypeBox)
	child2 := NewElement(TypeText)

	parent.Children(child1, child2)

	if parent.ChildrenCount() != 2 {
		t.Errorf("expected 2 children, got %d", parent.ChildrenCount())
	}
	if child1.parent != parent {
		t.Error("child1 parent should be parent")
	}
	if child2.parent != parent {
		t.Error("child2 parent should be parent")
	}
}

func TestVisibility_Display(t *testing.T) {
	e := NewElement(TypeBox)
	if !e.IsVisible() {
		t.Error("should be visible by default")
	}
	e.Visible(false)
	if e.IsVisible() {
		t.Error("should be invisible after Visible(false)")
	}
	e.Visible(true)
	if !e.IsVisible() {
		t.Error("should be visible after Visible(true)")
	}
}

func TestOpenClose_Modal(t *testing.T) {
	e := NewElement(TypeBox)
	e.Visible(false)
	if e.IsOpen() {
		t.Error("should not be open after Visible(false)")
	}
	e.Open()
	if !e.IsOpen() {
		t.Error("should be open after Open()")
	}
	e.Close()
	if e.IsOpen() {
		t.Error("should not be open after Close()")
	}
}

func TestInternalState(t *testing.T) {
	e := NewElement(TypeBox)
	st := e.State()

	st.Set("count", 0)
	if st.Int("count", -1) != 0 {
		t.Error("initial count should be 0")
	}

	st.Set("count", 5)
	if st.Int("count", -1) != 5 {
		t.Error("count should be 5 after set")
	}

	if st.Str("name", "default") != "default" {
		t.Error("unset string should return default")
	}
	st.Set("name", "RaitUI")
	if st.Str("name", "default") != "RaitUI" {
		t.Error("name should be RaitUI")
	}

	if st.Bool("active", true) != true {
		t.Error("unset bool should return default")
	}
	st.Set("active", false)
	if st.Bool("active", true) != false {
		t.Error("active should be false")
	}
}

func TestStore_Global(t *testing.T) {
	Set("test_int", 42)
	if Int("test_int", 0) != 42 {
		t.Error("global int should be 42")
	}

	Set("test_str", "hello")
	if Str("test_str", "") != "hello" {
		t.Error("global str should be 'hello'")
	}

	Set("test_bool", true)
	if !Bool("test_bool", false) {
		t.Error("global bool should be true")
	}

	Del("test_int")
	if Int("test_int", 99) != 99 {
		t.Error("deleted key should return default 99")
	}
}

func TestPlacement_Types(t *testing.T) {
	if PlaceTop != 0 {
		t.Error("PlaceTop should be 0")
	}
	if PlaceBottom != 1 {
		t.Error("PlaceBottom should be 1")
	}
	if PlaceLeft != 2 {
		t.Error("PlaceLeft should be 2")
	}
	if PlaceRight != 3 {
		t.Error("PlaceRight should be 3")
	}
}

func TestAlign_Types(t *testing.T) {
	if AlignLeft != 0 {
		t.Error("AlignLeft should be 0")
	}
	if AlignCenter != 1 {
		t.Error("AlignCenter should be 1")
	}
	if AlignRight != 2 {
		t.Error("AlignRight should be 2")
	}
}

func TestCheckbox_Switch_State(t *testing.T) {
	cb := NewElement(TypeCheckbox)
	cb.SetChecked(false)
	if cb.IsChecked() {
		t.Error("checkbox should not be checked")
	}
	cb.SetChecked(true)
	if !cb.IsChecked() {
		t.Error("checkbox should be checked")
	}

	sw := NewElement(TypeSwitch)
	sw.SetChecked(false)
	if sw.IsChecked() {
		t.Error("switch should not be checked")
	}
	sw.SetChecked(true)
	if !sw.IsChecked() {
		t.Error("switch should be checked")
	}
}

func TestPlaceholder(t *testing.T) {
	e := NewElement(TypeInput)
	e.SetPlaceholder("Enter text")
	if e.Placeholder() != "Enter text" {
		t.Errorf("expected placeholder 'Enter text', got %q", e.Placeholder())
	}
}

func TestClick_Hover_Handlers(t *testing.T) {
	e := NewElement(TypeBox)
	clicked := false
	e.OnClick(func() { clicked = true })
	e.OnHover(func(entered bool) {})
	if clicked {
		t.Error("click should not fire immediately")
	}
}

func TestTooltipPlacement(t *testing.T) {
	e := NewElement(TypeBox)
	e.TooltipPlacement(PlaceTop)
	if e._tooltipPlace != 0 {
		t.Error("tooltip should be PlaceTop (0)")
	}
	e.TooltipPlacement(PlaceBottom)
	if e._tooltipPlace != 1 {
		t.Error("tooltip should be PlaceBottom (1)")
	}
}

func TestScroll(t *testing.T) {
	e := NewElement(TypeBox)
	if e.ScrollY() != 0 {
		t.Error("initial scroll should be 0")
	}
	e.SetScrollMax(100)
	e.Scroll(30)
	if e.ScrollY() != 30 {
		t.Errorf("scrollY should be 30, got %v", e.ScrollY())
	}
	e.Scroll(100)
	if e.ScrollY() != 100 {
		t.Errorf("scrollY should be capped at 100, got %v", e.ScrollY())
	}
	e.Scroll(-50)
	if e.ScrollY() != 50 {
		t.Errorf("scrollY should be 50 after -50, got %v", e.ScrollY())
	}
}

func TestUnderline(t *testing.T) {
	e := NewElement(TypeText)
	if e.IsUnderline() {
		t.Error("underline should default to false")
	}
	e.SetUnderline(true)
	if !e.IsUnderline() {
		t.Error("underline should be true after set")
	}
}
