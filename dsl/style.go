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
	width      LayoutSize
	height     LayoutSize
}

type StyleModifier = func(style *StyleS)

type LayoutSize interface {
	isLayoutSize()
}

type fitS struct{}

func (f fitS) isLayoutSize() {}

func Fit() LayoutSize {
	return fitS{}
}

type growS struct {
	factor uint16
}

func (g growS) isLayoutSize() {}

func Grow(factor uint16) LayoutSize {
	return growS{
		factor: factor,
	}
}

type fixedS struct {
	Size uint16
}

func Fixed(size uint16) LayoutSize {
	return fixedS{
		Size: size,
	}
}

func (f fixedS) isLayoutSize() {}

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
		width:   Fit(),
		height:  Fit(),
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

func Width(width LayoutSize) StyleModifier {
	return func(style *StyleS) {
		style.width = width
	}
}

func Height(height LayoutSize) StyleModifier {
	return func(style *StyleS) {
		style.height = height
	}
}
