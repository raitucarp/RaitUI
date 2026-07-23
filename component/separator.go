package component

import (
	"raitui/core"
	"raitui/theme"
)

func Separator() *core.Element {
	return Box().
		Width("100").
		Height("1").
		BackgroundColor(theme.Gray200)
}
