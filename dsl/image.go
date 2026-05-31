package dsl

type ImageS struct {
	Style *StyleS
	Src   string
}

type ImageAttrs struct {
	Src string
}

type ImageOpt func(*ImageS)

func Image(src string, opts ...ImageOpt) ImageS {
	i := ImageS{Src: src}
	for _, o := range opts {
		o(&i)
	}
	return i
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
