package dsl

import "github.com/LeonardJouve/matryoshka-ui/utils"

type TextS struct {
	Style   *StyleS
	Content string
}

type TextOpt func(*TextS)

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

func FontSize(size uint16) TextOpt {
	return func(t *TextS) {
		t.Style.FontSize = size
	}
}
func TextColor(c utils.Color) TextOpt {
	return func(t *TextS) {
		t.Style.Color = c
	}
}

type TextAttrs struct {
	Content string
}
