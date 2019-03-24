package player

import (
	"log"
	"time"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/position"
)

// Render renders the match
func (p *Player) Render(renderer *sdl.Renderer) {
	start := time.Now()

	const sideLength = 50
	const cellSize = 10

	// invert y positions
	positions := append([]position.Position(nil), p.moves...)
	for i, position := range positions {
		positions[i].Y = 49 - position.Y
	}

	lastIndex := len(positions) - 1
	for index, position := range positions {
		if index == lastIndex {
			gfx.FilledCircleColor(
				renderer,
				int32(position.X)*cellSize+cellSize/2,
				int32(position.Y)*cellSize+cellSize/2,
				cellSize/2,
				p.colorPrimary,
			)
			continue
		}
		renderer.SetDrawColor(
			p.colorPrimary.R,
			p.colorPrimary.G,
			p.colorPrimary.B,
			p.colorPrimary.A,
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

	elapsed := time.Since(start)
	log.Printf("Render took %s", elapsed)
}
