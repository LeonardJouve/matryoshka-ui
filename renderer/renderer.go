package renderer

import (
	"github.com/LeonardJouve/matryoshka-ui/dsl"
)

type Renderer interface {
	InitWindow(width, height int, name string)
	CloseWindow()
	Render(element *dsl.Node)
}
