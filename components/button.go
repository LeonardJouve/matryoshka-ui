package components

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
)

// NOTE: the layout engine has no event/input layer, so these are
// purely visual — styled, padded labels. They convey affordance, not
// behavior. Wire interaction in at the renderer level if you need it.

// Button is a solid filled button with a centered label.
func Button(label string, fill utils.Color) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Color(fill),
			BorderRadius(0.3),
			Padding(Padding2(8, 16)),
			Justify(JustifyCenter),
			Align(AlignCenter),
		),
		Children(
			Text(label, TextStyle(FontSize(14), Color(T.Text))),
		),
	)
}

// PrimaryButton is a Button using the theme's primary color.
func PrimaryButton(label string) Element {
	return Button(label, T.Primary)
}

// OutlineButton is a transparent button with a colored border.
func OutlineButton(label string, c utils.Color) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Color(T.Surface),
			BorderRadius(0.3),
			BorderColor(c),
			BorderWidth(1),
			Padding(Padding2(8, 16)),
			Justify(JustifyCenter),
			Align(AlignCenter),
		),
		Children(
			Text(label, TextStyle(FontSize(14), Color(c))),
		),
	)
}

// Badge is a small rounded pill for statuses or counts.
func Badge(label string, c utils.Color) Element {
	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Color(c),
			BorderRadius(0.8),
			Padding(Padding2(3, 10)),
			Justify(JustifyCenter),
			Align(AlignCenter),
		),
		Children(
			Text(label, TextStyle(FontSize(11), Color(T.Bg))),
		),
	)
}

// Padding2 is a convenience for vertical + horizontal padding.
func Padding2(v uint16, h uint16) PaddingModifier {
	return func(p *PaddingS) {
		PaddingVertical(v)(p)
		PaddingHorizontal(h)(p)
	}
}
