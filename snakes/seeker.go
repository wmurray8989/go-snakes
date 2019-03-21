package snakes

import (
	"sort"

	"bitbucket.org/wmurray8989/go-snakes/position"
)

// Seeker moves directly towards opponent
func Seeker(self []position.Position, opponent []position.Position) position.Position {
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
}
