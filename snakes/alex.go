package snakes

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/direction"
	"github.com/wmurray8989/go-snakes/player"
	"github.com/wmurray8989/go-snakes/position"
)

// TemplateSnake moves up
var AlexSnake = player.Player{
	Name:           "Mr Noodle",
	Author:         "Alex Merkert",
	ColorPrimary:   sdl.Color{R: 170, G: 0, B: 20, A: 255},
	ColorSecondary: sdl.Color{R: 220, G: 200, B: 30, A: 255},
	Strategy: func(self []position.Position, opponent []position.Position) position.Position {
		currentPosition := self[len(self)-1]
		opponentPosition := opponent[len(opponent)-1]

		// If this is the first move, go towards the closest wall
		if len(self) == 1 {
			upProjection := position.Position{
				X: currentPosition.X,
				Y: 49,
			}
			rightProjection := position.Position{
				X: 49,
				Y: currentPosition.Y,
			}
			downProjection := position.Position{
				X: currentPosition.X,
				Y: 0,
			}
			leftProjection := position.Position{
				X: 0,
				Y: currentPosition.Y,
			}

			distanceUp := currentPosition.DistanceTo(upProjection)
			distanceRight := currentPosition.DistanceTo(rightProjection)
			distanceDown := currentPosition.DistanceTo(downProjection)
			distanceLeft := currentPosition.DistanceTo(leftProjection)

			if distanceUp <= distanceDown && distanceUp <= distanceLeft && distanceUp <= distanceRight {
				return currentPosition.DirectionAdd(direction.Up)
			}

			if distanceRight <= distanceDown && distanceRight <= distanceLeft {
				return currentPosition.DirectionAdd(direction.Right)
			}

			if distanceDown <= distanceLeft {
				return currentPosition.DirectionAdd(direction.Down)
			}

			return currentPosition.DirectionAdd(direction.Left)
		}

		// get current direction
		previousPosition := self[len(self)-2]
		currentDirection := previousPosition.DirectionTo(currentPosition)

		// Try moving forward
		forward := currentPosition.DirectionAdd(currentDirection)
		if forward.IsValidMove(self, opponent) {
			return forward
		} else {
			right := currentPosition.DirectionAdd(currentDirection.TurnRight())
			left := currentPosition.DirectionAdd(currentDirection.TurnLeft())
			rightValid := right.IsValidMove(self, opponent)
			leftValid := left.IsValidMove(self, opponent)
			if (rightValid && leftValid) {
				if (right.DistanceTo(opponentPosition) < left.DistanceTo(opponentPosition)) {
					return right
				}
				return left
			}
			if (rightValid) {
				return right
			}
			return left
		}
	},
}
