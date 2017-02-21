package main

import (
	"fmt"
	"math/rand"
)

type Map struct {
	w int
	h int
}

func (m Map) RandomPosition() Position {
	x := rand.Intn(m.w + 1)
	y := rand.Intn(m.h + 1)

	return Position{x, y}
}

type Position struct {
	x int
	y int
}

func (p Position) Equals(p2 Position) bool {
	return p.x == p2.x && p.y == p2.y
}

func (p Position) String() string {
	return fmt.Sprintf("Position{x: %v, y: %v}", p.x, p.y)
}

func (p Position) ManhattanDistance(p2 Position) int {
	return abs(p.x-p2.x) + abs(p.y-p2.y)
}

func (p Position) MoveTo(destination Position, steps int) Position {
	for i := 0; i < steps; i++ {
		if p.x < destination.x {
			p.x++
		} else if p.x > destination.x {
			p.x--
		} else if p.y < destination.y {
			p.y++
		} else if p.y > destination.y {
			p.y--
		} else {
			break
		}
	}

	return p
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
