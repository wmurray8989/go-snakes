package snakes

import (
	"math/rand"
	"time"

	"github.com/wmurray8989/go-snakes/position"
)

// Random moves randomly avoiding self collisions
func Random(self []position.Position, opponent []position.Position) position.Position {
	currentPosition := self[len(self)-1]

	validMoves := currentPosition.GetValidMoves(self, opponent)
	validMovesCount := len(validMoves)
	if validMovesCount == 0 {
		// Give up
		return currentPosition
	}

	rand.Seed(time.Now().UTC().UnixNano())
	randomIndex := rand.Intn(validMovesCount)

	nextMove := validMoves[randomIndex]

	return nextMove
}
