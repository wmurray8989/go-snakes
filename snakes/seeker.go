package snakes

import (
	"sort"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/player"
	"github.com/wmurray8989/go-snakes/position"
)

// Seeker moves directly towards opponent
var Seeker = player.Player{
	Name:           "Seeker",
	Author:         "Null",
	ColorPrimary:   sdl.Color{R: 0, G: 255, B: 0, A: 255},
	ColorSecondary: sdl.Color{R: 255, G: 0, B: 255, A: 255},
	Strategy: func(self []position.Position, opponent []position.Position) position.Position {
		currentPosition := self[len(self)-1]
		opponentPosition := opponent[len(opponent)-1]

		validMoves := currentPosition.GetValidMoves(self, opponent)
		validMovesCount := len(validMoves)
		if validMovesCount == 0 {
			// Give up
			return currentPosition
		}

		// sort moves by distance from opponent
		sort.Slice(validMoves, func(i, j int) bool {
			return validMoves[i].DistanceTo(opponentPosition) < validMoves[j].DistanceTo(opponentPosition)
		})

		return validMoves[0]
	},
}
