package dsl

type RootS struct {
	*Element
}

func Root(element *Element) *RootS {
	sizeLayout(element)
	constraint(element, element.layoutAxisSize(), element.crossAxisSize())
	overflow(element)
	grow(element)
	sizeCross(element)
	position(element, 0, 0)
	// TODO grow cross axis

	return &RootS{
		Element: element,
	}
}

func overflow(element *Element) {
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

		if g, ok := child.style.axisSize(el.style.layoutAxis).(growS); ok {
			childLine.TotalFactor += uint32(g.factor)
		}
		childLine.ChildrenAmount += 1
		childLine.Used += child.axisSize(el.style.layoutAxis)
	}

	for _, child := range el.Children() {
		line := lines[child.layout.Line]

		var left int32
		gap := (line.ChildrenAmount - 1) * el.style.layoutAxisGap()
		padding := el.style.layoutAxisPadding()
		left = int32(el.layoutAxisSize()) - int32(line.Used+padding.start+padding.end+gap)

		if g, ok := child.style.axisSize(el.style.layoutAxis).(growS); ok {
			child.axisSizeSet(el.style.layoutAxis, child.axisSize(el.style.layoutAxis)+uint16(float64(left)*float64(g.factor)/float64(line.TotalFactor)))
		}
	}

	for _, child := range el.Children() {
		grow(child)
	}
}

func sizeCross(el *Element) {
	for _, child := range el.Children() {
		sizeCross(child)
	}
	padding := el.style.crossAxisPadding()
	gap := el.style.crossAxisGap()
	size := padding.start + padding.end

	var maxLineSize uint16 = 0
	var line uint16 = 0
	for _, child := range el.Children() {
		if child.layout.Line != line {
			line += 1
			size += maxLineSize + gap
			maxLineSize = 0
		}
		maxLineSize = max(maxLineSize, child.axisSize(el.style.oppositeAxis()))
	}
	size += maxLineSize

	if crossFitSize, ok := el.style.crossAxisSize().(fixedS); ok {
		el.crossAxisSet(crossFitSize.Size)
	} else {
		el.crossAxisSet(size)
	}
}

func sizeLayout(el *Element) {
	//DFS postordre
	for _, child := range el.Children() {
		sizeLayout(child)
	}
	// Since there is 3 modes FIXED/FIT/GROW check that if the mode is FIXED we discard the work and set the fix sizeLayout
	if crossSize, ok := el.style.crossAxisSize().(fixedS); ok {
		el.crossAxisSet(crossSize.Size)
	}

	if layoutSize, ok := el.style.layoutAxisSize().(fixedS); ok {
		el.layoutAxisSet(layoutSize.Size)
		return
	}

	// Get Padding for the layout Axis
	padding := el.style.layoutAxisPadding()

	// Init the sizeLayout for layout axis depending on the layout and add padding
	size := padding.start + padding.end

	// Iterate every child and add their sizeLayout to the main axis for this element
	for _, child := range el.Children() {
		size += child.axisSize(el.style.layoutAxis)
	}

	// Add the gap sizeLayout between children to the main axis sizeLayout
	var gap uint16 = 0
	if n := len(el.children); n > 1 {
		gap = uint16(n-1) * el.style.layoutAxisGap()
	}

	size += gap
	el.layoutAxisSet(size)
}

func constraint(el *Element, maxLayoutSize uint16, maxCrossSize uint16) {
	if _, ok := el.style.layoutAxisSize().(fixedS); !ok {
		el.layoutAxisSet(min(el.layoutAxisSize(), maxLayoutSize))
	}

	for _, child := range el.Children() {
		layoutPadding := el.style.layoutAxisPadding()
		crossPadding := el.style.crossAxisPadding()
		if child.style.layoutAxis == el.style.layoutAxis {
			constraint(child, el.layoutAxisSize()-layoutPadding.start-layoutPadding.end, maxCrossSize-crossPadding.start-crossPadding.end)
		} else {
			constraint(child, maxCrossSize-crossPadding.start-crossPadding.end, el.layoutAxisSize()-layoutPadding.start-layoutPadding.end)
		}
	}
}
