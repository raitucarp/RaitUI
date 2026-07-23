package component

import (
	"raitui/core"
	"raitui/props"
)

func VStack() *core.Element {
	return Stack(props.FlexDirectionColumn)
}
