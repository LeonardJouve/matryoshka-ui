package dsl

type ImageS struct {
	Style *StyleS
	Src   string
}

type ImageAttrs struct {
	Src string
}

type ImageOpt func(*ImageS)

func Image(src string, opts ...ImageOpt) *ImageS {
	i := ImageS{Src: src, Style: NewStyle()}
	for _, o := range opts {
		o(&i)
	}
	return &i
}

func (i ImageS) build() *Node {
	return &Node{
		Kind:  KindImage,
		Style: i.Style,
		ImageAttrs: ImageAttrs{
			Src: i.Src,
		},
	}
}

type ImageStyleModifier interface {
	StyleModifier
	image()
}

func ImageStyle(modifiers ...ImageStyleModifier) ImageOpt {
	return func(i *ImageS) {
		for _, m := range modifiers {
			m.applyStyle(i.Style)
		}
	}
}
