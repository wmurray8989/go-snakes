package snakes

import (
	"github.com/wmurray8989/go-snakes/direction"
	"github.com/wmurray8989/go-snakes/position"
)

// SpiralIn moves to the outside and then spirals inward
func SpiralIn(self []position.Position, opponent []position.Position) position.Position {
	currentPosition := self[len(self)-1]

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
	}

	// try turning right
	turnRight := currentPosition.DirectionAdd(currentDirection.TurnRight())
	if turnRight.IsValidMove(self, opponent) {
		return turnRight
	}

	// try turning left
	turnLeft := currentPosition.DirectionAdd(currentDirection.TurnLeft())
	if turnLeft.IsValidMove(self, opponent) {
		return turnLeft
	}

	// Give up
	return currentPosition
}
