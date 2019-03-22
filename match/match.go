package match

import (
	"bitbucket.org/wmurray8989/go-snakes/position"
	"github.com/veandco/go-sdl2/gfx"
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

// Strategy is a function that takes a history of your positions and your opponents positions and returns a next position
type Strategy func(self []position.Position, opponent []position.Position) position.Position

// NewMatch creates a Match
func NewMatch(player1 Strategy, player2 Strategy) Match {
	var match = Match{}

	// setup strategies
	match.player1 = player1
	match.player2 = player2

	// setup starting positions
	match.player1History = append(match.player1History, position.Position{X: 24, Y: 24})
	match.player2History = append(match.player2History, position.Position{X: 26, Y: 26})

	// setup colors
	match.p1Color.R = 255
	match.p1Color.A = 255

	match.p2Color.G = 255
	match.p2Color.A = 255

	match.gridColor.R = 100
	match.gridColor.G = 100
	match.gridColor.B = 100
	match.gridColor.A = 100

	return match
}

// Update runs the match
func (p *Match) Update() {
	if p.status != matchRunning {
		return
	}

	p1Move := p.player1(p.player1History, p.player2History)
	if !(p1Move.IsValidMove(p.player1History, p.player2History)) {
		p.status = snake2Wins
	}
	p.player1History = append(p.player1History, p1Move)

	p2Move := p.player2(p.player2History, p.player1History)
	if !(p2Move.IsValidMove(p.player2History, p.player1History)) {
		p.status = snake1Wins
	}
	p.player2History = append(p.player2History, p2Move)
}

func renderPlayer(renderer *sdl.Renderer, positions []position.Position, color sdl.Color, cellSize int32) {
	lastIndex := len(positions) - 1
	for index, position := range positions {
		if index == lastIndex {
			gfx.FilledCircleColor(
				renderer,
				int32(position.X)*cellSize+cellSize/2,
				int32(position.Y)*cellSize+cellSize/2,
				cellSize/2,
				color,
			)
			continue
		}
		renderer.SetDrawColor(
			color.R,
			color.G,
			color.B,
			color.A,
		)
		renderer.FillRect(
			&sdl.Rect{
				X: int32(position.X) * cellSize,
				Y: int32(position.Y) * cellSize,
				W: cellSize,
				H: cellSize,
			},
		)
	}
}

// Render renders the particle system
func (p *Match) Render(renderer *sdl.Renderer) {

	const sideLength = 50
	const cellSize = 20

	// draw players
	renderPlayer(renderer, p.player1History, p.p1Color, cellSize)
	renderPlayer(renderer, p.player2History, p.p2Color, cellSize)

	// draw grid
	for x := int32(0); x < sideLength*cellSize; x = x + cellSize {
		for y := int32(0); y < sideLength*cellSize; y = y + cellSize {
			gfx.RectangleColor(
				renderer,
				int32(x),
				int32(y),
				int32(x+cellSize),
				int32(y+cellSize),
				p.gridColor,
			)
		}
	}

	// Draw Winning text
	if p.status == snake1Wins {
		gfx.StringColor(renderer, 100, 100, "Snake 1 Wins", p.p1Color)
	}
	if p.status == snake2Wins {
		gfx.StringColor(renderer, 100, 100, "Snake 2 Wins", p.p2Color)
	}

}
