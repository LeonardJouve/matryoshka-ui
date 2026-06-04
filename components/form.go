package components

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl"
)

// NOTE: no input layer in the engine — these render state you set at
// build time (filled text, on/off, checked). They show form structure,
// not live editing.

// Field is a labeled input row: a caption above a bordered value box.
// Pass the current value as text; empty value renders as a placeholder.
func Field(label string, value string, placeholder string) Element {
	shown := value
	textColor := T.Text
	if value == "" {
		shown = placeholder
		textColor = T.TextMute
	}
	return Div(
		Style(LayoutAxis(LAYOUT_VERTICAL), Width(Grow(1)), Gap(GapVertical(6))),
		Children(
			Label(label),
			Div(
				Style(
					LayoutAxis(LAYOUT_HORIZONTAL),
					Width(Grow(1)),
					Color(T.Bg),
					BorderRadius(0.15),
					BorderColor(T.Border),
					BorderWidth(1),
					Padding(Padding2(10, 12)),
					Align(AlignCenter),
				),
				Children(Text(shown, TextStyle(FontSize(13), Color(textColor)))),
			),
		),
	)
}

// Toggle is a pill switch. on=true slides the knob right and colors the track.
func Toggle(on bool) Element {
	track := T.Border
	knob := Row(0, Spacer(), knobDot())
	if on {
		track = T.Primary
		knob = Row(0, knobDot(), Spacer())
	}
	return Div(
		Style(
			LayoutAxis(LAYOUT_HORIZONTAL),
			Width(Fixed(44)),
			Height(Fixed(24)),
			Color(track),
			BorderRadius(1),
			Padding(PaddingAll(3)),
		),
		Children(knob),
	)
}

func knobDot() Element {
	return Div(Style(Width(Fixed(18)), Height(Fixed(18)), Color(T.Text), BorderRadius(1)))
}

// Checkbox is a square that fills with the primary color when checked.
func Checkbox(checked bool) Element {
	fill := T.Bg
	if checked {
		fill = T.Primary
	}
	return Div(Style(
		Width(Fixed(18)), Height(Fixed(18)),
		Color(fill),
		BorderRadius(0.2),
		BorderColor(T.Border),
		BorderWidth(1),
	))
}

// Radio is a circle, filled when selected.
func Radio(selected bool) Element {
	fill := T.Bg
	if selected {
		fill = T.Primary
	}
	return Div(Style(
		Width(Fixed(18)), Height(Fixed(18)),
		Color(fill),
		BorderRadius(1),
		BorderColor(T.Border),
		BorderWidth(1),
	))
}

// CheckRow pairs a checkbox with a label.
func CheckRow(label string, checked bool) Element {
	return Div(
		Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Gap(GapHorizontal(10)), Align(AlignCenter)),
		Children(Checkbox(checked), Body(label)),
	)
}

// ToggleRow pairs a label on the left with a toggle pushed to the right.
func ToggleRow(label string, on bool) Element {
	return Div(
		Style(LayoutAxis(LAYOUT_HORIZONTAL), Width(Grow(1)), Align(AlignCenter), Justify(JustifyBetween)),
		Children(Body(label), Toggle(on)),
	)
}
