package core

func (e *Element) OnClick(handler func()) *Element {
	e.onClick = handler
	return e
}

func (e *Element) OnHover(handler func(entered bool)) *Element {
	e.onHover = handler
	return e
}

func (e *Element) SetUnderline(v bool) *Element {
	e.underline = v
	return e
}

func (e *Element) IsUnderline() bool { return e.underline }

func (e *Element) SetChecked(v bool) *Element {
	e.checked = v
	return e
}

func (e *Element) IsChecked() bool { return e.checked }

func (e *Element) SetPlaceholder(p string) *Element {
	e.placeholder = p
	return e
}

func (e *Element) Placeholder() string { return e.placeholder }
