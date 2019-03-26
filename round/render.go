package round

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/assets"
)

// Render renders the round
func (p *Round) Render(renderer *sdl.Renderer, globalAssets *assets.Assets) {
	const sideLength = 50
	const cellSize = 20

	// draw players
	p.player1.Render(renderer, globalAssets)
	p.player2.Render(renderer, globalAssets)

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
