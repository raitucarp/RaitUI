package component

import (
	"raitui/core"
)

func Input(placeholder string) *core.Element {
	elem := core.NewElement(core.TypeInput)
	elem.SetPlaceholder(placeholder)
	elem.Width("200")
	elem.MinHeight("36")
	elem.Height("36")
	elem.FlexShrink(0)
	return elem
}

func TextArea(placeholder string) *core.Element {
	elem := core.NewElement(core.TypeInput)
	elem.SetPlaceholder(placeholder)
	elem.Width("300")
	elem.MinHeight("80")
	elem.Height("80")
	elem.FlexShrink(0)
	return elem
}
