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

func (s *StyleS) layoutAxisSize() LayoutSize {
	if s.layoutAxis == LAYOUT_HORIZONTAL {
		return s.width
	}

	return s.height
}

func (s *StyleS) crossAxisSize() LayoutSize {
	if s.layoutAxis == LAYOUT_HORIZONTAL {
		return s.height
	}

	return s.width
}

func (s *StyleS) layoutAxisGap() uint16 {
	if s.layoutAxis == LAYOUT_HORIZONTAL {
		return s.gap.horizontal
	}

	return s.gap.vertical
}

func (s *StyleS) crossAxisGap() uint16 {
	if s.layoutAxis == LAYOUT_HORIZONTAL {
		return s.gap.vertical
	}

	return s.gap.horizontal
}

type axisPadding struct {
	start uint16
	end   uint16
}

func (s *StyleS) layoutAxisPadding() axisPadding {
	if s.layoutAxis == LAYOUT_HORIZONTAL {
		return axisPadding{
			start: s.padding.left,
			end:   s.padding.right,
		}
	}

	return axisPadding{
		start: s.padding.top,
		end:   s.padding.bottom,
	}
}

func (s *StyleS) crossAxisPadding() axisPadding {
	if s.layoutAxis == LAYOUT_HORIZONTAL {
		return axisPadding{
			start: s.padding.top,
			end:   s.padding.bottom,
		}
	}

	return axisPadding{
		start: s.padding.left,
		end:   s.padding.right,
	}
}

func (s *StyleS) oppositeAxis() LayoutAxisT {
	if s.layoutAxis == LAYOUT_HORIZONTAL {
		return LAYOUT_VERTICAL
	}

	return LAYOUT_HORIZONTAL
}
