# RaitUI

Go GUI framework using [Ebitengine](https://github.com/hajimehoshi/ebiten) with a builder-pattern API inspired by [Chakra-UI](https://chakra-ui.com/).

```go
package main

import (
    "raitui"
    "raitui/theme"
)

func main() {
    app := raitui.App().
        Padding("24").Gap("16").
        BackgroundColor(theme.Gray50).
        Children(
            raitui.Text("Hello RaitUI!").TextColor(theme.Gray800).FontSize(18),
            raitui.Button("Click Me"),
        )

    raitui.Window().
        Title("My App").
        Width(800).Height(600).
        MinSize(400, 300).
        Children(app).
        Run()
}
```

## Quick Start

```bash
go install raitui/cmd/raitui@latest    # install CLI tool
raitui init                             # generate aliases.go
```

```go
package main

import (
    _ "raitui"    // or use aliases from generated aliases.go
    "raitui/theme"
)

func main() {
    app := App().
        Padding("20").Gap("12").
        BackgroundColor(theme.Gray50).
        Children(
            Heading("Welcome", 1),
            Button("Get Started"),
        )

    Window().
        Title("My App").
        Width(800).Height(600).
        MinSize(400, 300).
        Children(app).
        Run()
}
```
- **Chakra-UI component system** — Box, VStack, HStack, Text, Flex, Center, Container, Heading, Button, Input, and more
- **State management** — global + local state via sync.Map
- **Full color palette** — 13 colors x 10 shades
- **Keyboard input** — focus management, key repeat, cursor, shortcuts
- **Mouse events** — OnClick, OnHover with hit-testing
- **6 example apps** demonstrating components

## Quick Start

```
go get raitui
```

```go
package main

import (
    "raitui"
    "raitui/core"
    "raitui/theme"
)

func main() {
    root := raitui.VStack().
        Width("100%").Height("100%").
        Padding("20").Gap("12").
        BackgroundColor(theme.Gray50).
        Children(
            raitui.Heading("Welcome", 1),
            raitui.Button("Get Started"),
        )

    ctx := core.NewContext(theme.Gray50)
    ctx.SetMinWindowSize(400, 300)
    ctx.Run(root, "My App", 800, 600)
}
```

## Examples

```bash
go run ./examples/box/        # Basic layout with cards
go run ./examples/advanced/   # Dashboard with state, events, sidebar
go run ./examples/layout/     # Flex, Center, Container, Separator demos
go run ./examples/text/       # Heading, Code, Link, Blockquote, List
go run ./examples/form/       # Button, Checkbox, Switch, Input, TextArea
go run ./examples/minimal/    # Minimal hello-world
```

## Components

### Layout
`Box`, `VStack`, `HStack`, `Flex`, `Center`, `Container`, `Separator`, `Spacer`

### Typography
`Text`, `Heading`, `Code`, `Link`, `Blockquote`, `List`, `Kbd`, `Mark`

### Form
`Button`, `OutlineButton`, `Checkbox`, `Switch`, `Input`, `TextArea`

## Props Reference

Every element supports these builder methods:

**Dimensions**: `Width(v)`, `Height(v)`, `MinWidth(v)`, `MaxWidth(v)`, `MinHeight(v)`, `MaxHeight(v)`

**Spacing**: `Padding(v)`, `PaddingTop(v)`, `PaddingX(v)`, `Margin(v)`, `MarginTop(v)`, `Gap(v)`

**Flex**: `FlexDirection(d)`, `JustifyContent(j)`, `AlignItems(a)`, `FlexGrow(v)`, `FlexShrink(v)`, `FlexWrap(w)`, `FlexBasis(v)`

**Visual**: `BackgroundColor(c)`, `BorderColor(c)`, `TextColor(c)`, `BorderRadius(r)`, `BorderWidth(v)`, `BoxShadow(...)`, `Opacity(v)`, `FontSize(s)`

**HTML**: `ID(id)`, `Class(cls...)`, `Hidden(bool)`, `Title(t)`, `Lang(l)`, `Dir(d)`, `TabIndex(i)`

**Events**: `OnClick(func())`, `OnHover(func(entered bool))`

## Layout Constants

```go
import "raitui/props"

props.JustifySpaceBetween  props.AlignCenter        props.FlexDirectionRow
props.WrapWrap              props.PositionAbsolute    props.DisplayFlex
```

## State Management

```go
// Global state
raitui.SetState("clicks", 0)
getClicks := raitui.StateInt("clicks", 0)
setClicks := raitui.Setter[int]("clicks")

// Local state (per component)
raitui.SetLocalState("myForm", "value", "")
val, _ := raitui.GetLocalState("myForm", "value")
```

## License

MIT
