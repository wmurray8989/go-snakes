package match

import (
	"github.com/wmurray8989/go-snakes/position"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	matchRunning = 0
	snake1Wins   = 1
	snake2Wins   = 2
)

// Match simulates a match
type Match struct {
	status         int
	player1        Strategy
	player2        Strategy
	player1History []position.Position
	player2History []position.Position
	p1Color        sdl.Color
	p2Color        sdl.Color
	gridColor      sdl.Color
}
