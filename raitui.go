package raitui

import (
	"raitui/component"
	"raitui/core"
	"raitui/theme"
)

var (
	Box        = component.Box
	VStack     = component.VStack
	HStack     = component.HStack
	Text       = component.Text
	Flex       = component.Flex
	Center     = component.Center
	Container  = component.Container
	Separator  = component.Separator
	Spacer     = component.Spacer
	Heading    = component.Heading
	Code       = component.Code
	Link       = component.Link
	Blockquote = component.Blockquote
	List       = component.List
	Kbd        = component.Kbd
	Mark        = component.Mark
	Button      = component.Button
	OutlineButton = component.OutlineButton
	Checkbox    = component.Checkbox
	Switch      = component.Switch
	Input       = component.Input
	TextArea    = component.TextArea
	Badge       = component.Badge
	BadgeSolid  = component.BadgeSolid
	BadgeSubtle = component.BadgeSubtle
	BadgeOutline = component.BadgeOutline
	BadgeVariant = component.BadgeVariant
	Card        = component.Card
	Alert       = component.Alert
	Spinner     = component.Spinner
	Avatar      = component.Avatar
	Progress    = component.Progress
	Dialog      = component.Dialog
	DialogRoot  = component.DialogRoot
	Menu        = component.Menu
	MenuItem    = component.MenuItem
	Tooltip     = component.Tooltip
	Popper      = component.Popper
	Portal      = component.Portal
	Window      = component.Window
	App         = component.App
	Portals     = component.Portals

	DialogContent = component.DialogContent
	DialogFooter  = component.DialogFooter
	WithTooltip   = component.WithTooltip
	WithPopper    = component.WithPopper
	PopoverAt     = component.PopoverAt
	PlaceTop      = core.PlaceTop
	PlaceBottom   = core.PlaceBottom
	PlaceLeft     = core.PlaceLeft
	PlaceRight    = core.PlaceRight

	Colors = theme.Colors

	bgColor = theme.Gray50
)

func DefaultApp() *core.Element {
	return VStack().
		Width("100%").Height("100%")
}

func SetDefaultBackground(c theme.Color) {
	bgColor = c
}

func NewContext() *core.Context {
	return core.NewContext(bgColor)
}

func SetState(key string, value any)               { core.SetState(key, value) }
func GetState(key string) (any, bool)               { return core.GetState(key) }
func DeleteState(key string)                        { core.DeleteState(key) }
func SetLocalState(componentID, key string, v any)  { core.SetLocalState(componentID, key, v) }
func GetLocalState(componentID, key string) (any, bool) { return core.GetLocalState(componentID, key) }
func StateInt(key string, def int) func() int       { return core.StateInt(key, def) }
func StateStr(key string, def string) func() string { return core.StateStr(key, def) }
func StateGetter[T any](key string, def T) func() T { return core.StateGetter(key, def) }
func Setter[T any](key string) func(T)              { return core.Setter[T](key) }
