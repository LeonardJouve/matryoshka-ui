package components

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
)

// Row lays children out horizontally with a uniform gap.
func Row(gap uint16, children ...Element) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)),
			Gap(GapHorizontal(gap)),
		),
		Children(children...),
	)
}

// Column lays children out vertically with a uniform gap.
func Column(gap uint16, children ...Element) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_VERTICAL),
			Width(Grow(1)),
			Gap(GapVertical(gap)),
		),
		Children(children...),
	)
}

// Center places a single child centered on both axes inside a growing box.
func Center(child Element) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Grow(1)),
			Height(Grow(1)),
			Justify(JustifyCenter),
			Align(AlignCenter),
		),
		Children(child),
	)
}

// Spacer is a flexible gap that pushes siblings apart on the main axis.
func Spacer() Element {
	return Div(Style(Width(Grow(1)), Height(Grow(1))))
}

// FixedSpacer reserves an exact amount of space (square).
func FixedSpacer(size uint16) Element {
	return Div(Style(Width(Fixed(size)), Height(Fixed(size))))
}

// HDivider is a thin horizontal rule spanning the available width.
func HDivider() Element {
	return Div(Style(Width(Grow(1)), Height(Fixed(1)), Color(T.Border)))
}

// VDivider is a thin vertical rule spanning the available height.
func VDivider() Element {
	return Div(Style(Width(Fixed(1)), Height(Grow(1)), Color(T.Border)))
}

// Padded wraps a child in a uniform padding box.
func Padded(p uint16, child Element) Element {
	return Div(
		Style(Width(Grow(1)), Padding(PaddingAll(p))),
		Children(child),
	)
}

// Box is a bare colored container with rounded corners — a building block.
func Box(c utils.Color, child Element) Element {
	return Div(
		Style(
			Width(Grow(1)),
			Color(c),
			BorderRadius(0.1),
		),
		Children(child),
	)
}
