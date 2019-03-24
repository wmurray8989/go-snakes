package round

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/position"
)

const (
	roundRunning = 0
	snake1Wins   = 1
	snake2Wins   = 2
)

// Round simulates a round
type Round struct {
	status         int
	playerTurn     bool
	player1        Strategy
	player2        Strategy
	player1History []position.Position
	player2History []position.Position
	p1Color        sdl.Color
	p2Color        sdl.Color
	gridColor      sdl.Color
}
