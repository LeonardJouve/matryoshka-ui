package dsl

func Root(element Element) *Node {
	root := &Node{
		Kind:   KindDiv,
		Style:  NewStyle(),
		Layout: NewLayout(),
		DivAttrs: DivAttrs{
			Children: []*Node{element.build()},
		},
	}
	return root
}

func Layout(root *Node, width uint16, height uint16, measurer TextMeasurer) *Node {
	root.Style.Width = Fixed(width)
	root.Style.Height = Fixed(height)
	layoutEngine(root, measurer)
	return root
}

func layoutEngine(node *Node, measurer TextMeasurer) {
	sizing(node, measurer)
	flex(node)
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
			padding := node.Style.mainAxisPadding()
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

	mainPadding := node.Style.mainAxisPadding()
	mainAvailable := node.mainAxisSize() - mainPadding.start - mainPadding.end

	crossPadding := node.Style.crossAxisPadding()
	crossAvailable := node.crossAxisSize() - crossPadding.start - crossPadding.end

	for _, child := range node.Children {
		mainAvailable -= child.axisSize(node.Style.LayoutAxis)
	}

	for i, child := range node.Children {
		mainOffset, mainGap := justify(node.Style.Justify, mainAvailable, uint16(len(node.Children)))

		crossOffset := align(node.Style.Align, crossAvailable-child.axisSize(node.Style.oppositeAxis()))
		var offsetX, offsetY uint16

		if node.Style.LayoutAxis == LAYOUT_HORIZONTAL {
			if i == 0 {
				px += mainOffset
			} else {
				px += mainGap
			}
			offsetY = crossOffset
		} else {
			if i == 0 {
				py += mainOffset
			} else {
				py += mainGap
			}
			offsetX = crossOffset
		}

		position(child, px+offsetX, py+offsetY)

		if node.Style.LayoutAxis == LAYOUT_HORIZONTAL {
			px += child.Layout.Width

			if node.Style.Justify != JustifyBetween && node.Style.Justify != JustifyAround {
				px += node.Style.Gap.horizontal
			}
		} else {
			py += child.Layout.Height

			if node.Style.Justify != JustifyBetween && node.Style.Justify != JustifyAround {
				py += node.Style.Gap.vertical
			}
		}
	}
}

func flex(node *Node) {
	mainPadding := node.Style.mainAxisPadding()
	crossPadding := node.Style.crossAxisPadding()

	mainInner := node.mainAxisSize() - mainPadding.start - mainPadding.end
	crossInner := node.crossAxisSize() - crossPadding.start - crossPadding.end

	// resolve percentage children against this node's inner (content-box) size
	for _, child := range node.Children {
		if p, ok := child.Style.axisSize(node.Style.LayoutAxis).(percentS); ok {
			child.setAxisSize(node.Style.LayoutAxis, uint16(float64(mainInner)*p.Ratio))
		}
		if p, ok := child.Style.axisSize(node.Style.oppositeAxis()).(percentS); ok {
			child.setAxisSize(node.Style.oppositeAxis(), uint16(float64(crossInner)*p.Ratio))
		}
	}

	// grow main
	var totalFactor uint16 = 0
	used := mainPadding.start + mainPadding.end

	if n := len(node.Children); n > 1 {
		used += uint16(n-1) * node.Style.layoutAxisGap()
	}

	for _, child := range node.Children {
		if g, ok := child.Style.axisSize(node.Style.LayoutAxis).(growS); ok {
			totalFactor += g.factor
		}
		used += child.axisSize(node.Style.LayoutAxis)
	}

	var leftAvailable uint16 = 0
	if node.mainAxisSize() > used {
		leftAvailable = node.mainAxisSize() - used
	}

	for _, child := range node.Children {
		if g, ok := child.Style.axisSize(node.Style.LayoutAxis).(growS); ok && totalFactor > 0 {
			growSize := uint16(float64(leftAvailable) * float64(g.factor) / float64(totalFactor))
			child.setAxisSize(node.Style.LayoutAxis, child.axisSize(node.Style.LayoutAxis)+growSize)
		}
	}

	// grow cross
	for _, child := range node.Children {
		if _, ok := child.Style.axisSize(node.Style.oppositeAxis()).(growS); ok {
			child.setAxisSize(node.Style.oppositeAxis(), crossInner)
		}
	}

	for _, child := range node.Children {
		flex(child)
	}
}

func justify(j JustifyT, available uint16, count uint16) (offset uint16, gap uint16) {
	if count == 0 {
		return 0, 0
	}

	switch j {
	case JustifyCenter:
		return available / 2, 0
	case JustifyEnd:
		return available, 0
	case JustifyBetween:
		if count == 1 {
			return 0, 0
		}

		return 0, available / (count - 1)
	case JustifyAround:
		g := available / count
		return g / 2, g
	default:
		return 0, 0
	}
}

func align(a AlignT, available uint16) uint16 {
	switch a {
	case AlignCenter:
		return available / 2
	case AlignEnd:
		return available
	default:
		return 0
	}
}
