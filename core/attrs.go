package core

import goda "goda"

func (e *Element) ID(id string) *Element {
	e.elemID = id
	return e
}

func (e *Element) Class(classes ...string) *Element {
	for _, c := range classes {
		e.GNode.AddClass(c)
	}
	return e
}

func (e *Element) AddClass(cls string) *Element {
	e.GNode.AddClass(cls)
	return e
}

func (e *Element) HasClass(cls string) bool {
	return e.GNode.HasClass(cls)
}

func (e *Element) GetClasses() []string {
	return e.GNode.GetClasses()
}

func (e *Element) Hidden(v bool) *Element {
	e.visible = !v
	return e
}

func (e *Element) Title(title string) *Element {
	e.tooltip = title
	return e
}

func (e *Element) Lang(lang string) *Element {
	e.lang = lang
	return e
}

func (e *Element) Dir(dir string) *Element {
	e.dir = dir
	if dir == "rtl" {
		e.GNode.SetDirection(goda.DirectionRTL)
	} else {
		e.GNode.SetDirection(goda.DirectionLTR)
	}
	return e
}

func (e *Element) TabIndex(idx int) *Element {
	e.tabIndex = idx
	return e
}
