package component

import (
	goda "goda"

	"raitui/core"
)

func Portal() *core.Element {
	elem := core.NewElement(core.TypeBox)
	elem.GNode.SetPositionType(goda.PositionTypeAbsolute)
	elem.Width("100%").Height("100%")
	return elem
}

func OpenPortal(ctx *core.Context, portal *core.Element) {
	ctx.PushLayer(portal)
}

func ClosePortal(ctx *core.Context, portal *core.Element) {
	ctx.PopLayer(portal)
}
