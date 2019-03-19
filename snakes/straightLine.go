package snakes

import "bitbucket.org/wmurray8989/go-snakes/simulation"

// StraightLine goes in a straight line
func StraightLine(self []simulation.Position, opponent []simulation.Position) simulation.Position {
	currentPosition := self[len(self)-1]

	nextPosition := simulation.Position{
		X: currentPosition.X + 1,
		Y: currentPosition.Y + 1,
	}

	return nextPosition
}
