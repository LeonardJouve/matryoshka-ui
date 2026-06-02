package renderer

import (
	"fmt"
	"os"
	"time"

	"github.com/LeonardJouve/matryoshka-ui/dsl"
	"github.com/LeonardJouve/matryoshka-ui/utils"
	"github.com/gdamore/tcell/v2"
)

type TerminalRenderer struct {
	screen tcell.Screen
	// Scale factors to convert pixels to terminal cells.
	// Defaults usually assume an 8x16 pixel terminal font.
	scaleX float64
	scaleY float64
}

func NewTerminalRenderer() *TerminalRenderer {
	return &TerminalRenderer{
		scaleX: 8.0,  // 8 pixels per terminal column
		scaleY: 16.0, // 16 pixels per terminal row
	}
}

// InitWindow creates and initializes the terminal screen.
func (renderer *TerminalRenderer) InitWindow(width, height int, name string) {
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create screen: %v\n", err)
		os.Exit(1)
	}
	if err := s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize screen: %v\n", err)
		os.Exit(1)
	}

	// Set default terminal styling
	defStyle := tcell.StyleDefault.
		Background(tcell.ColorReset).
		Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	renderer.screen = s
}

func (renderer *TerminalRenderer) CloseWindow() {
	if renderer.screen != nil {
		renderer.screen.Fini()
	}
}

// Render runs the main application loop.
func (renderer *TerminalRenderer) Render(node *dsl.Node) {
	quit := make(chan struct{})

	// Event loop to handle terminal events (like exiting)
	go func() {
		for {
			ev := renderer.screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				// Exit on Escape or Ctrl+C
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					close(quit)
					return
				}
			case *tcell.EventResize:
				renderer.screen.Sync()
			}
		}
	}()

	// Rendering loop (simulating your 60FPS target)
	ticker := time.NewTicker(time.Second / 60)
	defer ticker.Stop()

	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			// Clear background (using a custom dark gray to match your Raylib version)
			bgStyle := tcell.StyleDefault.Background(tcell.NewRGBColor(24, 24, 27))
			renderer.screen.SetStyle(bgStyle)
			renderer.screen.Clear()

			renderer.renderNode(node)

			renderer.screen.Show()
		}
	}
}

func (renderer *TerminalRenderer) renderNode(node *dsl.Node) {
	switch node.Kind {
	case dsl.KindDiv:
		renderer.renderDiv(node)
		for _, child := range node.DivAttrs.Children {
			renderer.renderNode(child)
		}
	case dsl.KindText:
		style := tcell.StyleDefault.Foreground(colorToTcell(node.Style.Color))

		// Scale the text coordinates
		x := int(float64(node.Layout.X) / renderer.scaleX)
		y := int(float64(node.Layout.Y) / renderer.scaleY)

		// Draw string character by character
		for i, ch := range node.TextAttrs.Content {
			renderer.screen.SetContent(x+i, y, ch, nil, style)
		}
	case dsl.KindImage:
		style := tcell.StyleDefault.
			Background(tcell.ColorDarkGray).
			Foreground(tcell.ColorWhite)

		// Scale image coordinates and dimensions
		x := int(float64(node.Layout.X) / renderer.scaleX)
		y := int(float64(node.Layout.Y) / renderer.scaleY)
		w := int(float64(node.Layout.Width) / renderer.scaleX)
		h := int(float64(node.Layout.Height) / renderer.scaleY)

		// Ensure it renders at least 1 cell wide/high if it exists
		if w == 0 && node.Layout.Width > 0 {
			w = 1
		}
		if h == 0 && node.Layout.Height > 0 {
			h = 1
		}

		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				char := ' '
				if j == 0 && w >= 3 {
					if i == 0 {
						char = 'i'
					}
					if i == 1 {
						char = 'm'
					}
					if i == 2 {
						char = 'g'
					}
				}
				renderer.screen.SetContent(x+i, y+j, char, nil, style)
			}
		}
	}
}

func (renderer *TerminalRenderer) renderDiv(node *dsl.Node) {
	style := tcell.StyleDefault.Background(colorToTcell(node.Style.Color))

	// Scale the div coordinates and dimensions
	x := int(float64(node.Layout.X) / renderer.scaleX)
	y := int(float64(node.Layout.Y) / renderer.scaleY)
	w := int(float64(node.Layout.Width) / renderer.scaleX)
	h := int(float64(node.Layout.Height) / renderer.scaleY)

	// Prevent elements with real dimensions from scaling down to 0 and vanishing
	if w == 0 && node.Layout.Width > 0 {
		w = 1
	}
	if h == 0 && node.Layout.Height > 0 {
		h = 1
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			renderer.screen.SetContent(x+i, y+j, ' ', nil, style)
		}
	}
}

func (renderer *TerminalRenderer) MeasureText(content string, fontSize uint16) (uint16, uint16) {
	// Text width is just the string length.
	// We multiply it back into "pixel space" so the underlying layout engine
	// calculates the bounding boxes correctly.
	pixelWidth := uint16(float64(len(content)) * renderer.scaleX)
	pixelHeight := uint16(1.0 * renderer.scaleY)

	return pixelWidth, pixelHeight
}

func colorToTcell(color utils.Color) tcell.Color {
	return tcell.NewRGBColor(int32(color.Red), int32(color.Green), int32(color.Blue))
}
