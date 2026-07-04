package component

import (
	goda "goda"

	"raitui/core"
	"raitui/theme"
)

type dialogRoot struct {
	backdrop *core.Element
	card     *core.Element
	header   *core.Element
	body     *core.Element
	footer   *core.Element
}

func DialogRoot() *dialogRoot {
	backdrop := core.NewElement(core.TypeBox)
	backdrop.BackgroundColor(colorWithAlpha(theme.Black, 60))
	backdrop.Width("100%").Height("100%")
	backdrop.GNode.SetFlexDirection(goda.FlexDirectionColumn)
	backdrop.GNode.SetJustifyContent(goda.JustifyCenter)
	backdrop.GNode.SetAlignItems(goda.AlignCenter)
	backdrop.GNode.SetFlexShrink(0)
	backdrop.Visible(false)

	card := core.NewElement(core.TypeBox)
	card.BackgroundColor(theme.White)
	card.BorderRadius(12)
	card.Padding("24").Gap("16")
	card.GNode.SetMinWidth(400)
	card.GNode.SetFlexDirection(goda.FlexDirectionColumn)
	card.GNode.SetFlexShrink(0)

	header := core.NewElement(core.TypeBox)
	header.GNode.SetFlexShrink(0)

	body := core.NewElement(core.TypeVStack)
	body.FlexDirection(goda.FlexDirectionColumn)
	body.GNode.SetFlexGrow(1).SetFlexShrink(1)
	body.Gap("12")

	footer := core.NewElement(core.TypeHStack)
	footer.FlexDirection(goda.FlexDirectionRow)
	footer.Gap("8")
	footer.GNode.SetFlexShrink(0)

	card.AppendChild(header)
	card.AppendChild(body)
	card.AppendChild(footer)
	backdrop.Children(card)

	return &dialogRoot{
		backdrop: backdrop,
		card:     card,
		header:   header,
		body:     body,
		footer:   footer,
	}
}

func (d *dialogRoot) Title(label string) *dialogRoot {
	t := Text(label)
	t.FontSize(18)
	t.TextColor(theme.Gray800)
	t.GNode.SetFlexShrink(0)
	d.header.Children(t)
	return d
}

func (d *dialogRoot) Body(children ...*core.Element) *dialogRoot {
	for _, c := range children {
		d.body.AppendChild(c)
	}
	return d
}

func (d *dialogRoot) Footer(children ...*core.Element) *dialogRoot {
	for _, c := range children {
		d.footer.AppendChild(c)
	}
	return d
}

func (d *dialogRoot) CloseButton(label string) *core.Element {
	btn := Button(label)
	btn.OnClick(func() { d.backdrop.Close() })
	return btn
}

func (d *dialogRoot) ActionTrigger(label string) *core.Element {
	btn := OutlineButton(label)
	btn.OnClick(func() { d.backdrop.Close() })
	return btn
}

func (d *dialogRoot) Trigger(trigger *core.Element) *core.Element {
	trigger.OnClick(func() { d.backdrop.Open() })
	return trigger
}

func (d *dialogRoot) Element() *core.Element {
	return d.backdrop
}

func (d *dialogRoot) Open()  { d.backdrop.Open() }
func (d *dialogRoot) Close() { d.backdrop.Close() }

func Dialog(title string) *core.Element { return DialogRoot().Title(title).Element() }
func DialogContent(dlg *core.Element) *core.Element { return dlg }
func DialogFooter(dlg *core.Element) *core.Element { return dlg }
