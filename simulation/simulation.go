package simulation

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	SimulationRunning = 0
	Snake1Wins        = 1
	Snake2Wins        = 2
	SimulationDraw    = 3
)

// Simulation simulates a simulation
type Simulation struct {
	player1History []Position
	player2History []Position
}

// NewSimulation creates a Simulation
func NewSimulation() Simulation {
	var simulation = Simulation{}

	simulation.player1History = append(simulation.player1History, Position{10, 10})
	simulation.player2History = append(simulation.player2History, Position{40, 40})

	return simulation
}

// Position contains an X and Y coordinate
type Position struct {
	X int
	Y int
}

// Strategy is a function that takes a history of your positions and your opponents positions and returns a next position
type Strategy func(self []Position, opponent []Position) Position

// Update runs the simulation
func (p *Simulation) Update(player1 Strategy, player2 Strategy) int {
	p1Loss := false
	p2Loss := false

	p1Next := player1(p.player1History, p.player2History)
	p2Next := player2(p.player2History, p.player1History)

	if p1Next.X < 0 || p1Next.X > 50 || p1Next.Y < 0 || p1Next.Y > 50 {
		p1Loss = true
	}

	if p2Next.X < 0 || p2Next.X > 50 || p2Next.Y < 0 || p2Next.Y > 50 {
		p2Loss = true
	}

	p.player1History = append(p.player1History, p1Next)
	p.player2History = append(p.player2History, p2Next)

	if p1Loss && p2Loss {
		return SimulationDraw
	}

	if p1Loss {
		return Snake2Wins
	}

	if p2Loss {
		return Snake1Wins
	}

	return SimulationRunning
}

// Render renders the particle system
func (p *Simulation) Render(renderer *sdl.Renderer) {

	// draw grid
	const sideLength = 50
	const cellSize = 20
	for x := int32(0); x < sideLength*cellSize; x = x + cellSize {
		for y := int32(0); y < sideLength*cellSize; y = y + cellSize {
			gfx.RectangleColor(
				renderer,
				int32(x),
				int32(y),
				int32(x+cellSize),
				int32(y+cellSize),
				sdl.Color{R: 50, G: 50, B: 50, A: 255},
			)
		}
	}

	// draw player 1
	for _, position := range p.player1History {
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.FillRect(&sdl.Rect{X: int32(position.X) * cellSize, Y: int32(position.Y) * cellSize, W: cellSize, H: cellSize})
	}

	// draw player 2
	for _, position := range p.player2History {
		renderer.SetDrawColor(0, 0, 255, 255)
		renderer.FillRect(&sdl.Rect{X: int32(position.X) * cellSize, Y: int32(position.Y) * cellSize, W: cellSize, H: cellSize})
	}
}
