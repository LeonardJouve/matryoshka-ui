package dsl

import "github.com/LeonardJouve/matryoshka-ui/utils"

type DivS struct {
	StyleS *StyleS
	Kids   []Element
}

type DivAttrs struct {
	Children []*Node
}

type DivOpt func(*DivS)
type DivStyleModifier interface {
	StyleModifier
	div()
}
type DivOnly func(*StyleS)

func (f DivOnly) applyStyle(s *StyleS) { f(s) }
func (f DivOnly) div()                 {}

func Div(opts ...DivOpt) *DivS {
	d := DivS{
		StyleS: NewStyle(),
	}
	for _, opt := range opts {
		opt(&d)
	}
	return &d
}

func (d *DivS) build() *Node {
	n := &Node{
		Kind:     KindDiv,
		Style:    d.StyleS,
		Layout:   NewLayout(),
		DivAttrs: DivAttrs{},
	}

	for _, k := range d.Kids {
		n.Children = append(n.Children, k.build())
	}

	return n
}

func Children(els ...Element) DivOpt {
	return func(d *DivS) {
		d.Kids = append(d.Kids, els...)
	}
}

func Style(modifiers ...DivStyleModifier) DivOpt {
	return func(d *DivS) {
		for _, m := range modifiers {
			m.applyStyle(d.StyleS)
		}
	}
}

func BorderRadius(r float32) DivOnly {
	return func(d *StyleS) {
		r = max(0, r)
		r = min(1, r)
		d.BorderRadius = r
	}
}

func BorderWidth(width uint16) DivOnly {
	return func(d *StyleS) {
		d.BorderWidth = width
	}
}

func BorderColor(c utils.Color) DivOnly {
	return func(d *StyleS) {
		d.BorderColor = c
	}
}

func LayoutAxis(layoutAxis LayoutAxisT) DivOnly {
	return func(style *StyleS) {
		style.LayoutAxis = layoutAxis
	}
}

func Padding(modifiers ...PaddingModifier) DivOnly {
	return func(style *StyleS) {
		for _, modifier := range modifiers {
			modifier(style.Padding)
		}
	}
}

func Gap(modifiers ...GapModifier) DivOnly {
	return func(style *StyleS) {
		for _, modifier := range modifiers {
			modifier(style.Gap)
		}
	}
}

func Justify(j JustifyT) DivOnly {
	return func(style *StyleS) {
		style.Justify = j
	}
}

func Align(align AlignT) DivOnly {
	return func(style *StyleS) {
		style.Align = align
	}
}
