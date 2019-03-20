package snakes

import (
	"math/rand"
	"time"

	"bitbucket.org/wmurray8989/go-snakes/simulation"
)

func includesPosition(elements []simulation.Position, p simulation.Position) bool {
	for _, element := range elements {
		if element.X == p.X && element.Y == p.Y {
			return true
		}
	}
	return false
}

func positionIsFree(self []simulation.Position, opponent []simulation.Position, p simulation.Position) bool {
	if includesPosition(self, p) || includesPosition(opponent, p) {
		return false
	}
	return true
}

func getAvailableMoves(self []simulation.Position, opponent []simulation.Position, currentPosition simulation.Position) []simulation.Position {
	availableMoves := []simulation.Position{}

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if positionIsFree(self, opponent, simulation.Position{
				X: currentPosition.X + x,
				Y: currentPosition.Y + y,
			}) {
				availableMoves = append(availableMoves, simulation.Position{
					X: currentPosition.X + x,
					Y: currentPosition.Y + y,
				})
			}
		}
	}

	return availableMoves
}

// Brownian moves randomly avoiding self collisions
func Brownian(self []simulation.Position, opponent []simulation.Position) simulation.Position {
	currentPosition := self[len(self)-1]

	availableMoves := getAvailableMoves(self, opponent, currentPosition)
	availableMovesCount := len(availableMoves)

	if availableMovesCount == 0 {
		// Give up
		return currentPosition
	}

	rand.Seed(time.Now().UTC().UnixNano())
	randomIndex := rand.Intn(availableMovesCount)

	nextMove := availableMoves[randomIndex]

	return nextMove
}
