package component

import (
	"raitui/core"
	"raitui/props"
)

func Portal() *core.Element {
	return Box().
		Position(props.PositionAbsolute).
		Width("100%").Height("100%")
}

func OpenPortal(ctx *core.Context, portal *core.Element) {
	ctx.PushLayer(portal)
}

func ClosePortal(ctx *core.Context, portal *core.Element) {
	ctx.PopLayer(portal)
}
