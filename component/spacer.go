package component

import (
	"raitui/core"
)

func Spacer() *core.Element {
	return Box().FlexGrow(1)
}
