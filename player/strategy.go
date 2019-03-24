package player

import "github.com/wmurray8989/go-snakes/position"

// Strategy is a function that takes a history of your positions and your opponents positions and returns a next position
type Strategy func(self []position.Position, opponent []position.Position) position.Position
