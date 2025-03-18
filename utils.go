package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func anyArrowKeyIsPressed() bool {
	for key, _ := range keyDirections {
		if inpututil.IsKeyJustPressed(key) {
			return true
		}
	}
	return false
}
