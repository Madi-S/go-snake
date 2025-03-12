package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	GRID_SIZE  int = 20
	MAX_WIDTH  int = 110
	MAX_HEIGHT int = 480
)

type Game struct {
	snake Snake
}

func (g *Game) Update() error {
	// log.Printf("FPS: %.2f\n", ebiten.ActualFPS())
	if ebiten.IsWindowBeingClosed() {
		log.Println("Game was closed")
		return errors.New("Game was closed")
	}

	// keys := []any{ebiten.KeyArrowDown, ebiten.KeyArrowLeft}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		log.Println("Up arrow pressed!")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		log.Println("Down arrow pressed!")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		log.Println("Left arrow pressed!")
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		log.Println("Right arrow pressed!")
	}

	return nil
}

const (
	WELCOME_TEXT           = "Welcome to Snake2D"
	WELCOME_TEXT_WIDTH int = 110
	WELCOME_TEXT_X     int = (MAX_WIDTH - WELCOME_TEXT_WIDTH) / 2
	WELCOME_TEXT_Y     int = 10
)

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, WELCOME_TEXT, WELCOME_TEXT_X, WELCOME_TEXT_Y)
	for _, c := range g.snake.coords {
		vector.DrawFilledRect(
			screen, float32(c.x*GRID_SIZE), float32(c.y*GRID_SIZE),
			float32(GRID_SIZE), float32(GRID_SIZE), g.snake.color, true,
		)
		fmt.Printf("%f %f %f\n", float32(c.x*GRID_SIZE), float32(c.y*GRID_SIZE), float32(GRID_SIZE))
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
