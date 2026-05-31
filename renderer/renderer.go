package renderer

import (
	"fmt"

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
	case dsl.KindText:
		textColor := node.Style.Color
		rl.DrawText(node.TextAttrs.Content, int32(node.Layout.X), int32(node.Layout.Y), int32(node.Style.FontSize), colorToRL(textColor))
	case dsl.KindImage:
		if tex, ok := loadTexture(node.ImageAttrs.Src); ok {
			// scale the texture into the node's computed rect.
			src := rl.NewRectangle(0, 0, float32(tex.Width), float32(tex.Height))
			dst := rl.NewRectangle(float32(node.Layout.X), float32(node.Layout.Y), float32(node.Layout.Width), float32(node.Layout.Height))
			fmt.Println()
			rl.DrawTexturePro(tex, src, dst, rl.NewVector2(0, 0), 0, rl.White)
		} else {
			// fallback placeholder when the image file isn't found.
			rl.DrawRectangle(int32(node.Layout.X), int32(node.Layout.Y), int32(node.Layout.Width), int32(node.Layout.Height), rl.NewColor(200, 200, 200, 255))
			rl.DrawRectangleLines(int32(node.Layout.X), int32(node.Layout.Y), int32(node.Layout.Width), int32(node.Layout.Height), rl.DarkGray)
			rl.DrawText("img", int32(node.Layout.X)+4, int32(node.Layout.Y)+4, 10, rl.DarkGray)
		}
	}
}

func (renderer *RaylibRenderer) renderDiv(node *dsl.Node) {
	renderRectangle(node.Layout.X, node.Layout.Y, node.Layout.Width, node.Layout.Height, node.Style.Color)
}

func renderRectangle(x uint16, y uint16, width uint16, height uint16, color utils.Color) {
	rl.DrawRectangle(int32(x), int32(y), int32(width), int32(height), colorToRL(color))
}

func (renderer *RaylibRenderer) MeasureText(content string, fontSize uint16) (uint16, uint16) {
	w := rl.MeasureText(content, int32(fontSize))
	return uint16(w), fontSize
}

func colorToRL(color utils.Color) rl.Color {
	return rl.NewColor(color.Red, color.Green, color.Blue, 255)
}

var textureCache = map[string]rl.Texture2D{}

func loadTexture(path string) (rl.Texture2D, bool) {
	if t, found := textureCache[path]; found {
		return t, t.ID != 0
	}

	t := rl.LoadTexture(path)
	textureCache[path] = t

	return t, t.ID != 0
}
