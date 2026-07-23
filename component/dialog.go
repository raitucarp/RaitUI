package component

import (
	"raitui/core"
	"raitui/props"
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
	backdrop := core.NewElement(core.TypeDialog)
	backdrop.BackgroundColor(colorWithAlpha(theme.Black, 60))
	backdrop.Width("100%").Height("100%")
	backdrop.FlexDirection(props.FlexDirectionColumn)
	backdrop.JustifyContent(props.JustifyCenter)
	backdrop.AlignItems(props.AlignCenter)
	backdrop.FlexShrink(0)
	backdrop.Visible(false)

	card := Card().Padding("24").Gap("16")
	header := Box()
	body := VStack().FlexGrow(1).FlexShrink(1).Gap("12")
	footer := HStack().Gap("8")

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
	t := Text(label).FontSize(18).TextColor(theme.Gray800)
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
