package dsl

import "github.com/LeonardJouve/matryoshka-ui/utils"

type LayoutAxisT uint16

const (
	LAYOUT_HORIZONTAL LayoutAxisT = iota
	LAYOUT_VERTICAL
)

type StyleS struct {
	LayoutAxis LayoutAxisT
	Color      utils.Color
	Padding    *PaddingS
	Gap        *GapS
	Width      LayoutSize
	Height     LayoutSize
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
		LayoutAxis: LAYOUT_HORIZONTAL,
		Color: utils.Color{
			Red:   0,
			Green: 0,
			Blue:  0,
		},
		Padding: &PaddingS{},
		Gap:     &GapS{},
		Width:   Fit(),
		Height:  Fit(),
	}
}

func LayoutAxis(layoutAxis LayoutAxisT) StyleModifier {
	return func(style *StyleS) {
		style.LayoutAxis = layoutAxis
	}
}

func Padding(modifiers ...PaddingModifier) StyleModifier {
	return func(style *StyleS) {
		for _, modifier := range modifiers {
			modifier(style.Padding)
		}
	}
}

func Gap(modifiers ...GapModifier) StyleModifier {
	return func(style *StyleS) {
		for _, modifier := range modifiers {
			modifier(style.Gap)
		}
	}
}

func Color(color utils.Color) StyleModifier {
	return func(style *StyleS) {
		style.Color = color
	}
}

func Width(width LayoutSize) StyleModifier {
	return func(style *StyleS) {
		style.Width = width
	}
}

func Height(height LayoutSize) StyleModifier {
	return func(style *StyleS) {
		style.Height = height
	}
}

func (s *StyleS) layoutAxisSize() LayoutSize {
	if s.LayoutAxis == LAYOUT_HORIZONTAL {
		return s.Width
	}

	return s.Height
}

func (s *StyleS) crossAxisSize() LayoutSize {
	if s.LayoutAxis == LAYOUT_HORIZONTAL {
		return s.Height
	}

	return s.Width
}

func (s *StyleS) axisSize(layout LayoutAxisT) LayoutSize {
	if layout == LAYOUT_HORIZONTAL {
		return s.Width
	}

	return s.Height
}

func (s *StyleS) axisGap(l LayoutAxisT) uint16 {
	if l == LAYOUT_HORIZONTAL {
		return s.Gap.horizontal
	}

	return s.Gap.vertical
}

func (s *StyleS) layoutAxisGap() uint16 {
	return s.axisGap(s.LayoutAxis)
}

func (s *StyleS) crossAxisGap() uint16 {
	return s.axisGap(s.oppositeAxis())
}

type axisPaddingS struct {
	start uint16
	end   uint16
}

func (s *StyleS) axisPadding(l LayoutAxisT) axisPaddingS {
	if l == LAYOUT_HORIZONTAL {
		return axisPaddingS{
			start: s.Padding.left,
			end:   s.Padding.right,
		}
	}

	return axisPaddingS{
		start: s.Padding.top,
		end:   s.Padding.bottom,
	}
}

func (s *StyleS) layoutAxisPadding() axisPaddingS {
	return s.axisPadding(s.LayoutAxis)
}

func (s *StyleS) crossAxisPadding() axisPaddingS {
	return s.axisPadding(s.oppositeAxis())
}

func (s *StyleS) oppositeAxis() LayoutAxisT {
	if s.LayoutAxis == LAYOUT_HORIZONTAL {
		return LAYOUT_VERTICAL
	}

	return LAYOUT_HORIZONTAL
}
