package core

import (
	"testing"

	goda "goda"
)

func TestLayout_GapVStack(t *testing.T) {
	vstack := NewElement(TypeVStack)
	vstack.FlexDirection(goda.FlexDirectionColumn)
	vstack.Width("200")
	vstack.Gap("4")

	for _, s := range []string{"Revenue", "$12,430", "+12.5%"} {
		txt := NewElement(TypeText)
		txt.SetTextContent(s)
		txt.GNode.SetWidth(float32(len([]rune(s))) * 8)
		txt.GNode.SetHeight(20)
		vstack.AppendChild(txt)
	}

	goda.CalculateNodeLayout(vstack.GNode, 200, 200, goda.DirectionLTR)

	for i := 0; i < vstack.ChildrenCount(); i++ {
		child := vstack.ChildAt(i)
		lo := child.GNode.LayoutOut()
		expectedY := float32(i) * (20 + 4)
		if lo.Top != expectedY {
			t.Errorf("GAP FAIL %q: y=%.0f, expected y=%.0f", child.TextContent(), lo.Top, expectedY)
		} else {
			t.Logf("OK %q: y=%.0f", child.TextContent(), lo.Top)
		}
	}
}

func TestLayout_CardInHStack(t *testing.T) {
	root := NewElement(TypeVStack)
	root.FlexDirection(goda.FlexDirectionColumn)
	root.Width("800").Height("600")
	root.Padding("24").Gap("16")

	row := NewElement(TypeHStack)
	row.FlexDirection(goda.FlexDirectionRow)
	row.Gap("12")

	for range 3 {
		card := NewElement(TypeBox)
		card.Padding("16")
		card.GNode.SetFlexGrow(1)

		vst := NewElement(TypeVStack)
		vst.FlexDirection(goda.FlexDirectionColumn)
		vst.Gap("4")

		txt := NewElement(TypeText)
		txt.SetTextContent("Test")
		txt.GNode.SetWidth(40).SetHeight(20)
		vst.AppendChild(txt)

		card.AppendChild(vst)
		row.AppendChild(card)
	}

	root.AppendChild(row)

	goda.CalculateNodeLayout(root.GNode, 800, 600, goda.DirectionLTR)

	rowLO := row.GNode.LayoutOut()
	t.Logf("Row: pos=(%.0f,%.0f) size=(%.0fx%.0f)", rowLO.Left, rowLO.Top, rowLO.Width, rowLO.Height)

	// Cards should be in a ROW (x positions different, same y)
	for i := 0; i < row.ChildrenCount(); i++ {
		card := row.ChildAt(i)
		cardLO := card.GNode.LayoutOut()
		t.Logf("  Card[%d]: pos=(%.0f,%.0f) size=(%.0fx%.0f)", i, cardLO.Left, cardLO.Top, cardLO.Width, cardLO.Height)

		// In a row, all cards should have the same Y
		if i > 0 {
			prev := row.ChildAt(i - 1)
			prevLO := prev.GNode.LayoutOut()
			if cardLO.Top != prevLO.Top {
				t.Errorf("ROW FAIL: Card[%d].y=%.0f != Card[0].y=%.0f", i, cardLO.Top, prevLO.Top)
			}
			if cardLO.Left <= prevLO.Left {
				t.Errorf("ROW FAIL: Card[%d].x=%.0f <= Card[%d].x=%.0f", i, cardLO.Left, i-1, prevLO.Left)
			}
		}
	}
}

