package dsl

type RootS struct {
	*Element
}

func Root(element *Element) *RootS {
	size(element)
	overflow(element)
	grow(element)
	position(element, 0, 0)
	// TODO grow cross axis

	return &RootS{
		Element: element,
	}
}

func overflow(element *Element) {
	// DFS postordre
	for _, child := range element.Children() {
		overflow(child)
	}

	padding := element.style.layoutAxisPadding()
	var layout = padding.start + padding.end
	var line uint16 = 0

	for _, child := range element.Children() {
		childSize := child.axisSize(element.style.layoutAxis)
		if layout+childSize > element.layoutAxisSize() {
			line += 1
			layout = padding.start + padding.end + childSize
		} else {
			layout += childSize + element.style.layoutAxisGap()
		}
		child.layout.Line = line
	}
}

func position(element *Element, layoutOffset uint16, crossOffset uint16) {
	element.layoutPositionSet(layoutOffset)
	element.crossPositionSet(crossOffset)

	layoutOffset += element.style.layoutAxisPadding().start
	crossOffset += element.style.crossAxisPadding().start

	var maxCrossChild uint16 = 0

	var line uint16 = 0
	for _, child := range element.Children() {
		if line != child.layout.Line {
			line += 1
			crossOffset += maxCrossChild + element.style.crossAxisGap()
			layoutOffset = element.layoutAxisPosition() + element.style.layoutAxisPadding().start
			maxCrossChild = 0
		}
		maxCrossChild = max(maxCrossChild, child.axisSize(element.style.oppositeAxis()))

		if child.style.layoutAxis == element.style.layoutAxis {
			position(child, layoutOffset, crossOffset)
		} else {
			position(child, crossOffset, layoutOffset)
		}

		layoutOffset += element.style.layoutAxisGap() + child.axisSize(element.style.layoutAxis)
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

		if g, ok := child.style.layoutAxisSize().(growS); ok {
			childLine.TotalFactor += uint32(g.factor)
		}
		childLine.ChildrenAmount += 1
		childLine.Used += child.layoutAxisSize()
	}

	for _, child := range el.Children() {
		line := lines[child.layout.Line]

		var left int32
		gap := (line.ChildrenAmount - 1) * el.style.layoutAxisGap()
		padding := el.style.layoutAxisPadding()
		left = int32(el.layoutAxisSize()) - int32(line.Used+padding.start+padding.end+gap)

		if g, ok := child.style.layoutAxisSize().(growS); ok {
			child.layoutAxisSet(child.layoutAxisSize() + uint16(float64(left)*float64(g.factor)/float64(line.TotalFactor)))
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
	// Get Padding for the both Axis Horizontal/Vertical
	layoutPadding := el.style.layoutAxisPadding()
	crossPadding := el.style.crossAxisPadding()

	// Init the size for both axis depending on the layout and add padding
	layoutAxis := layoutPadding.start + layoutPadding.end
	crossAxis := crossPadding.start + crossPadding.end

	// Iterate every child and add their size to the main axis for this element, Calculate the biggest element on the cross axis for its size
	var maxCrossSize uint16 = 0
	for _, child := range el.Children() {
		layoutAxis += child.axisSize(el.style.layoutAxis)
		maxCrossSize = max(maxCrossSize, child.axisSize(el.style.oppositeAxis()))
	}

	// Add the gap size between children to the main axis size
	if n := len(el.children); n > 1 {
		layoutAxis += uint16(n-1) * el.style.layoutAxisGap()
	}
	// Set the cross axis size
	crossAxis += maxCrossSize

	// Since there is 3 modes FIXED/FIT/GROW check that if the mode is FIXED we discard the work and set the fix size
	if layoutSize, ok := el.style.layoutAxisSize().(fixedS); ok {
		el.layoutAxisSet(layoutSize.Size)
	} else {
		el.layoutAxisSet(layoutAxis)
	}

	if crossSize, ok := el.style.crossAxisSize().(fixedS); ok {
		el.crossAxisSet(crossSize.Size)
	} else {
		el.crossAxisSet(crossAxis)
	}
}
