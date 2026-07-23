package component

import (
	"raitui/core"
)

func AspectRatio(ratio float32) *core.Element {
	return Box().AspectRatio(ratio)
}
