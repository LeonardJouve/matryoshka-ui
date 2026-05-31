package dsl2

func Layout(root *Node) {
	// Root gets no outer constraint (0 means "unconstrained / use intrinsic").
	measure(root, 0, 0)
	flex(root)
	position(root, 0, 0)
}

// measure computes a node's size. availW/availH are the content-box dimensions
// offered by the parent (0 = unconstrained). A container resolves its own size
// BEFORE breaking lines so wrapping and flex have a real main-axis extent.
func measure(n *Node, availW, availH int) (w, h int) {
	switch n.Kind {
	//case dsl2.KindText:
	//	w, h = measureText(n.Content, n.Style)

	//case dsl2.KindImage:
	//	w, h = n.Style.W, n.Style.H

	case KindDiv:
		dir := n.Style.Dir
		pad := n.Style.Padding

		// Resolve this container's size first: explicit Style wins, else the
		// space the parent offered along the relevant axis, else intrinsic.
		if n.Style.W > 0 {
			n.Layout.W = n.Style.W
		} else if availW > 0 {
			n.Layout.W = availW
		}
		if n.Style.H > 0 {
			n.Layout.H = n.Style.H
		} else if availH > 0 {
			n.Layout.H = availH
		}

		// What we offer children:
		//  - Cross axis: always offer the parent's content extent. In block
		//    flow a child spans its parent's cross size (e.g. a Row inside a
		//    fixed-width Column may use the full width). Offered as available,
		//    not forced — a child with an explicit/ intrinsic size ignores it.
		//  - Main axis: offer it only when THIS container wraps, since only a
		//    wrapping container needs its children to know the main extent for
		//    line-breaking. Non-wrap containers keep children intrinsic so
		//    siblings and flex-grow size correctly.
		var cw, ch int
		if dir == Row {
			ch = n.Layout.H - pad*2 // cross = height
			if n.Style.Wrap == WrapOn {
				cw = n.Layout.W - pad*2 // main = width
			}
		} else {
			cw = n.Layout.W - pad*2 // cross = width
			if n.Style.Wrap == WrapOn {
				ch = n.Layout.H - pad*2 // main = height
			}
		}
		for _, c := range n.Children {
			measure(c, cw, ch)
		}

		n.lines = breakLines(n, dir)

		maxMain, sumCross := 0, 0
		for _, ln := range n.lines {
			if ln.mainUsed > maxMain {
				maxMain = ln.mainUsed
			}
			sumCross += ln.crossSize
		}

		// Intrinsic size from content, used only where nothing was resolved.
		if dir == Row {
			w, h = maxMain+pad*2, sumCross+pad*2
		} else {
			w, h = sumCross+pad*2, maxMain+pad*2
		}
	}

	// A resolved (explicit or inherited) dimension overrides intrinsic.
	if n.Layout.W > 0 {
		w = n.Layout.W
	}
	if n.Layout.H > 0 {
		h = n.Layout.H
	}
	if n.Style.W > 0 {
		w = n.Style.W
	}
	if n.Style.H > 0 {
		h = n.Style.H
	}

	n.Layout.W, n.Layout.H = w, h
	return w, h
}
