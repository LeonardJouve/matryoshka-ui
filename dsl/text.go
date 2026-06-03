package dsl

type TextS struct {
	Style   *StyleS
	Content string
}

type TextOpt func(*TextS)
type TextStyleModifier interface {
	StyleModifier
	text()
}
type TextOnly func(*StyleS)

func (f TextOnly) applyStyle(s *StyleS) { f(s) }
func (f TextOnly) text()                {}

func Text(content string, opts ...TextOpt) *TextS {
	t := TextS{Content: content, Style: NewStyle()}
	for _, o := range opts {
		o(&t)
	}
	return &t
}

func (t *TextS) build() *Node {
	return &Node{
		Kind:  KindText,
		Style: t.Style,
		TextAttrs: TextAttrs{
			Content: t.Content,
		},
	}
}

func TextStyle(modifiers ...TextStyleModifier) TextOpt {
	return func(t *TextS) {
		for _, m := range modifiers {
			m.applyStyle(t.Style)
		}
	}
}

type TextAttrs struct {
	Content string
}

func FontSize(size uint16) TextOnly {
	return func(style *StyleS) {
		style.FontSize = size
	}
}
