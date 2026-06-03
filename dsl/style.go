package dsl

import "github.com/LeonardJouve/matryoshka-ui/utils"

type LayoutAxisT uint16

const (
	LAYOUT_HORIZONTAL LayoutAxisT = iota
	LAYOUT_VERTICAL
)

type JustifyT uint16

const (
	JustifyStart JustifyT = iota
	JustifyCenter
	JustifyEnd
	JustifyBetween
	JustifyAround
)

type AlignT uint16

const (
	AlignStart AlignT = iota
	AlignCenter
	AlignEnd
)

type WrapT int

const (
	NoWrap WrapT = iota // Default
	WrapYes
)

type StyleModifier interface{ applyStyle(*StyleS) }

type StyleS struct {
	LayoutAxis   LayoutAxisT
	Color        utils.Color
	Padding      *PaddingS
	Gap          *GapS
	Width        LayoutSize
	Height       LayoutSize
	FontSize     uint16
	Justify      JustifyT
	Align        AlignT
	Wrap         WrapT
	BorderRadius float32
	BorderWidth  uint16
	BorderColor  utils.Color
}

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
		Padding:  &PaddingS{},
		Gap:      &GapS{},
		Width:    Fit(),
		Height:   Fit(),
		FontSize: 16,
		Justify:  JustifyStart,
		Align:    AlignStart,
	}
}

type Shared func(*StyleS)

func (f Shared) applyStyle(s *StyleS) { f(s) }
func (f Shared) div()                 {}
func (f Shared) text()                {}
func (f Shared) image()               {}

func Color(color utils.Color) Shared {
	return func(style *StyleS) {
		style.Color = color
	}
}

func Width(width LayoutSize) Shared {
	return func(style *StyleS) {
		style.Width = width
	}
}

func Height(height LayoutSize) Shared {
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

func (s *StyleS) mainAxisPadding() axisPaddingS {
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
