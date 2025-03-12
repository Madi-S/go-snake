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
	keyDirections := map[ebiten.Key]Direction{
		ebiten.KeyArrowUp:    UP,
		ebiten.KeyArrowDown:  DOWN,
		ebiten.KeyArrowLeft:  LEFT,
		ebiten.KeyArrowRight: RIGHT,
	}

	for key, keyDirection := range keyDirections {
		if inpututil.IsKeyJustPressed(key) {
			g.snake.direction = keyDirection

			break
		}
	}

	newHead := g.snake.coords[0]
	switch g.snake.direction {
	case UP:
		newHead.y -= 1
	case DOWN:
		newHead.y += 1
	case LEFT:
		newHead.x -= 1
	case RIGHT:
		newHead.x += 1
	}

	factualX, factualY := newHead.x*GRID_SIZE, newHead.y*GRID_SIZE
	if factualX >= MAX_WIDTH || factualX < 0 || factualY >= MAX_HEIGHT || factualY < 0 {
		return errors.New("Snake is out of bounds, game is over")
	}

	coordsWithoutTail := g.snake.coords[:len(g.snake.coords)-1]
	g.snake.coords = append([]Coordinate{newHead}, coordsWithoutTail...)
	return nil
}
