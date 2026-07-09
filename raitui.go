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

	Tabs               = component.Tabs
	TabList            = component.TabList
	Tab                = component.Tab
	TabPanels          = component.TabPanels
	TabPanel           = component.TabPanel
	Accordion          = component.Accordion
	AccordionItem      = component.AccordionItem
	AccordionHeader    = component.AccordionHeader
	AccordionPanel     = component.AccordionPanel
	Breadcrumb         = component.Breadcrumb
	BreadcrumbItem     = component.BreadcrumbItem
	BreadcrumbLink     = component.BreadcrumbLink
	BreadcrumbSeparator = component.BreadcrumbSeparator
	Table               = component.Table
	TableHead           = component.TableHead
	TableBody           = component.TableBody
	TableRow            = component.TableRow
	TableCell           = component.TableCell
	TableHeaderCell     = component.TableHeaderCell
	Timeline            = component.Timeline
	TimelineItem        = component.TimelineItem
	TimelineDot         = component.TimelineDot
	TimelineContent     = component.TimelineContent
	Stat                = component.Stat
	StatLabel           = component.StatLabel
	StatNumber          = component.StatNumber
	StatHelpText        = component.StatHelpText
	Skeleton            = component.Skeleton
	Tag                 = component.Tag
	TagLabel            = component.TagLabel
	TagCloseButton      = component.TagCloseButton
	Wrap                = component.Wrap
	ScrollArea          = component.ScrollAreaRoot
	ScrollAreaRoot      = component.ScrollAreaRoot
	ScrollAreaViewport  = component.ScrollAreaViewport
	ScrollAreaContent   = component.ScrollAreaContent
	ScrollAreaScrollbar = component.ScrollAreaScrollbar
	ScrollAreaThumb     = component.ScrollAreaThumb
	AspectRatio         = component.AspectRatio
	Drawer              = component.Drawer
	Collapsible         = component.Collapsible
	CollapsibleTrigger  = component.CollapsibleTrigger
	CollapsibleContent  = component.CollapsibleContent
	Pagination          = component.Pagination
	PageButton          = component.PageButton
	Steps               = component.Steps
	Step                = component.Step
	StepIndicator       = component.StepIndicator
	StepSeparator       = component.StepSeparator
	CloseButton         = component.CloseButton
	Status              = component.Status
	EmptyState          = component.EmptyState
	ProgressCircle      = component.ProgressCircle
	RadioGroup          = component.RadioGroup
	Radio               = component.Radio
	Image               = component.Image
	NumberInput         = component.NumberInput
	SegmentedControl    = component.SegmentedControl
	Segment             = component.Segment

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
