package component

import (
	"raitui/core"
	"raitui/props"
)

func HStack() *core.Element {
	return Stack(props.FlexDirectionRow)
}
