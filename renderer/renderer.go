package renderer

import (
	"github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderer interface {
	InitWindow(width, height int, name string)
	CloseWindow()
	Render(element *dsl.Element)
}

type RaylibRenderer struct {
}

type RaylibInit struct {
	width, height int
	name          string
}

func NewRaylibRenderer() *RaylibRenderer {
	return &RaylibRenderer{}
}

func (renderer *RaylibRenderer) InitWindow(width, height int32, name string) {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.InitWindow(width, height, name)
}

func (renderer *RaylibRenderer) SetWindowFlag(flag uint32) {
	rl.SetWindowState(flag)
}

func (renderer *RaylibRenderer) CloseWindow() {
	rl.CloseWindow()
}

func (renderer *RaylibRenderer) Render(node *dsl.Node) {
	switch node.Kind {
	case dsl.KindDiv:
		renderer.renderDiv(node)
		for _, child := range node.DivAttrs.Children {
			renderer.Render(child)
		}
	}
}

func (renderer *RaylibRenderer) renderDiv(node *dsl.Node) {
	renderRectangle(node.Layout.X, node.Layout.Y, node.Layout.Width, node.Layout.Height, node.Style.Color)
}

func renderRectangle(x uint16, y uint16, width uint16, height uint16, color utils.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), rl.NewColor(color.Red, color.Green, color.Blue, 255))
}
