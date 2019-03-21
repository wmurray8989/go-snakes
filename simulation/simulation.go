package simulation

import (
	"bitbucket.org/wmurray8989/go-snakes/position"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	simulationRunning = 0
	snake1Wins        = 1
	snake2Wins        = 2
)

// Simulation simulates a simulation
type Simulation struct {
	status         int
	player1History []position.Position
	player2History []position.Position
	p1Color        sdl.Color
	p2Color        sdl.Color
	gridColor      sdl.Color
}

// NewSimulation creates a Simulation
func NewSimulation() Simulation {
	var simulation = Simulation{}

	// setup starting positions
	simulation.player1History = append(simulation.player1History, position.Position{24, 24})
	simulation.player2History = append(simulation.player2History, position.Position{26, 26})

	// setup colors
	simulation.p1Color.R = 255
	simulation.p1Color.A = 255

	simulation.p2Color.G = 255
	simulation.p2Color.A = 255

	simulation.gridColor.R = 100
	simulation.gridColor.G = 100
	simulation.gridColor.B = 100
	simulation.gridColor.A = 100

	return simulation
}

func moveIsValid(move position.Position, self []position.Position, opponent []position.Position) bool {
	// Is within bounds
	if !move.IsWithinBounds() {
		return false
	}

	// Distance from last position is 1
	lastPosition := self[len(self)-1]
	if move.DistanceTo(lastPosition) != 1 {
		return false
	}

	// Is unoccupied
	if !move.IsUnoccupied(self, opponent) {
		return false
	}
	return true
}

// Strategy is a function that takes a history of your positions and your opponents positions and returns a next position
type Strategy func(self []position.Position, opponent []position.Position) position.Position

// Update runs the simulation
func (p *Simulation) Update(player1 Strategy, player2 Strategy) {
	if p.status != simulationRunning {
		return
	}

	p1Move := player1(p.player1History, p.player2History)
	if !(moveIsValid(p1Move, p.player1History, p.player2History)) {
		p.status = snake2Wins
	}
	p.player1History = append(p.player1History, p1Move)

	p2Move := player2(p.player2History, p.player1History)
	if !(moveIsValid(p2Move, p.player2History, p.player1History)) {
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
func (p *Simulation) Render(renderer *sdl.Renderer) {

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
}
