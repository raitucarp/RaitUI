package component

import "raitui/core"

func Image() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.BorderRadius(8)
	elem.FlexShrink(0)
	return elem
}
