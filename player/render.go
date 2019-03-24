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

	// position length
	positionLength := len(positions)

	// generate points and rects from positions
	rects := make([]sdl.Rect, positionLength)
	points := make([]sdl.Point, positionLength)
	for index, position := range positions {
		rects[index] = sdl.Rect{
			X: int32(position.X) * cellSize,
			Y: int32(position.Y) * cellSize,
			W: cellSize,
			H: cellSize,
		}
		points[index] = sdl.Point{
			X: (int32(position.X) * cellSize) + cellSize/2,
			Y: (int32(position.Y) * cellSize) + cellSize/2,
		}
	}

	// draw main body of snake
	renderer.SetDrawColor(
		p.colorPrimary.R,
		p.colorPrimary.G,
		p.colorPrimary.B,
		p.colorPrimary.A,
	)
	renderer.FillRects(rects)

	// draw snake spine
	renderer.SetDrawColor(
		p.colorSecondary.R,
		p.colorSecondary.G,
		p.colorSecondary.B,
		p.colorSecondary.A,
	)
	renderer.DrawLines(points)

	// draw snake head
	lastPosition := positions[positionLength-1]
	gfx.FilledCircleColor(
		renderer,
		int32(lastPosition.X)*cellSize+cellSize/2,
		int32(lastPosition.Y)*cellSize+cellSize/2,
		cellSize/2,
		p.colorSecondary,
	)
}
