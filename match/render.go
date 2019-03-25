package match

import (
	"fmt"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// Render renders the match
func (m *Match) Render(renderer *sdl.Renderer) {

	// render active round
	m.activeRound.Render(renderer)

	// Draw Points player 1
	renderer.SetDrawColor(
		m.player1.ColorPrimary.R,
		m.player1.ColorPrimary.G,
		m.player1.ColorPrimary.B,
		m.player1.ColorPrimary.A,
	)
	renderer.FillRect(&sdl.Rect{
		X: 0,
		Y: 500,
		W: 250,
		H: 30,
	})
	gfx.StringColor(renderer, 10, 510, fmt.Sprintf("%s: %d", m.player1.Name, m.player1Points), m.player1.ColorSecondary)

	// Draw Points player 2
	renderer.SetDrawColor(
		m.player2.ColorPrimary.R,
		m.player2.ColorPrimary.G,
		m.player2.ColorPrimary.B,
		m.player2.ColorPrimary.A,
	)
	renderer.FillRect(&sdl.Rect{
		X: 250,
		Y: 500,
		W: 250,
		H: 30,
	})
	gfx.StringColor(renderer, 260, 510, fmt.Sprintf("%s: %d", m.player2.Name, m.player2Points), m.player2.ColorSecondary)

	// Draw time remaining
	white := sdl.Color{R: 255, B: 255, G: 255, A: 255}
	gfx.StringColor(renderer, 200, 575, fmt.Sprintf("Time remaining: %f", m.timeRemaining.Seconds()), white)

}
