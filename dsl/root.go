package dsl

func Root(element Element, measurer TextMeasurer) *Node {
	root := element.build()
	layoutEngine(root, measurer)
	return root
}

func layoutEngine(node *Node, measurer TextMeasurer) {
	sizing(node, measurer)
	grow(node)
	growCross(node)
	position(node, 0, 0)
}

func sizing(node *Node, measurer TextMeasurer) {
	switch node.Kind {
	case KindImage:
		if width, ok := node.Style.Width.(fixedS); ok {
			node.Layout.Width = width.Size
		}
		if height, ok := node.Style.Height.(fixedS); ok {
			node.Layout.Height = height.Size
		}
	case KindText:
		w, h := measurer.MeasureText(node.TextAttrs.Content, node.Style.FontSize)
		node.Layout.Width = w
		node.Layout.Height = h
	case KindDiv:
		{
			padding := node.Style.layoutAxisPadding()
			paddingCross := node.Style.crossAxisPadding()
			mainSize := padding.start + padding.end
			var crossSize uint16 = 0

			for _, child := range node.Children {
				sizing(child, measurer)
				mainSize += child.axisSize(node.Style.LayoutAxis)
				crossSize = max(crossSize, child.axisSize(node.Style.oppositeAxis()))
			}

			crossSize += paddingCross.start + paddingCross.end

			if n := len(node.Children); n > 1 {
				mainSize += uint16(n-1) * node.Style.layoutAxisGap()
			}

			switch size := node.Style.layoutAxisSize().(type) {
			case fixedS:
				node.setMainAxisSize(size.Size)
			default:
				node.setMainAxisSize(mainSize)
			}

			switch size := node.Style.crossAxisSize().(type) {
			case fixedS:
				node.setCrossAxisSize(size.Size)
			default:
				node.setCrossAxisSize(crossSize)
			}
		}
	}
}

func position(node *Node, x uint16, y uint16) {
	node.Layout.X = x
	node.Layout.Y = y

	px := x + node.Style.Padding.left
	py := y + node.Style.Padding.top

	for _, child := range node.Children {
		position(child, px, py)

		if node.Style.LayoutAxis == LAYOUT_HORIZONTAL {
			px += child.Layout.Width + node.Style.Gap.horizontal
		} else {
			py += child.Layout.Height + node.Style.Gap.vertical
		}
	}
}

func grow(node *Node) {
	// DFS preordre
	var totalFactor uint16 = 0
	padding := node.Style.layoutAxisPadding()
	used := padding.start + padding.end
	if n := len(node.Children); n > 1 {
		used += uint16(n-1) * node.Style.layoutAxisGap()
	}

	for _, child := range node.Children {
		if g, ok := child.Style.axisSize(node.Style.LayoutAxis).(growS); ok {
			totalFactor += g.factor
		}
		used += child.axisSize(node.Style.LayoutAxis)
	}

	var leftAvailable = max(node.mainAxisSize()-used, 0)
	for _, child := range node.Children {
		if g, ok := child.Style.axisSize(node.Style.LayoutAxis).(growS); ok {
			growSize := uint16(float64(leftAvailable) * float64(g.factor) / float64(totalFactor))
			child.setAxisSize(node.Style.LayoutAxis, child.axisSize(node.Style.LayoutAxis)+growSize)
		}
	}

	for _, child := range node.Children {
		grow(child)
	}
}

func growCross(node *Node) {
	// DFS preordre
	padding := node.Style.crossAxisPadding()

	var leftAvailable = node.crossAxisSize() - (padding.start + padding.end)
	for _, child := range node.Children {
		if _, ok := child.Style.axisSize(node.Style.oppositeAxis()).(growS); ok {
			child.setAxisSize(node.Style.oppositeAxis(), leftAvailable)
		}
	}

	for _, child := range node.Children {
		growCross(child)
	}
}
