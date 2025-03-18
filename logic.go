package main

import (
	"errors"
	"log"
	"slices"
	"time"

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

var keyDirections map[ebiten.Key]Direction = map[ebiten.Key]Direction{
	ebiten.KeyArrowUp:    UP,
	ebiten.KeyArrowDown:  DOWN,
	ebiten.KeyArrowLeft:  LEFT,
	ebiten.KeyArrowRight: RIGHT,
}

// Handles user input of key arrows by calculating new
// coordinates for head of the snake and removing its tail.
//
// Throws error if snake is out of bounds.
func handleKeyArrowsInput(g *Game) error {
	for key, keyDirection := range keyDirections {
		if inpututil.IsKeyJustPressed(key) {
			if keyDirection != g.snake.direction && keyDirection != (g.snake.direction^1) {
				g.snake.direction = keyDirection
				break
			}
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

	if slices.Contains(g.snake.coords, newHead) {
		return errors.New("Snake ate itself, game is over")
	}

	coordUntil := len(g.snake.coords) - 1
	for i, food := range g.foods {
		if newHead == food.coord {
			coordUntil = len(g.snake.coords)
			g.foods = append(g.foods[:i], g.foods[i+1:]...)
			break
		}
	}

	coordsWithoutTail := g.snake.coords[:coordUntil]
	g.snake.coords = append([]Coordinate{newHead}, coordsWithoutTail...)
	return nil
}

// Handles random food spawn
func handleFoodSpawn(g *Game) {
	if time.Since(g.lastFoodSpawn) > FOOD_SPAWN_INTERVAL_SECONDS {
		foodX := generateRandomNumber(0, MAX_WIDTH) / GRID_SIZE
		foodY := generateRandomNumber(0, MAX_HEIGHT) / GRID_SIZE
		foodColor := yellowish
		for foodColor == yellowish {
			foodColor = colors[generateRandomNumber(0, len(colors)-1)]
		}
		newFood := Food{color: foodColor, coord: Coordinate{foodX, foodY}}
		g.foods = append(g.foods, newFood)
		g.lastFoodSpawn = time.Now()
	}
}
