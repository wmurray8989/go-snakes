package match

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

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
