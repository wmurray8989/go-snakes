package simulation

import (
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
	player1History []Position
	player2History []Position
	p1Color        sdl.Color
	p2Color        sdl.Color
	gridColor      sdl.Color
}

// NewSimulation creates a Simulation
func NewSimulation() Simulation {
	var simulation = Simulation{}

	// setup starting positions
	simulation.player1History = append(simulation.player1History, Position{24, 24})
	simulation.player2History = append(simulation.player2History, Position{26, 26})

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

// Position contains an X and Y coordinate
type Position struct {
	X int
	Y int
}

func isWithinBounds(position Position) bool {
	return !(position.X < 0 || position.X > 50 || position.Y < 0 || position.Y > 50)
}

func integerAbs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func distance(a Position, b Position) int {
	return integerAbs(a.X-b.X) + integerAbs(a.Y-b.Y)
}

func includesPosition(elements []Position, p Position) bool {
	for _, element := range elements {
		if element.X == p.X && element.Y == p.Y {
			return true
		}
	}
	return false
}

func moveIsValid(move Position, player []Position, opponent []Position) bool {
	// Is within bounds
	if !isWithinBounds(move) {
		return false
	}

	// Distance from last position is 1
	lastPosition := player[len(player)-1]
	if distance(move, lastPosition) != 1 {
		return false
	}

	// Is unoccupied
	if includesPosition(player, move) || includesPosition(opponent, move) {
		return false
	}
	return true
}

// Strategy is a function that takes a history of your positions and your opponents positions and returns a next position
type Strategy func(self []Position, opponent []Position) Position

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

func renderPlayer(renderer *sdl.Renderer, positions []Position, color sdl.Color, cellSize int32) {
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
