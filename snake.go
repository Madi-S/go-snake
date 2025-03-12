package main

type Coordinate struct {
	x, y int
}

type Direction int

type Snake struct {
	coords    []Coordinate
	color     Color
	direction Direction
}
