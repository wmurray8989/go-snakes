package match

import (
	"github.com/wmurray8989/go-snakes/position"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

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
