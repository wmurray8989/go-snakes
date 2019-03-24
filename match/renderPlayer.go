package match

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/position"
)

func renderPlayer(renderer *sdl.Renderer, uninvertedPositions []position.Position, color sdl.Color, cellSize int32) {
	// invert y positions
	positions := append([]position.Position(nil), uninvertedPositions...)
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
