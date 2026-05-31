package dsl

type TextMeasurer interface {
	MeasureText(content string, fontSize uint16) (width uint16, height uint16)
}