func TestLayout_FullDashboard(t *testing.T) {
	root := NewElement(TypeVStack)
	root.FlexDirection(goda.FlexDirectionColumn)
	root.Width("800").Height("600")
	root.Padding("20").Gap("16")

	header := NewElement(TypeBox)
	header.Width("100%").Padding("16")
	header.AppendChild(NewElement(TypeText))

	body := NewElement(TypeHStack)
	body.FlexDirection(goda.FlexDirectionRow)
	body.GNode.SetFlexGrow(1)
	body.Gap("16")

	side := NewElement(TypeBox)
	side.Width("180").Padding("8")
	body.AppendChild(side)

	main := NewElement(TypeVStack)
	main.FlexDirection(goda.FlexDirectionColumn)
	main.GNode.SetFlexGrow(1)
	main.Gap("16")

	card := NewElement(TypeBox)
	card.Width("100%").Padding("20")
	main.AppendChild(card)

	stats := NewElement(TypeHStack)
	stats.FlexDirection(goda.FlexDirectionRow)
	stats.Gap("16")
	for range 3 {
		sc := NewElement(TypeBox)
		sc.Padding("16")
		sc.GNode.SetFlexGrow(1)
		sc.Helper_SetSize(100, 20)
		stats.AppendChild(sc)
	}
	main.AppendChild(stats)

	tags := NewElement(TypeBox)
	tags.Width("100%").Padding("16")
	main.AppendChild(tags)

	body.AppendChild(main)
	root.AppendChild(header)
	root.AppendChild(body)

	goda.CalculateNodeLayout(root.GNode, 800, 600, goda.DirectionLTR)

	bodyLO := body.GNode.LayoutOut()
	t.Logf("Body: pos=(%.0f,%.0f) size=(%.0fx%.0f)", bodyLO.Left, bodyLO.Top, bodyLO.Width, bodyLO.Height)

	sideLO := side.GNode.LayoutOut()
	mainLO := main.GNode.LayoutOut()
	t.Logf("Sidebar: pos=(%.0f,%.0f) size=(%.0fx%.0f)", sideLO.Left, sideLO.Top, sideLO.Width, sideLO.Height)
	t.Logf("Main: pos=(%.0f,%.0f) size=(%.0fx%.0f)", mainLO.Left, mainLO.Top, mainLO.Width, mainLO.Height)

	// Sidebar and main should be in a row (same Y, different X)
	if sideLO.Top != mainLO.Top {
		t.Errorf("ALIGN FAIL: sidebar.y=%.0f != main.y=%.0f", sideLO.Top, mainLO.Top)
	}
	if mainLO.Left <= sideLO.Left {
		t.Errorf("ALIGN FAIL: main.x=%.0f <= sidebar.x=%.0f", mainLO.Left, sideLO.Left)
	}

	// Stats cards should be in a row
	if stats.ChildrenCount() >= 2 {
		c0 := stats.ChildAt(0).GNode.LayoutOut()
		c1 := stats.ChildAt(1).GNode.LayoutOut()
		if c0.Top != c1.Top {
			t.Errorf("STATS FAIL: card[0].y=%.0f != card[1].y=%.0f", c0.Top, c1.Top)
		}
	}
}

func (e *Element) Helper_SetSize(w, h float32) {
	e.GNode.SetWidth(w).SetHeight(h)
}

func TestLayout_TextDoesNotShrink(t *testing.T) {
	// Simulate window resize: narrow parent should NOT shrink fixed-width text

	for _, availW := range []float32{800, 400, 200, 100} {
		parent := NewElement(TypeVStack)
		parent.FlexDirection(goda.FlexDirectionColumn)
		parent.GNode.SetAlignItems(goda.AlignFlexStart)
		parent.Width("100%")
		parent.Padding("20")

		txt := NewElement(TypeText)
		txt.SetTextContent("Hello World")
		txt.GNode.SetWidth(99)
		txt.GNode.SetMinWidth(99)
		txt.GNode.SetHeight(22)
		txt.GNode.SetMinHeight(22)
		txt.GNode.SetFlexShrink(0)

		parent.AppendChild(txt)

		goda.CalculateNodeLayout(parent.GNode, availW, 600, goda.DirectionLTR)

		txtLO := txt.GNode.LayoutOut()
		t.Logf("Available width=%d: text width=%.0f height=%.0f", int(availW), txtLO.Width, txtLO.Height)

		if txtLO.Width != 99 {
			t.Errorf("SHRINK BUG: availW=%d, text width=%.0f, expected 99", int(availW), txtLO.Width)
		}
	}
}

func TestLayout_StatCardsDontShrink(t *testing.T) {
	for _, availW := range []float32{800, 500, 300} {
		row := NewElement(TypeHStack)
		row.FlexDirection(goda.FlexDirectionRow)
		row.Gap("16")

		for range 3 {
			card := NewElement(TypeBox)
			card.Padding("16")
			card.GNode.SetFlexShrink(0)

			vst := NewElement(TypeVStack)
			vst.FlexDirection(goda.FlexDirectionColumn)
			vst.Gap("4")

			txt := NewElement(TypeText)
			txt.SetTextContent("Revenue")
			txt.GNode.SetWidth(56).SetMinWidth(56).SetHeight(22).SetMinHeight(22)
			txt.GNode.SetFlexShrink(0)
			vst.AppendChild(txt)

			card.AppendChild(vst)
			row.AppendChild(card)
		}

		goda.CalculateNodeLayout(row.GNode, availW, 200, goda.DirectionLTR)

		for i := 0; i < row.ChildrenCount(); i++ {
			card := row.ChildAt(i)
			cardLO := card.GNode.LayoutOut()
			vst := card.ChildAt(0)
			vstLO := vst.GNode.LayoutOut()
			txt := vst.ChildAt(0)
			txtLO := txt.GNode.LayoutOut()

			t.Logf("availW=%d card[%d]: card=%.0fx%.0f vst=%.0fx%.0f txt=%.0fx%.0f",
				int(availW), i, cardLO.Width, cardLO.Height, vstLO.Width, vstLO.Height, txtLO.Width, txtLO.Height)

			if txtLO.Width != 56 {
				t.Errorf("TEXT SHRINK: availW=%d, txt[%d] width=%.0f, expected 56", int(availW), i, txtLO.Width)
			}
		}
	}
}
