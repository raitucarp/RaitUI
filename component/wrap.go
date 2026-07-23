package component

import (
	"raitui/core"
	"raitui/props"
)

func Wrap() *core.Element {
	return Flex(props.FlexDirectionRow).
		FlexWrap(props.WrapWrap).
		AlignItems(props.AlignFlexStart)
}
