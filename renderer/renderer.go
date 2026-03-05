package renderer

import (
	"github.com/LeonardJouve/matryoshka-ui/elements"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderer interface {
	Render()
}

type RaylibRenderer struct {
}

func newRaylibRenderer() RaylibRenderer {
	return RaylibRenderer{}
}

func (renderer *RaylibRenderer) Render(element *elements.Element) {
	var el = element
	renderRectangle(el.GetX(), el.GetY(), el.GetWidth(), el.GetHeight(), el.Color)
	for _, child := range el.Children {
		renderer.Render(child)
	}
}

func renderRectangle(x uint16, y uint16, width uint16, height uint16, color utils.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), rl.NewColor(color.Red, color.Green, color.Blue, 255))
}
