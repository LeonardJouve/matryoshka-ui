package dsl

type GapS struct {
	vertical   uint16
	horizontal uint16
}

type GapModifier = func(gap *GapS)

func GapVertical(vertical uint16) GapModifier {
	return func(gap *GapS) {
		gap.vertical = vertical
	}
}

func GapHorizontal(horizontal uint16) GapModifier {
	return func(gap *GapS) {
		gap.horizontal = horizontal
	}
}
