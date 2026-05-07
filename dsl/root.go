package dsl

type RootS struct {
	*Element
}

func Root(element *Element) *RootS {
	size(element)
	position(element, nil, 0, 0)

	return &RootS{
		Element: element,
	}
}

func position(element *Element, parent *Element, horizontalOffset uint16, verticalOffset uint16) {
	// DFS preordre
	// TODO handle overflows

	if parent == nil {
		element.layout.X = 0
		element.layout.Y = 0
	} else {
		element.layout.X = horizontalOffset
		element.layout.Y = verticalOffset
	}

	horizontalOffset += element.style.padding.left
	verticalOffset += element.style.padding.top

	var maxWidthChild uint16 = 0
	var maxHeightChild uint16 = 0

	for _, child := range element.Children() {

		if element.style.layoutAxis == LAYOUT_HORIZONTAL {
			if horizontalOffset+child.Width() >= element.X()+element.Width() {
				horizontalOffset = element.X() + element.style.padding.left
				verticalOffset += element.style.gap.vertical + maxHeightChild
			}
		} else {
			if verticalOffset+child.Height() >= element.Y()+element.Height() {
				verticalOffset = element.Y() + element.style.padding.top
				horizontalOffset += element.style.gap.horizontal + maxWidthChild
			}
		}

		position(child, element, horizontalOffset, verticalOffset)
		maxWidthChild = max(maxWidthChild, child.Width())
		maxHeightChild = max(maxHeightChild, child.Height())

		if element.style.layoutAxis == LAYOUT_HORIZONTAL {
			horizontalOffset += child.layout.Width
			horizontalOffset += element.style.gap.horizontal
		} else {
			verticalOffset += child.layout.Height
			verticalOffset += element.style.gap.vertical
		}
	}
}

func size(el *Element) {
	//DFS postordre
	for _, child := range el.Children() {
		size(child)
	}

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
