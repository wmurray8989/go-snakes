package snakes

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/player"
	"github.com/wmurray8989/go-snakes/position"
)

// Random moves randomly avoiding self collisions
var Random = player.Player{
	Name:           "Random",
	Author:         "Null",
	ColorPrimary:   sdl.Color{R: 255, G: 0, B: 0, A: 255},
	ColorSecondary: sdl.Color{R: 255, G: 255, B: 255, A: 255},
	Strategy: func(self []position.Position, opponent []position.Position) position.Position {
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
	},
}
