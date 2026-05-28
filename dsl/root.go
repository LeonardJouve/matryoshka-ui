package dsl

type RootS struct {
	*Element
}

func Root(element *Element) *RootS {
	size(element)
	overflow(element)
	// TODO only grow layout axis
	grow(element)
	position(element, nil, 0, 0)
	// TODO grow cross axis

	return &RootS{
		Element: element,
	}
}

func overflow(element *Element) {
	// DFS postordre
	for _, child := range element.children {
		overflow(child)
	}

	var width = element.style.padding.left
	var height = element.style.padding.top
	var line uint16 = 0

	for _, child := range element.Children() {
		if element.style.layoutAxis == LAYOUT_HORIZONTAL {
			if width+child.layout.Width > element.layout.Width {
				line += 1
				width = element.style.padding.left + child.layout.Width
			} else {
				width += child.layout.Width + element.style.gap.horizontal
			}
			child.layout.Line = line
		} else {
			if height+child.Height() > element.Height() {
				line += 1
				height = element.style.padding.top + child.layout.Height
			} else {
				height += child.layout.Height + element.style.gap.vertical
			}
			child.layout.Line = line
		}
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

	var line uint16 = 0
	for _, child := range element.Children() {
		if element.style.layoutAxis == LAYOUT_HORIZONTAL {
			if line != child.layout.Line {
				line += 1
				verticalOffset += maxHeightChild + element.style.gap.vertical
				horizontalOffset = element.X() + element.style.padding.left
				maxHeightChild = 0
			}
			maxHeightChild = max(maxHeightChild, child.layout.Height)
		} else {
			if line != child.layout.Line {
				line += 1
				horizontalOffset += maxWidthChild + element.style.gap.horizontal
				verticalOffset = element.Y() + element.style.padding.top
				maxWidthChild = 0
			}
			maxWidthChild = max(maxWidthChild, child.layout.Width)
		}

		position(child, element, horizontalOffset, verticalOffset)

		if element.style.layoutAxis == LAYOUT_HORIZONTAL {
			horizontalOffset += child.layout.Width
			horizontalOffset += element.style.gap.horizontal
		} else {
			verticalOffset += child.layout.Height
			verticalOffset += element.style.gap.vertical
		}
	}
}

type GrowLine struct {
	Used           uint16
	TotalFactor    uint32
	ChildrenAmount uint16
}

func grow(el *Element) {
	// DFS preordre

	lines := map[uint16]*GrowLine{}

	for _, child := range el.Children() {
		if _, ok := lines[child.layout.Line]; !ok {
			lines[child.layout.Line] = &GrowLine{}
		}
		childLine := lines[child.layout.Line]

		if el.style.layoutAxis == LAYOUT_HORIZONTAL {
			if g, ok := child.style.width.(growS); ok {
				childLine.TotalFactor += uint32(g.factor)
			}
			childLine.ChildrenAmount += 1
			childLine.Used += child.Width()
		} else {
			if g, ok := child.style.height.(growS); ok {
				childLine.TotalFactor += uint32(g.factor)
			}
			childLine.ChildrenAmount += 1
			childLine.Used += child.Height()
		}
	}

	for _, child := range el.Children() {
		line := lines[child.layout.Line]

		var left int32
		if el.style.layoutAxis == LAYOUT_HORIZONTAL {
			gap := (line.ChildrenAmount-1)*el.style.gap.horizontal - 1
			left = int32(el.Width()) - int32(line.Used+el.style.padding.left+el.style.padding.right+gap)
		} else {
			gap := (line.ChildrenAmount - 1) * el.style.gap.vertical
			left = int32(el.Height()) - int32(line.Used+el.style.padding.top+el.style.padding.bottom+gap)
		}

		if el.style.layoutAxis == LAYOUT_HORIZONTAL {
			if g, ok := child.style.width.(growS); ok {
				child.layout.Width += uint16(float64(left) * float64(g.factor) / float64(line.TotalFactor))
			}
		} else {
			if g, ok := child.style.height.(growS); ok {
				child.layout.Height += uint16(float64(left) * float64(g.factor) / float64(line.TotalFactor))
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
