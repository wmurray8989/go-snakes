package snakes

import (
	"github.com/wmurray8989/go-snakes/direction"
	"github.com/wmurray8989/go-snakes/position"
)

// Panic10 moves up then panics on the 10th move
func Panic10(self []position.Position, opponent []position.Position) position.Position {
	currentPosition := self[len(self)-1]

	if len(self) >= 10 {
		panic("I'm Panicking!!!")
	}

	nextMove := currentPosition.DirectionAdd(direction.Up)

	return nextMove
}
