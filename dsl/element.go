package dsl

import "github.com/LeonardJouve/matryoshka-ui/utils"

type Element struct {
	id       uint
	style    *StyleS
	children []*Element
	layout   Layout
}

type ElementModifier func(element *Element)

var idGenerator = utils.NewIDGenerator()

func Children(els ...*Element) ElementModifier {
	return func(el *Element) {
		for _, child := range els {
			el.children = append(el.children, child)
		}
	}
}

func Style(modifiers ...StyleModifier) ElementModifier {
	return func(el *Element) {
		for _, modifier := range modifiers {
			modifier(el.style)
		}
	}
}

func (el *Element) Width() uint16 {
	return el.layout.Width
}

func (el *Element) Height() uint16 {
	return el.layout.Height
}

func (el *Element) Children() []*Element {
	return el.children
}

func (el *Element) Color() utils.Color {
	return el.style.color
}

func (el *Element) X() uint16 {
	return el.layout.X
}

func (el *Element) Y() uint16 {
	return el.layout.Y
}
