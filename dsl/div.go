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

func Div(opts ...ElementModifier) *Element {
	el := Element{
		id:       <-idGenerator,
		children: []*Element{},
		layout:   NewLayout(),
		style:    NewStyle(),
	}
	for _, opt := range opts {
		opt(&el)
	}
	computeFit(&el)
	return &el
}

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

func computeFit(el *Element) {
	width := el.style.padding.left + el.style.padding.right
	height := el.style.padding.top + el.style.padding.bottom

	var maxCrossSize uint16 = 0
	for _, child := range el.children {
		if el.style.layoutAxis == LAYOUT_HORIZONTAL {
			width += child.Width()
			maxCrossSize = max(maxCrossSize, child.Height())
		} else {
			height += child.Height()
			maxCrossSize = max(maxCrossSize, child.Width())
		}
	}

	if el.style.layoutAxis == LAYOUT_HORIZONTAL {
		gaps := uint16(len(el.children)-1) * el.style.gap.horizontal
		width += gaps
		height += maxCrossSize
	} else {
		gaps := uint16(len(el.children)-1) * el.style.gap.vertical
		height += gaps
		width += maxCrossSize
	}

	if el.style.width != 0 {
		el.layout.Width = el.style.width
	} else {
		el.layout.Width = width
	}

	if el.style.height != 0 {
		el.layout.Height = el.style.height
	} else {
		el.layout.Height = height
	}
}

/*
Old algo
horizontalOffset := el.style.padding.left
	verticalOffset := el.style.padding.top
	var maxCrossSize uint16 = 0

	for _, child := range el.children {
		if el.style.layoutAxis == LAYOUT_HORIZONTAL {
			width := child.Width()
			maxCrossSize = max(maxCrossSize, child.Height())

			// if fixed width and content overflows, increment cross offset
			if el.style.width != 0 && horizontalOffset != 0 && horizontalOffset+width > el.style.width {
				verticalOffset += maxCrossSize + el.style.gap.vertical // max height
				horizontalOffset = el.style.padding.left
				maxCrossSize = 0
			}

			child.layout.X = horizontalOffset
			child.layout.Y = verticalOffset
			horizontalOffset += width + el.style.gap.horizontal
		} else {
			height := child.Height()
			maxCrossSize = max(maxCrossSize, child.Width())

			// if fixed width and content overflows, increment cross offset
			if el.style.height != 0 && verticalOffset != 0 && verticalOffset+height > el.style.height {
				horizontalOffset += maxCrossSize + el.style.gap.horizontal // max height
				verticalOffset = el.style.padding.top
				maxCrossSize = 0
			}

			child.layout.Y = verticalOffset
			child.layout.X = horizontalOffset
			verticalOffset += height + el.style.gap.vertical
		}
	}

	if el.style.width != 0 {
		el.layout.Width = el.style.width
	} else {
		el.layout.Width = horizontalOffset
	}

	if el.style.height != 0 {
		el.layout.Height = el.style.height
	} else {
		el.layout.Height = verticalOffset
	}
*/
