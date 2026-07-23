# RaitUI

A GUI framework for Go, built on [Ebitengine](https://github.com/hajimehoshi/ebiten). Compose interfaces declaratively with a builder-pattern API inspired by [Chakra UI](https://chakra-ui.com/).

---

## Getting Started

### Installation

```bash
go get raitui
```

### Hello World

```go
package main

import (
    "raitui"
    "raitui/theme"
)

func main() {
    raitui.Window().
        Title("My App").
        Width(800).Height(600).
        Children(
            raitui.VStack().
                Width("100%").Height("100%").
                Padding("24").Gap("16").
                BackgroundColor(theme.Gray50).
                Children(
                    raitui.Heading("Welcome to RaitUI", 1),
                    raitui.Text("Build native GUI apps with Go and Ebitengine").FontSize(14),
                    raitui.Button("Get Started"),
                ),
        ).
        Run()
}
```

---

## Architecture

RaitUI is built on a layered architecture that keeps the layout engine (`goda`) isolated from application code:

| Layer | Role | Examples |
|-------|------|----------|
| **goda** | Flexbox layout engine | Node positioning, size calculation |
| **Primitives** | Direct layout manipulation | `Box`, `Flex`, `Stack`, `Text`, `Canvas` |
| **Composites** | Built from primitives only | `Button`, `Card`, `Dialog`, `Table`, `Tabs` |
| **Application** | Your code | Composing primitives and composites |

Only primitive components interact with `goda`. All composite components are built purely by composing primitives together. This keeps the framework modular and easy to extend.

---

## Component Reference

### Layout Primitives

These are the building blocks. They are the only components that create layout nodes.

| Component | Description | Usage |
|-----------|-------------|-------|
| `Box()` | The fundamental container. Every component is built on it. | `Box().Width("200").BackgroundColor(theme.Blue500)` |
| `Flex(dir)` | Flexible container with configurable direction. | `Flex(props.FlexDirectionRow).Gap("16")` |
| `Stack(dir)` | Flex with convenient defaults for stacking. | `Stack(props.FlexDirectionColumn)` |
| `VStack()` | Vertical stack. Children are laid out top-to-bottom. | `VStack().Gap("8")` |
| `HStack()` | Horizontal stack. Children are laid out left-to-right. | `HStack().Gap("12")` |
| `Center()` | Centers its children both horizontally and vertically. | `Center().Children(Text("Centered"))` |
| `Text(content)` | Renders text with automatic sizing. | `Text("Hello").FontSize(16).TextColor(theme.Gray800)` |
| `Image()` | Container for rendering images. | `Image().Width("48").Height("48")` |
| `Canvas(draw)` | Custom drawing surface powered by Ebitengine. | `Canvas(func(ctx CanvasContext) { ... })` |

### Layout

Containers and spacing utilities for organizing your interface.

| Component | Description |
|-----------|-------------|
| `Container()` | Constrains content width with a max-width (1280px). |
| `Grid()` | Wrapping grid layout for arranging items in rows and columns. |
| `GridItem()` | An item inside a Grid that grows to fill available space. |
| `SimpleGrid()` | A grid with pre-configured gap spacing. |
| `Wrap()` | Horizontal row that wraps children to the next line. |
| `Spacer()` | Flexible space that grows to fill available room. |
| `Separator()` | Horizontal divider line. |
| `AspectRatio(ratio)` | Constrains a child to a specific aspect ratio. |

### Typography

Text and content components for displaying written information.

| Component | Description |
|-----------|-------------|
| `Heading(content, level)` | Semantic heading (level 1-6, descending size). |
| `Text(content)` | Paragraph text. |
| `Link(content)` | Blue underlined text for navigation. |
| `Code(content)` | Inline code with monospace styling. |
| `Kbd(content)` | Keyboard key display with raised shadow. |
| `Mark(content)` | Highlighted text with yellow background. |
| `Blockquote(lines...)` | Quoted content with left accent border. |
| `BlockquoteCaption(cite)` | Author citation below a blockquote. |
| `List(items, ordered)` | Bulleted or numbered list of items. |
| `ListItem(content)` | A single list item. |

### Buttons

| Component | Description |
|-----------|-------------|
| `Button(label)` | Standard button with blue background and hover effect. |
| `OutlineButton(label)` | Outlined button with transparent background. |
| `CloseButton()` | Small square button with an × icon. |

### Forms

Input controls for collecting user data.

| Component | Description |
|-----------|-------------|
| `Input(placeholder)` | Single-line text input field. |
| `TextArea(placeholder)` | Multi-line text input area. |
| `NumberInput()` | Numeric input with stepper buttons. |
| `Checkbox(label, checked)` | Toggle checkbox with label. |
| `Switch(label, checked)` | Toggle switch with label. |
| `Radio(label)` | Radio button with circular indicator. |
| `RadioGroup()` | Container for grouping related radio buttons. |
| `SegmentedControl()` | Horizontal control for toggling between segments. |
| `Segment(label)` | A single segment inside a SegmentedControl. |

### Overlays

Floating components that appear above other content.

| Component | Description |
|-----------|-------------|
| `Dialog(title)` | Modal dialog with title, body, and footer. |
| `DialogRoot()` | Dialog builder with Title, Body, Footer, CloseButton, and Trigger methods. |
| `Menu()` | Dropdown menu container. |
| `MenuItem(label)` | Clickable menu item with hover highlight. |
| `Tooltip(label)` | Hover tooltip with dark background. |
| `WithTooltip(target, tooltip)` | Attaches a tooltip to an element to show on hover. |
| `Popper()` | Positioned popup container with shadow. |
| `PopoverAt(target, pop, placement)` | Click-triggered popover attached to an element. |
| `Portal()` | Absolutely-positioned overlay layer. |

### Disclosure

Components for showing and hiding content sections.

| Component | Description |
|-----------|-------------|
| `Accordion()` | Vertical container for collapsible items. |
| `AccordionItem()` | A single collapsible section with border. |
| `AccordionHeader(label)` | Clickable header that expands the section. |
| `AccordionPanel()` | Expandable content panel (hidden by default). |
| `Breadcrumb()` | Horizontal breadcrumb trail container. |
| `BreadcrumbItem(label)` | A breadcrumb step. |
| `BreadcrumbLink(label)` | Clickable breadcrumb link (uses `Link`). |
| `BreadcrumbSeparator()` | Slash separator between breadcrumb items. |
| `Tabs()` | Vertical container for a tabbed interface. |
| `TabList()` | Horizontal row of tab triggers. |
| `Tab(label)` | Clickable tab trigger with hover effect. |
| `TabPanels()` | Container for tab content panels. |
| `TabPanel()` | Content panel for a tab (hidden by default). |
| `Collapsible()` | Collapsible section container. |
| `CollapsibleTrigger(label)` | Clickable trigger to toggle collapse. |
| `CollapsibleContent()` | Collapsible content panel. |
| `Pagination()` | Horizontal container for page navigation. |
| `PageButton(label)` | Pagination button with hover effect. |
| `Steps()` | Horizontal steps indicator container. |
| `Step(active)` | A step in a steps indicator. |
| `StepIndicator(number, active)` | Numbered circle indicator for a step. |
| `StepSeparator()` | Horizontal line between step indicators. |

### Feedback

Status indicators, loading states, and notifications.

| Component | Description |
|-----------|-------------|
| `Alert(status)` | Status message with left accent border (info/success/warning/error). |
| `AlertTitle(title)` | Bold title for an alert. |
| `AlertDescription(desc)` | Muted description text for an alert. |
| `EmptyState()` | Centered container for empty state content. |
| `EmptyStateTitle(title)` | Title for an empty state. |
| `EmptyStateDescription(desc)` | Description for an empty state. |
| `Progress(value)` | Linear progress bar (0-100). |
| `ProgressCircle(value)` | Circular progress indicator (0-100). |
| `Skeleton()` | Loading placeholder box. |
| `Spinner()` | Animated loading spinner. |
| `Status(variant)` | Small colored dot indicator (success/warning/error/info). |

### Data Display

Components for presenting structured information.

| Component | Description |
|-----------|-------------|
| `Avatar(name)` | Circular avatar with initials derived from a name. |
| `Badge(label)` | Small label for status or categorization (subtle variant). |
| `BadgeSolid(label, color)` | Solid background badge with color scheme. |
| `BadgeSubtle(label, color)` | Subtle background badge. |
| `BadgeOutline(label, color)` | Outlined badge with colored border. |
| `Card()` | Content container with white background, border, and shadow. |
| `CardHeader()` | Top section of a card. |
| `CardBody()` | Main content area of a card. |
| `CardFooter()` | Bottom action area of a card (horizontal layout). |
| `DataList()` | Vertical container for key-value data rows. |
| `DataListItem()` | A single row in a data list. |
| `DataListLabel(text)` | Label column in a data list row. |
| `DataListValue(text)` | Value column in a data list row. |
| `Stat()` | Vertical container for a statistic. |
| `StatLabel(text)` | Muted label above a statistic value. |
| `StatNumber(text)` | Large bold number for a statistic. |
| `StatHelpText(text)` | Small helper text below a statistic. |
| `Table()` | Data table container with border and rounded corners. |
| `TableHead()` | Table header section with background. |
| `TableBody()` | Table body section. |
| `TableRow()` | Horizontal table row with bottom border. |
| `TableCell(label)` | Data cell with label. |
| `TableHeaderCell(label)` | Header cell with muted styling. |
| `Tag()` | Horizontal tag container for categorization. |
| `TagLabel(label)` | Text label inside a tag. |
| `TagCloseButton(onClose)` | Small close button inside a tag. |
| `Timeline()` | Vertical timeline container. |
| `TimelineItem()` | Horizontal timeline row with dot and content. |
| `TimelineDot()` | Circular indicator dot on the timeline. |
| `TimelineContent()` | Content panel with left connector line. |

### Navigation

Page-level and application shell components.

| Component | Description |
|-----------|-------------|
| `Drawer()` | Side panel with absolute positioning and shadow. |
| `Window()` | Application window builder with title, size, and theme. |
| `App()` | Application root container. |

### Utilities

Helper components for controlling visibility and behavior.

| Component | Description |
|-----------|-------------|
| `Show(condition)` | Conditionally shows or hides content. |
| `VisuallyHidden()` | Makes content invisible but keeps it in layout. |
| `Icon(name)` | Small centered icon placeholder. |
| `Float()` | Absolutely-positioned floating container. |

---

## Styling

Every component supports chained builder methods for styling:

```go
raitui.Box().
    Width("200").Height("48").           // dimensions
    Padding("12").PaddingX("16").        // spacing
    BackgroundColor(theme.Blue500).      // fill color
    BorderColor(theme.Blue600).          // border color
    BorderWidth("2").                    // border size
    BorderRadius(8).                     // corner rounding
    BoxShadow(0, 2, 8, 0, theme.Black). // shadow
    Opacity(0.9)                         // transparency
```

### Colors

The `theme` package includes a full color palette:

```go
theme.Gray50  ... theme.Gray900    // 9 gray shades
theme.Blue50  ... theme.Blue900    // 9 blue shades
theme.Red50   ... theme.Red900     // 9 red shades
theme.Green50 ... theme.Green900   // 9 green shades
// ... and 9 more color families: Orange, Yellow, Teal, Cyan, Purple, Pink
```

### Layout Constants

Use the `props` package for flexbox properties:

```go
import "raitui/props"

props.FlexDirectionRow      props.FlexDirectionColumn
props.JustifyCenter          props.JustifySpaceBetween
props.AlignCenter            props.AlignFlexStart
props.WrapWrap               props.PositionAbsolute
```

---

## Events

Handle user interaction with click and hover callbacks:

```go
btn := raitui.Button("Click Me")

btn.OnClick(func() {
    fmt.Println("Button clicked!")
})

btn.OnHover(func(entered bool) {
    if entered {
        btn.BackgroundColor(theme.Blue600)
    } else {
        btn.BackgroundColor(theme.Blue500)
    }
})
```

---

## State Management

RaitUI includes a reactive state system for managing application data:

```go
import "raitui"

// Global state
raitui.SetState("count", 0)
count := raitui.StateInt("count", 0)
setCount := raitui.Setter[int]("count")

// Bound text
counter := raitui.Text("0").Use(func() string {
    return fmt.Sprintf("%d", raitui.StateInt("count", 0)())
})
```

---

## Canvas

The `Canvas` primitive gives you direct access to Ebitengine's drawing API for custom rendering:

```go
import (
    "image/color"
    "math"

    "github.com/hajimehoshi/ebiten/v2/vector"
    "raitui"
)

spinner := raitui.Canvas(func(ctx raitui.CanvasContext) {
    cx := ctx.W / 2
    cy := ctx.H / 2
    r := float32(10)

    // Animated rotating arc
    angle := float64(ctx.Frame) * 0.15
    x := cx + float32(math.Cos(angle))*r
    y := cy + float32(math.Sin(angle))*r

    clr := color.NRGBA{R: 66, G: 133, B: 244, A: 255}
    vector.DrawFilledCircle(ctx.Screen, x, y, 4, clr, true)
}).Width("48").Height("48")
```

`CanvasContext` provides:
- `Screen` — the `*ebiten.Image` drawing surface
- `X`, `Y` — top-left position within the parent
- `W`, `H` — available width and height
- `Frame` — incremented each frame for animations

---

## License

MIT
