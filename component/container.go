package component

import (
	"raitui/core"
)

func Container() *core.Element {
	return Box().
		MaxWidth("1280").
		MarginX("16")
}
