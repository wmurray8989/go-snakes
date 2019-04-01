package snakes

import (
	"sort"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/direction"
	"github.com/wmurray8989/go-snakes/player"
	"github.com/wmurray8989/go-snakes/position"
)

// Scaredy will maximize distance from opponent, minimize distance from the nearest corner, and avoid obvious hellholes
var Scaredy = player.Player{
	Name:           "Scaredy",
	Author:         "Molly Henderson",
	ColorPrimary:   sdl.Color{R: 0, G: 126, B: 181, A: 255},
	ColorSecondary: sdl.Color{R: 255, G: 114, B: 0, A: 255},
	Strategy: func(self []position.Position, opponent []position.Position) position.Position {
		currentPosition := self[len(self)-1]
		opponentPosition := opponent[len(opponent)-1]

		validMoves := currentPosition.GetValidMoves(self, opponent)
		validMovesCount := len(validMoves)
		if validMovesCount == 0 {
			// Give up
			return currentPosition
		}

		currentDirection := direction.Null
		if len(self) > 1 {
			currentDirection = self[len(self)-2].DirectionTo(currentPosition)
		}

		// sort moves by our heuristic
		sort.Slice(validMoves, func(i, j int) bool {
			iPos := validMoves[i]
			jPos := validMoves[j]

			nearestX := 0
			nearestY := 0
			if currentPosition.X >= 25 {
				nearestX = 49
			}
			if currentPosition.Y >= 25 {
				nearestY = 49
			}

			nearestCorner := position.Position{X: nearestX, Y: nearestY}

			// high penalty for entering a single-wide path
			baseI := 0
			baseJ := 0
			if (currentDirection == direction.Up || currentDirection == direction.Down) &&
				(!iPos.DirectionAdd(direction.Left).IsUnoccupied(self, opponent) &&
					!iPos.DirectionAdd(direction.Right).IsUnoccupied(self, opponent)) {
				baseI += 100000
			}
			if (currentDirection == direction.Left || currentDirection == direction.Right) &&
				(!iPos.DirectionAdd(direction.Up).IsUnoccupied(self, opponent) &&
					!iPos.DirectionAdd(direction.Down).IsUnoccupied(self, opponent)) {
				baseI += 100000
			}

			if (currentDirection == direction.Up || currentDirection == direction.Down) &&
				(!jPos.DirectionAdd(direction.Left).IsUnoccupied(self, opponent) &&
					!jPos.DirectionAdd(direction.Right).IsUnoccupied(self, opponent)) {
				baseJ += 100000
			}
			if (currentDirection == direction.Left || currentDirection == direction.Right) &&
				(!jPos.DirectionAdd(direction.Up).IsUnoccupied(self, opponent) &&
					!jPos.DirectionAdd(direction.Down).IsUnoccupied(self, opponent)) {
				baseJ += 100000
			}

			iHeuristic := iPos.DistanceTo(opponentPosition) - 2*iPos.DistanceTo(nearestCorner) - baseI
			jHeuristic := jPos.DistanceTo(opponentPosition) - 2*jPos.DistanceTo(nearestCorner) - baseJ

			// sort in desc order by heuristic value
			return iHeuristic > jHeuristic
		})

		return validMoves[0]
	},
}
