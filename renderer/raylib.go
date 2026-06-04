package renderer

import (
	"github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RaylibRenderer struct {
}

func NewRaylibRenderer() *RaylibRenderer {
	return &RaylibRenderer{}
}

func (renderer *RaylibRenderer) InitWindow(width, height uint16, name string) {
	rl.InitWindow(int32(width), int32(height), name)
}

func (renderer *RaylibRenderer) SetWindowFlag(flag uint32) {
	rl.SetWindowState(flag)
}

func (renderer *RaylibRenderer) CloseWindow() {
	rl.CloseWindow()
}

func (renderer *RaylibRenderer) Render(layout *dsl.Node) {
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{R: 24, G: 24, B: 27, A: 255})
		root := dsl.Layout(layout, uint16(rl.GetScreenWidth()), uint16(rl.GetScreenHeight()), renderer)
		renderer.render(root)

		rl.EndDrawing()
	}
}

func (renderer *RaylibRenderer) render(node *dsl.Node) {
	switch node.Kind {
	case dsl.KindDiv:
		renderer.renderDiv(node)
		for _, child := range node.DivAttrs.Children {
			renderer.render(child)
		}
	case dsl.KindText:
		textColor := node.Style.Color
		rl.DrawText(node.TextAttrs.Content, int32(node.Layout.X), int32(node.Layout.Y), int32(node.Style.FontSize), colorToRL(textColor))
	case dsl.KindImage:
		if tex, ok := loadTexture(node.ImageAttrs.Src); ok {
			// scale the texture into the node's computed rect.
			src := rl.NewRectangle(0, 0, float32(tex.Width), float32(tex.Height))
			dst := rl.NewRectangle(float32(node.Layout.X), float32(node.Layout.Y), float32(node.Layout.Width), float32(node.Layout.Height))
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
	renderRectangle(node.Layout.X, node.Layout.Y, node.Layout.Width, node.Layout.Height, node.Style.BorderRadius, node.Style.BorderWidth, node.Style.BorderColor, node.Style.Color)
}

func renderRectangle(x uint16, y uint16, width uint16, height uint16, radius float32, borderWidth uint16, borderColor utils.Color, color utils.Color) {
	rec := rl.NewRectangle(float32(x), float32(y), float32(width), float32(height))

	// fill
	if radius == 0 {
		rl.DrawRectanglePro(rec, rl.NewVector2(0, 0), 0, colorToRL(color))
	} else {
		rl.DrawRectangleRounded(rec, radius, 16, colorToRL(color))
	}

	// border
	if borderWidth > 0 {
		if radius <= 0 {
			rl.DrawRectangleLinesEx(rec, float32(borderWidth), colorToRL(borderColor))
		} else {
			rl.DrawRectangleRoundedLinesEx(rec, radius, 16, float32(borderWidth), colorToRL(borderColor))
		}
	}

	rl.DrawRectangleRounded(rec, radius, 32, colorToRL(color))
}

func (renderer *RaylibRenderer) MeasureText(content string, fontSize uint16) (uint16, uint16) {
	w := rl.MeasureText(content, int32(fontSize))
	return uint16(w), fontSize
}

func colorToRL(color utils.Color) rl.Color {
	return rl.NewColor(color.Red, color.Green, color.Blue, color.Alpha)
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
