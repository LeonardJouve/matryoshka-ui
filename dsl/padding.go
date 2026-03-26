package dsl

type PaddingS struct {
	left   uint16
	top    uint16
	right  uint16
	bottom uint16
}

type PaddingModifier = func(padding *PaddingS)

func PaddingVertical(vertical uint16) PaddingModifier {
	return func(padding *PaddingS) {
		padding.top = vertical
		padding.bottom = vertical
	}
}

func PaddingHorizontal(horizontal uint16) PaddingModifier {
	return func(padding *PaddingS) {
		padding.left = horizontal
		padding.right = horizontal
	}
}

func PaddingLeft(left uint16) PaddingModifier {
	return func(padding *PaddingS) {
		padding.left = left
	}
}

func PaddingRight(right uint16) PaddingModifier {
	return func(padding *PaddingS) {
		padding.right = right
	}
}

func PaddingBottom(bottom uint16) PaddingModifier {
	return func(padding *PaddingS) {
		padding.bottom = bottom
	}
}

func PaddingTop(top uint16) PaddingModifier {
	return func(padding *PaddingS) {
		padding.top = top
	}
}
