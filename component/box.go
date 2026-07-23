package component

import (
		"raitui/core"
)

func Box() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.FlexShrink(0)
	return elem
}
