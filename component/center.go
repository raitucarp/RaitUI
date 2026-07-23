package component

import (
	"raitui/core"
	"raitui/props"
)

func Center() *core.Element {
	return Flex(props.FlexDirectionColumn).
		JustifyContent(props.JustifyCenter).
		AlignItems(props.AlignCenter)
}
