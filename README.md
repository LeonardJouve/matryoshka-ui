# MatryoshkaUI

A small, declarative UI layout library for Go. You describe your interface as a tree of nested elements (hence *matryoshka*), the layout engine computes positions and sizes with a flexbox-like model, and a pluggable renderer draws the result. The default renderer ships on top of [raylib](https://www.raylib.com/).

---

## Requirements

- **Go 1.25+**
- **A C compiler** — the default renderer uses [`raylib-go`](https://github.com/gen2brain/raylib-go), which is a cgo binding to the native raylib library, so cgo (and therefore a C toolchain) is mandatory.
    - **Linux**: `gcc` plus the usual X11/OpenGL/audio dev headers
      ```sh
      sudo apt install build-essential libgl1-mesa-dev libx11-dev libxcursor-dev \
        libxrandr-dev libxinerama-dev libxi-dev libxxf86vm-dev libasound2-dev
      ```
    - **macOS**: the Xcode command line tools (`xcode-select --install`)
    - **Windows**: a MinGW-w64 `gcc` toolchain
- Make sure cgo is enabled (it is by default unless you've set `CGO_ENABLED=0`).

## Installation

```sh
git clone https://github.com/LeonardJouve/matryoshka-ui.git
cd matryoshka-ui
go mod download
```

Run the example UI:

```sh
go run .
```

This project also wires in [`air`](https://github.com/air-verse/air) as a Go tool dependency for live-reload during development:

```sh
go tool air
```

---

## Concepts

A UI is a tree of **elements**. There are three kinds:

- `Div` — a layout container that holds children
- `Text` — a string of text
- `Image` — an image loaded from a path

Every element carries a **style**. Containers lay their children out along a **main axis** (horizontal or vertical); the perpendicular direction is the **cross axis**. Sizing, spacing, alignment, and growth are all expressed relative to those axes, much like CSS flexbox.

You build the tree with `Div(...)`, `Text(...)`, `Image(...)`, wrap the top of it in `Root(...)`, and hand that to a renderer.

---

## Building a UI

### Containers — `Div`

`Div(opts ...DivOpt)` accepts these options:

- `Children(els ...Element)` — nested elements
- `Style(modifiers ...DivStyleModifier)` — style modifiers (see below)

### Text — `Text`

`Text(content string, opts ...TextOpt)`:

- `TextStyle(modifiers ...TextStyleModifier)` — style modifiers valid on text

### Image — `Image`

`Image(src string, opts ...ImageOpt)`:

- `ImageStyle(modifiers ...ImageStyleModifier)` — style modifiers valid on images

### Root

`Root(element Element) *Node` wraps your top-level element into the root node that the layout engine and renderer expect.

---

## Style reference

Modifiers are passed into `Style(...)`, `TextStyle(...)`, or `ImageStyle(...)`. Some are shared across all element kinds; others are specific.

### Shared (Div, Text, Image)

| Modifier | Description |
|---|---|
| `Color(utils.Color)` | Background/fill color (text color for `Text`) |
| `Width(LayoutSize)` | Width sizing (see *Sizing*) |
| `Height(LayoutSize)` | Height sizing (see *Sizing*) |

### Div-only

| Modifier | Description |
|---|---|
| `LayoutAxis(LayoutAxisT)` | `LAYOUT_HORIZONTAL` (default) or `LAYOUT_VERTICAL` |
| `Padding(...PaddingModifier)` | Inner padding |
| `Gap(...GapModifier)` | Spacing between children |
| `Justify(JustifyT)` | Main-axis distribution |
| `Align(AlignT)` | Cross-axis alignment |
| `BorderRadius(float32)` | Corner radius, clamped to `0..1` |
| `BorderWidth(uint16)` | Border thickness |
| `BorderColor(utils.Color)` | Border color |

### Text-only

| Modifier | Description |
|---|---|
| `FontSize(uint16)` | Font size (default `16`) |

### Sizing — `LayoutSize`

| Constructor | Meaning |
|---|---|
| `Fit()` | Size to content (default) |
| `Fixed(size uint16)` | Exact size in pixels |
| `Grow(factor uint16)` | Expand to fill leftover space, shared across siblings proportionally to `factor` |

### Padding — `PaddingModifier`

`PaddingAll(p)`, `PaddingVertical(v)`, `PaddingHorizontal(h)`, `PaddingTop(t)`, `PaddingBottom(b)`, `PaddingLeft(l)`, `PaddingRight(r)`.

### Gap — `GapModifier`

`GapAll(g)`, `GapVertical(v)`, `GapHorizontal(h)`.

### Justify — `JustifyT` (main axis)

`JustifyStart` (default), `JustifyCenter`, `JustifyEnd`, `JustifyBetween`, `JustifyAround`.

### Align — `AlignT` (cross axis)

`AlignStart` (default), `AlignCenter`, `AlignEnd`.

---

## Example

A music-player card: a vertical card containing album art, track info, a progress bar, and a control row.

```go
package main

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/renderer"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func col(r, g, b uint8) utils.Color { return utils.Color{Red: r, Green: g, Blue: b, Alpha: 255} }

func ui() *Node {
	var (
		bg     = col(16, 18, 27)
		card   = col(28, 31, 46)
		track  = col(45, 50, 75)
		white  = col(235, 238, 248)
		muted  = col(140, 146, 175)
		green  = col(0, 255, 0)
		border = col(42, 48, 68)
	)

	iconBtn := func(src string, size uint16) Element {
		return Image(src, ImageStyle(Width(Fixed(size)), Height(Fixed(size))))
	}

	return Root(Div(
		Style(
			Width(Grow(1)), Height(Grow(1)),
			LayoutAxis(LAYOUT_VERTICAL),
			Color(bg),
			Justify(JustifyCenter),
			Align(AlignCenter),
		),
		Children(
			Div(
				Style(
					LayoutAxis(LAYOUT_VERTICAL),
					Padding(PaddingAll(24)),
					Gap(GapAll(16)),
					Color(card),
					BorderRadius(0.1),
					BorderColor(border),
					BorderWidth(2),
				),
				Children(
					// album
					Div(
						Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Justify(JustifyCenter)),
						Children(
							Image("assets/album.png", ImageStyle(Width(Fixed(240)), Height(Fixed(240)))),
						),
					),

					// title / artist
					Text("Midnight City", TextStyle(FontSize(20), Color(white))),
					Text("M83", TextStyle(Color(muted), FontSize(14))),

					// player
					Div(
						Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Gap(GapHorizontal(10)), Justify(JustifyBetween), Align(AlignCenter)),
						Children(
							Text("1:24", TextStyle(Color(muted), FontSize(12))),
							Div(
								Style(Width(Grow(1)), Height(Grow(1)), Color(track), Padding(PaddingAll(4)), BorderRadius(0.8)),
								Children(Div(Style(Width(Fixed(60)), Height(Fixed(8)), Color(green), BorderRadius(0.8)))),
							),
							Text("4:03", TextStyle(Color(muted), FontSize(12))),
						),
					),

					// controls
					Div(
						Style(
							LayoutAxis(LAYOUT_HORIZONTAL),
							Width(Grow(1)),
							Justify(JustifyBetween),
						),
						Children(
							iconBtn("assets/shuffle.png", 24),
							iconBtn("assets/prev.png", 24),
							iconBtn("assets/play.png", 24),
							iconBtn("assets/next.png", 24),
							iconBtn("assets/repeat.png", 24),
						),
					),
				),
			),
		),
	))
}

func main() {
	renderer := renderer.NewRaylibRenderer()
	renderer.InitWindow(800, 800, "MusicPlayer")
	renderer.SetWindowFlag(rl.FlagWindowResizable | rl.FlagWindowHighdpi)
	defer renderer.CloseWindow()
	renderer.Render(ui())
}
```

---

## Bringing your own renderer

The layout engine is renderer-agnostic. It only needs to know how to measure text, and a renderer only needs to know how to draw the resulting node tree.

### 1. Implement `TextMeasurer`

Layout can't size `Text` nodes without knowing how wide a string is at a given font size. Provide this interface:

```go
type TextMeasurer interface {
	MeasureText(content string, fontSize uint16) (width uint16, height uint16)
}
```

### 2. Run the layout pass

For a given viewport size, call:

```go
root := dsl.Layout(node, width, height, measurer)
```

`Layout` fixes the root to the given dimensions and computes `Layout.X`, `Layout.Y`, `Layout.Width`, and `Layout.Height` on every node in the tree.

### 3. Walk the tree and draw

Each `*dsl.Node` exposes:

- `Kind` — `KindDiv`, `KindText`, or `KindImage`
- `Layout` — computed `X`, `Y`, `Width`, `Height`
- `Style` — colors, border, font size, etc.
- `DivAttrs.Children` — child nodes (Div)
- `TextAttrs.Content` — the string (Text)
- `ImageAttrs.Src` — the image path (Image)

Recurse over `Children`, and for each node draw according to its `Kind` and `Layout` rectangle. Conform to the `Renderer` interface so callers can swap renderers:

```go
type Renderer interface {
	InitWindow(width, height uint16, name string)
	CloseWindow()
	Render(element *dsl.Node)
}
```

A minimal render loop typically: clears the screen, calls `dsl.Layout` with the current window size (so the UI is responsive on resize), recursively draws the node tree, then repeats. The bundled `RaylibRenderer` is the reference implementation — useful to read alongside this section. A `TerminalRenderer` is also referenced as a commented-out alternative in `main.go`, showing the same UI tree can target a completely different backend.