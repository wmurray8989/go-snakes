package snakes

import (
	"github.com/wmurray8989/go-snakes/position"
)

// SpiralIn moves to the outside and then spirals inward
func SpiralIn(self []position.Position, opponent []position.Position) position.Position {
	currentPosition := self[len(self)-1]
	
	// If this is the first move, go up
	if (len(self) == 1) {
		return position.Position{
			X: currentPosition.X,
			Y: currentPosition.Y + 1,
		}
	}
	
	// get current direction 
	previousPosition := self[len(self) - 2]
	currentDirection := previousPosition.DirectionTo(currentPosition)

	// Try moving forward
	forward := currentPosition.DirectionAdd(currentDirection)
	if (forward.IsValidMove(self, opponent)) {
		return forward
	}

	// try turning right
	turnRight := currentPosition.DirectionAdd(currentDirection.TurnRight())
	if (turnRight.IsValidMove(self, opponent)) {
		return turnRight
	}

	// try turning left
	turnLeft := currentPosition.DirectionAdd(currentDirection.TurnLeft())
	if (turnLeft.IsValidMove(self, opponent)) {
		return turnLeft
	}	

	// Give up
	return currentPosition
}
