# Changelog

## [Unreleased]

### Architecture Refactoring

- **Primitive layer**: Only `Box`, `Flex`, `Stack`, `VStack`, `HStack`, `Image`, `Canvas`, `Text`, `Center` create goda nodes
- **Composite layer**: All other components (Button, Card, Dialog, Menu, Tabs, etc.) compose purely from primitives
- **Zero goda in composites**: Composite components no longer import `goda` or access `GNode` directly
- **`props` package**: Layout constants (Justify, Align, FlexDirection, Wrap, Position, Display, Overflow) re-exported for composite use
- **Nested composition**: Composites can compose from other composites (e.g., `BreadcrumbLink` uses `Link`, `Dialog` uses `Card`)

### Added

- **`Canvas` primitive**: Custom Ebitengine drawing surface via callback with `CanvasContext` (Screen, position, dimensions, frame counter)
- **`Stack` primitive**: Configurable-direction stack, base for VStack/HStack
- **`CardHeader`**, **`CardBody`**, **`CardFooter`**: Card sub-components matching Chakra UI structure
- **`AlertTitle`**, **`AlertDescription`**: Alert sub-components
- **`BlockquoteCaption`**: Blockquote author caption
- **`EmptyStateTitle`**, **`EmptyStateDescription`**: Empty state sub-components
- **`ListItem`**: List item component
- **`TypeCanvas` element type**: Canvas render dispatch

### Changed

- **VStack**: Now composes from `Stack(FlexDirectionColumn)` instead of creating own goda node
- **HStack**: Now composes from `Stack(FlexDirectionRow)` instead of creating own goda node
- **Center**: Now composes from `Flex` with justify/align presets
- **Box**: Uses `FlexShrink()` public API instead of raw `GNode`
- **Flex**: Uses `AlignItems()`/`FlexShrink()` public API
- **Text**: Uses `Width()`/`Height()` public API
- **Image**: Uses `FlexShrink()` public API
- **Button**: Composes from `Center()` + `Text()` child
- **Badge**: Composes from `Center()` + `Text()` child
- **Card**: Composes from `VStack()` with border/shadow
- **Alert**: Composes from `VStack()` with left accent border
- **Avatar**: Composes from `Center()` + `Text()` (circle via BorderRadius)
- **Blockquote**: Composes from `Box()` + `VStack()` > `Text()` children
- **Code/Kbd**: Compose from `Box()` + `Text()` child
- **Heading**: Composes from `Text()` with sized font
- **Link**: Composes from `Text()` with underline + blue color
- **Mark**: Composes from `Box()` with highlight bg + `Text()` child
- **CloseButton**: Composes from `Center()` + `Text("×")`
- **Tag**: Composes from `HStack()` / `Text()` / `Center()` (CloseButton)
- **Breadcrumb**: Composes from `HStack()` / `Link()` / `Text()`
- **Tabs**: Compose from `VStack()` / `HStack()` / `Center()` / `Text()`
- **Accordion**: Composes from `VStack()` / `HStack()` + hover
- **Table**: Composes from `VStack()` / `HStack()` / `Box()` + `Text()`
- **Timeline**: Composes from `VStack()` / `HStack()` / `Box()` + `props`
- **Stat**: Composes from `VStack()` + `Text()`
- **List**: Composes from `VStack()` + `Text()`
- **EmptyState**: Composes from `Center()` + `Text()`
- **Wrap**: Composes from `Flex()` + `FlexWrap(props.WrapWrap)`
- **Grid**: Composes from `Wrap()` / `Box()`
- **Icon**: Composes from `Center()`
- **Skeleton**: Composes from `Box()` with dimensions
- **Spacer**: Composes from `Box().FlexGrow(1)`
- **Show/VisuallyHidden**: Compose from `Box()`
- **Status**: Composes from `Box()` with circle styling
- **SegmentedControl**: Composes from `HStack()` + `Center()` segments
- **Radio**: Composes from `HStack()` + `Box()` dots + `Text()`
- **DataList**: Composes from `VStack()` / `HStack()` / `Text()`
- **Splitter/Float**: Compose from `Box()` + positioning
- **Container**: Composes from `Box()` with MaxWidth
- **Navigation** (Drawer, Collapsible, Pagination, Steps): Compose from primitives + props
- **NumberInput**: Composes from `HStack()` + `Center()` stepper buttons
- **Window/App/Portals**: Compose from `VStack()` / `Center()` + props
- **Dialog**: Composes from `Card()` + `VStack()` + `HStack()` + `Button()`
- **Tooltip/Popper/Popover**: Use props constants, public API
- **Menu**: Composes from `Box()` / `Text()` for MenuItem
- **Progress/Spinner/Switch/Checkbox/Input**: Use public API, keep type for rendering
- **Separator**: Composes from `Box()` with background color
- **AspectRatio**: Composes from `Box().AspectRatio(ratio)`
- **Portal/Drawer/Float/ScrollArea**: Compose from `Box()` + props

### Removed

- **examples/**: All example applications removed (to be recreated with new architecture)

### Technical

- `NewElement(TypeX)` eliminated in 40+ composite components
- All direct `GNode` accesses removed from composites
- All `goda` imports removed from 31 composite files
- `CanvasContext` struct with Screen, X, Y, W, H, Frame fields
- `SetDrawFunc`/`DrawFunc` methods on Element
- `renderCanvas` function in render pipeline
