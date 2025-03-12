package main

// Resolution params
const (
	GRID_SIZE  int = 20
	MAX_WIDTH  int = 640
	MAX_HEIGHT int = 480
)

// Welcome text params
const (
	WELCOME_TEXT           = "Welcome to Snake2D"
	WELCOME_TEXT_WIDTH int = 110
	WELCOME_TEXT_X     int = (MAX_WIDTH - WELCOME_TEXT_WIDTH) / 2
	WELCOME_TEXT_Y     int = 10
)

// Snake directions
const (
	UP    Direction = iota // 0
	DOWN                   // 1
	LEFT                   // 2
	RIGHT                  // 3
)
