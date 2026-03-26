package dsl

import "github.com/LeonardJouve/matryoshka-ui/utils"

type LayoutAxisT uint16

const (
	LAYOUT_HORIZONTAL LayoutAxisT = iota
	LAYOUT_VERTICAL
)

type StyleS struct {
	layoutAxis LayoutAxisT
	color      utils.Color
	padding    *PaddingS
	gap        *GapS
	width      uint16
	height     uint16
}

type StyleModifier = func(style *StyleS)

func NewStyle() *StyleS {
	return &StyleS{
		layoutAxis: LAYOUT_HORIZONTAL,
		color: utils.Color{
			Red:   0,
			Green: 0,
			Blue:  0,
		},
		padding: &PaddingS{},
		gap:     &GapS{},
		width:   0,
		height:  0,
	}
}

func LayoutAxis(layoutAxis LayoutAxisT) StyleModifier {
	return func(style *StyleS) {
		style.layoutAxis = layoutAxis
	}
}

func Padding(modifiers ...PaddingModifier) StyleModifier {
	return func(style *StyleS) {
		for _, modifier := range modifiers {
			modifier(style.padding)
		}
	}
}

func Gap(modifiers ...GapModifier) StyleModifier {
	return func(style *StyleS) {
		for _, modifier := range modifiers {
			modifier(style.gap)
		}
	}
}

func Color(color utils.Color) StyleModifier {
	return func(style *StyleS) {
		style.color = color
	}
}

func Width(width uint16) StyleModifier {
	return func(style *StyleS) {
		style.width = width
	}
}

func Height(height uint16) StyleModifier {
	return func(style *StyleS) {
		style.height = height
	}
}
