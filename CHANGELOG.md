# Changelog

## [Unreleased]

### Added
- Core element system with builder-pattern API (Box, VStack, HStack, Text)
- 40+ CSS-like props on Element (Width, Height, Padding, Margin, Flex*, Align*, etc.)
- HTML global attributes (ID, Class, Hidden, Title, Lang, Dir, TabIndex)
- Visual props (BackgroundColor, BorderColor, TextColor, BorderRadius, Shadow, Opacity)
- Mouse events (OnClick, OnHover) with hit-test dispatch
- Keyboard input handling with focus management, key repeat, cursor positioning
- State management via sync.Map (global + local state with component prefix)
- Ebitengine integration with goda (Yoga) flexbox layout engine
- Min window size support via `ebiten.SetWindowSizeLimits`

### Components
- **Layout**: Flex, Center, Container, Separator, Spacer
- **Typography**: Heading, Code, Link, Blockquote, List, Kbd, Mark
- **Form**: Button (solid + outline variants), Checkbox, Switch, Input, TextArea

### Theme
- Full Chakra-UI color palette (13 colors x 10 shades)
- Design tokens and color system

### Infrastructure
- Keyboard input package (`core/keyboard`) with key repeat, shortcuts, modifiers
- Layout test suite verifying goda position/size computation
- 6 example applications in `examples/`
- `props` package re-exporting goda layout constants (Justify, Align, Wrap, etc.)
