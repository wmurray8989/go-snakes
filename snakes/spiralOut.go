package snakes

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/player"
	"github.com/wmurray8989/go-snakes/position"
)

// SpiralOut spirals outward from starting position
var SpiralOut = player.Player{
	Name:           "Spiral Out",
	Author:         "Null",
	ColorPrimary:   sdl.Color{R: 255, G: 255, B: 255, A: 255},
	ColorSecondary: sdl.Color{R: 0, G: 0, B: 0, A: 255},
	Strategy: func(self []position.Position, opponent []position.Position) position.Position {
		currentPosition := self[len(self)-1]

		// If this is the first move, go up
		if len(self) == 1 {
			return position.Position{
				X: currentPosition.X,
				Y: currentPosition.Y + 1,
			}
		}

		// get current direction
		previousPosition := self[len(self)-2]
		currentDirection := previousPosition.DirectionTo(currentPosition)

		// try turning right
		turnRight := currentPosition.DirectionAdd(currentDirection.TurnRight())
		if turnRight.IsValidMove(self, opponent) {
			return turnRight
		}

		// Try moving forward
		forward := currentPosition.DirectionAdd(currentDirection)
		if forward.IsValidMove(self, opponent) {
			return forward
		}

		// try turning left
		turnLeft := currentPosition.DirectionAdd(currentDirection.TurnLeft())
		if turnLeft.IsValidMove(self, opponent) {
			return turnLeft
		}

		// Give up
		return currentPosition
	},
}
