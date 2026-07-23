package component

import (
	"raitui/core"
)

func Canvas(draw func(core.CanvasContext)) *core.Element {
	elem := core.NewElement(core.TypeCanvas)
	elem.FlexShrink(0)
	elem.SetDrawFunc(draw)
	return elem
}
