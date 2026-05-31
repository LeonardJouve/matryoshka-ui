package main

import (
	. "github.com/LeonardJouve/matryoshka-ui/dsl2"
	"github.com/LeonardJouve/matryoshka-ui/renderer"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func ui() *RootS {
	return Root(Div(Width(100), Height(200), Background(Color{R: 200, G: 100, B: 50}), LayoutAxis(Column), Padding(PaddingHorizontal(10)), Children(
		Div(Height(30), LayoutAxis(Row), Background(Color{R: 100, G: 10, B: 250})),
		Div(Height(12), Width(10)),
	)))
}

func main() {
	var width uint16 = 800
	var height uint16 = 800
	renderer := renderer.NewRaylibRenderer()

	renderer.InitWindow(int32(width), int32(height), "Testing101")
	renderer.SetWindowFlag(rl.FlagWindowResizable)
	defer renderer.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Color{R: 24, G: 24, B: 27, A: 255})
		renderer.Render(ui())

		rl.EndDrawing()
	}
}
