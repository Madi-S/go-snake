package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Returns error if user wishes to close the game.
func handleClose() error {
	if ebiten.IsWindowBeingClosed() {
		log.Println("Game was closed")
		return errors.New("Game was closed")
	}
	return nil
}

// Handles user input of key arrows by calculating new
// coordinates for head of the snake and removing its tail.
//
// Throws error if snake is out of bounds.
func handleKeyArrowsInput(g *Game) error {
	keyCoordDeltas := map[ebiten.Key]Coordinate{
		ebiten.KeyArrowUp:    {0, -1},
		ebiten.KeyArrowDown:  {0, 1},
		ebiten.KeyArrowLeft:  {-1, 0},
		ebiten.KeyArrowRight: {1, 0},
	}

	newCoord := g.snake.coords[0]

	for key, coordDelta := range keyCoordDeltas {
		if inpututil.IsKeyJustPressed(key) {
			newCoord.x += coordDelta.x
			newCoord.y += coordDelta.y

			factualX := newCoord.x * GRID_SIZE
			factualY := newCoord.y * GRID_SIZE
			if factualX >= MAX_WIDTH || factualX < 0 || factualY >= MAX_HEIGHT || factualY < 0 {
				return errors.New("Snake is out of bounds, game is over")
			}
			break
		}
	}

	g.snake.coords = append([]Coordinate{newCoord}, g.snake.coords...)
	return nil
}
