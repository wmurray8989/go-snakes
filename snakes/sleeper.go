package snakes

import (
	"math/rand"
	"time"

	"github.com/wmurray8989/go-snakes/direction"
	"github.com/wmurray8989/go-snakes/position"
)

// Sleeper moves up and can randomly sleep
func Sleeper(self []position.Position, opponent []position.Position) position.Position {
	currentPosition := self[len(self)-1]

	rand.Seed(time.Now().UTC().UnixNano())
	randomDigit := rand.Intn(10)

	if randomDigit == 9 {
		time.Sleep(time.Minute)
	}

	nextMove := currentPosition.DirectionAdd(direction.Up)

	return nextMove
}
