package simulation

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	simulationRunning = 0
	snake1Wins        = 1
	snake2Wins        = 2
	simulationDraw    = 3
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

	simulation.p2Color.B = 255

	simulation.gridColor.R = 100
	simulation.gridColor.G = 100
	simulation.gridColor.B = 100

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
func (p *Simulation) Update(player1 Strategy, player2 Strategy) {
	if p.status != simulationRunning {
		return
	}

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
		p.status = simulationDraw
	}

	if p1Loss {
		p.status = snake2Wins
	}

	if p2Loss {
		p.status = snake1Wins
	}
}

func renderPlayer(renderer *sdl.Renderer, positions []Position, red uint8, green uint8, blue uint8, cellSize int32) {
	lastIndex := len(positions) - 1
	for index, position := range positions {
		print(position.X)
		print("-")
		println(position.Y)
		if index == lastIndex {
			gfx.FilledCircleRGBA(
				renderer,
				int32(position.X)*cellSize+cellSize/2,
				int32(position.Y)*cellSize+cellSize/2,
				cellSize/2,
				red,
				green,
				blue,
				255,
			)
			continue
		}
		renderer.SetDrawColor(
			red,
			green,
			blue,
			255,
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
	renderPlayer(renderer, p.player1History, 255, 0, 0, cellSize)
	renderPlayer(renderer, p.player2History, 0, 0, 255, cellSize)

	// draw grid
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
}
