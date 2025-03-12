package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	snake Snake
}

func (g *Game) Update() error {
	log.Printf("FPS: %.2f\n", ebiten.ActualFPS())

	if err := handleClose(); err != nil {
		return err
	}
	if err := handleKeyArrowsInput(g); err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, WELCOME_TEXT, WELCOME_TEXT_X, WELCOME_TEXT_Y)
	for _, c := range g.snake.coords {
		vector.DrawFilledRect(
			screen, float32(c.x*GRID_SIZE), float32(c.y*GRID_SIZE),
			float32(GRID_SIZE), float32(GRID_SIZE), g.snake.color, true,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return MAX_WIDTH, MAX_HEIGHT
}

func setup() {
	ebiten.SetWindowTitle("Snake 2D")
	ebiten.SetWindowSize(MAX_WIDTH, MAX_HEIGHT)
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	ebiten.SetFullscreen(false)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(120)
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetRunnableOnUnfocused(false)
	ebiten.SetWindowClosingHandled(true)
}

func main() {
	setup()

	game := Game{
		snake: Snake{
			coords: []Coordinate{{x: MAX_WIDTH/2/GRID_SIZE - 1, y: MAX_HEIGHT/2/GRID_SIZE - 1}},
			color:  yellowish,
		},
	}
	gameOptions := ebiten.RunGameOptions{
		SkipTaskbar:       false,
		InitUnfocused:     false,
		ScreenTransparent: false,
		ColorSpace:        ebiten.ColorSpaceSRGB,
	}

	if err := ebiten.RunGameWithOptions(&game, &gameOptions); err != nil {
		log.Fatal(err)
	}
}
