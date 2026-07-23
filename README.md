# RaitUI

Go GUI framework using [Ebitengine](https://github.com/hajimehoshi/ebiten) with a builder-pattern API inspired by [Chakra-UI](https://chakra-ui.com/).

## Architecture

```
goda (layout engine)
   ↑
Primitive Components (Box, Flex, Stack, VStack, HStack, Image, Canvas, Text)
   ↑
Composite Components (Button, Card, Dialog, Menu, Tabs, etc.)
   ↑
Application Code
```

Only **primitive components** interact with `goda`. **Composite components** are built purely by composing primitives.

## Quick Start

```go
package main

import (
    "raitui"
    "raitui/theme"
)

func main() {
    app := raitui.VStack().
        Width("100%").Height("100%").
        Padding("24").Gap("16").
        BackgroundColor(theme.Gray50).
        Children(
            raitui.Heading("Welcome", 1),
            raitui.Button("Get Started"),
        )

    raitui.Window().
        Title("My App").
        Width(800).Height(600).
        MinSize(400, 300).
        Children(app).
        Run()
}
```

## Primitives

| Primitive | Description |
|-----------|-------------|
| `Box()` | Fundamental layout node, the root of all components |
| `Flex(dir)` | Flexbox container with configurable direction |
| `Stack(dir)` | Flex with predefined defaults (basis for VStack/HStack) |
| `VStack()` | Vertical stack = `Stack(FlexDirectionColumn)` |
| `HStack()` | Horizontal stack = `Stack(FlexDirectionRow)` |
| `Center()` | Flex with `justifyContent: center` + `alignItems: center` |
| `Text(content)` | Text rendering with auto-sizing |
| `Image()` | Image-capable flex container |
| `Canvas(draw)` | Custom Ebitengine drawing surface |

### Canvas

```go
raitui.Canvas(func(ctx raitui.CanvasContext) {
    // ctx.Screen  - *ebiten.Image
    // ctx.X, ctx.Y - top-left position
    // ctx.W, ctx.H - dimensions
    // ctx.Frame   - animation frame counter
    vector.DrawFilledCircle(ctx.Screen, ctx.W/2, ctx.H/2, 50, clr, true)
}).Width("100").Height("100")
```

## Composite Components

### Layout
`Container`, `Separator`, `Spacer`, `Grid`, `GridItem`, `SimpleGrid`, `Wrap`, `AspectRatio`

### Typography
`Heading`, `Code`, `Link`, `Blockquote`, `List`, `ListItem`, `Kbd`, `Mark`

### Buttons
`Button`, `OutlineButton`, `CloseButton`, `IconButton`

### Forms
`Checkbox`, `Switch`, `Input`, `TextArea`, `NumberInput`, `RadioGroup`, `Radio`, `SegmentedControl`, `Segment`

### Overlays
`Dialog`, `DialogRoot`, `Menu`, `MenuItem`, `Tooltip`, `Popper`, `Popover`, `Portal`

### Disclosure
`Accordion`, `AccordionItem`, `AccordionHeader`, `AccordionPanel`, `Breadcrumb`, `BreadcrumbItem`, `BreadcrumbLink`, `BreadcrumbSeparator`, `Tabs`, `TabList`, `Tab`, `TabPanels`, `TabPanel`, `Collapsible`, `Pagination`, `Steps`

### Feedback
`Alert`, `AlertTitle`, `AlertDescription`, `EmptyState`, `Progress`, `ProgressCircle`, `Skeleton`, `Spinner`, `Status`

### Data Display
`Avatar`, `Badge`, `BadgeSolid`, `BadgeSubtle`, `BadgeOutline`, `Card`, `CardHeader`, `CardBody`, `CardFooter`, `DataList`, `Stat`, `StatLabel`, `StatNumber`, `StatHelpText`, `Table`, `Tag`, `Timeline`

### Navigation
`Drawer`, `PageButton`, `Step`, `StepIndicator`, `StepSeparator`

### Utilities
`Show`, `VisuallyHidden`, `Icon`

## Props Reference

Every element supports chained builder methods:

**Dimensions:** `Width(v)`, `Height(v)`, `MinWidth(v)`, `MaxWidth(v)`, `MinHeight(v)`, `MaxHeight(v)`
**Spacing:** `Padding(v)`, `PaddingTop(v)`, `PaddingX(v)`, `Margin(v)`, `Gap(v)`
**Flex:** `FlexDirection(d)`, `JustifyContent(j)`, `AlignItems(a)`, `FlexGrow(v)`, `FlexShrink(v)`, `FlexWrap(w)`, `FlexBasis(v)`
**Visual:** `BackgroundColor(c)`, `BorderColor(c)`, `TextColor(c)`, `BorderRadius(r)`, `BorderWidth(v)`, `BoxShadow(...)`, `Opacity(v)`, `FontSize(s)`
**Events:** `OnClick(func())`, `OnHover(func(entered bool))`

## Layout Constants

```go
import "raitui/props"

props.JustifyCenter       props.AlignCenter        props.FlexDirectionRow
props.WrapWrap            props.PositionAbsolute    props.FlexDirectionColumn
```

## State Management

```go
raitui.SetState("count", 0)
count := raitui.StateInt("count", 0)
setCount := raitui.Setter[int]("count")
```

## License

MIT
