package snakes

import (
	"math/rand"
	"time"

	"github.com/wmurray8989/go-snakes/direction"
	"github.com/wmurray8989/go-snakes/position"
)

// Panicker moves up and can randomly panick
func Panicker(self []position.Position, opponent []position.Position) position.Position {
	currentPosition := self[len(self)-1]

	rand.Seed(time.Now().UTC().UnixNano())
	randomDigit := rand.Intn(10)

	if randomDigit == 9 {
		panic("I'm Panicking!!!")
	}

	nextMove := currentPosition.DirectionAdd(direction.Up)

	return nextMove
}
