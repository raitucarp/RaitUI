package component

import (
	"raitui/core"
)

func Grid() *core.Element {
	return Wrap()
}

func GridItem() *core.Element {
	return Box().
		FlexShrink(0).
		FlexGrow(1).
		FlexBasis("200")
}

func SimpleGrid() *core.Element {
	return Wrap().
		Gap("16")
}

func Icon(name string) *core.Element {
	elem := Center().
		Width("24").MinWidth("24").
		Height("24").MinHeight("24").
		FlexShrink(0)
	elem.SetTextContent(name)
	return elem
}
