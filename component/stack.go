package component

import (
	goda "goda"
	"raitui/core"
)

func Stack(dir goda.FlexDirection) *core.Element {
	return Flex(dir)
}
