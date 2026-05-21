package dsl

type RootS struct {
	*Element
}

func Root(element *Element) *RootS {
	size(element)
	// TODO only grow layout axis
	grow(element)
	position(element, nil, 0, 0)
	// TODO grow cross axis

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

func grow(el *Element) {
	// DFS preordre

	// TODO grow on layout axis instead of width
	growWidthChildrenAmount := 0
	var usedWidth uint16 = 0
	var totalWidthGrowFactor uint32 = 0

	for _, child := range el.Children() {
		if g, ok := child.style.width.(growS); ok {
			growWidthChildrenAmount += 1
			totalWidthGrowFactor += uint32(g.factor)
		}
		usedWidth += child.Width()
	}

	gap := uint16(len(el.Children())-1) * el.style.gap.horizontal
	leftWidth := int32(el.Width()) - int32(usedWidth+el.style.padding.left+el.style.padding.right+gap)
	if leftWidth > 0 {
		for _, child := range el.Children() {
			if g, ok := child.style.width.(growS); ok {
				child.layout.Width += uint16(float64(leftWidth) * float64(g.factor) / float64(totalWidthGrowFactor))
			}
		}
	}

	for _, child := range el.Children() {
		grow(child)
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

	if layoutSize, ok := el.style.width.(fixedS); ok {
		el.layout.Width = layoutSize.Size
	} else {
		el.layout.Width = width
	}

	if layoutSize, ok := el.style.height.(fixedS); ok {
		el.layout.Height = layoutSize.Size
	} else {
		el.layout.Height = height
	}
}
