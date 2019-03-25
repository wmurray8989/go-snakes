package player

import (
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/position"
)

// Render renders the match
func (p *Player) Render(renderer *sdl.Renderer) {
	const sideLength = 50
	const cellSize = 20

	// invert y positions
	positions := append([]position.Position(nil), p.Moves...)
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
		p.ColorPrimary.R,
		p.ColorPrimary.G,
		p.ColorPrimary.B,
		p.ColorPrimary.A,
	)
	renderer.FillRects(rects)

	// draw snake spine
	renderer.SetDrawColor(
		p.ColorSecondary.R,
		p.ColorSecondary.G,
		p.ColorSecondary.B,
		p.ColorSecondary.A,
	)
	renderer.DrawLines(points)

	// draw snake head
	lastPosition := positions[positionLength-1]
	gfx.FilledCircleColor(
		renderer,
		int32(lastPosition.X)*cellSize+cellSize/2,
		int32(lastPosition.Y)*cellSize+cellSize/2,
		cellSize/2,
		p.ColorSecondary,
	)
}
