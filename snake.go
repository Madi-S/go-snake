package main

type Coordinate struct {
	x, y int
}

type Snake struct {
	coords []Coordinate
	color  Color
}
