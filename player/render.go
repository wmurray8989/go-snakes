package player

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/position"
)

// Render renders the match
func (p *Player) Render(renderer *sdl.Renderer) {
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
				p.color,
			)
			continue
		}
		renderer.SetDrawColor(
			p.color.R,
			p.color.G,
			p.color.B,
			p.color.A,
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
