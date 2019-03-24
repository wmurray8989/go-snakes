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
	p1PrimaryColor := m.player1.GetColorPrimary()
	renderer.SetDrawColor(
		p1PrimaryColor.R,
		p1PrimaryColor.G,
		p1PrimaryColor.B,
		p1PrimaryColor.A,
	)
	renderer.FillRect(&sdl.Rect{
		X: 0,
		Y: 500,
		W: 250,
		H: 30,
	})
	gfx.StringColor(renderer, 10, 510, fmt.Sprintf("%s: %d", m.player2.GetName(), m.player2Points), m.player2.GetColorSecondary())

	// Draw Points player 1
	p2PrimaryColor := m.player2.GetColorPrimary()
	renderer.SetDrawColor(
		p2PrimaryColor.R,
		p2PrimaryColor.G,
		p2PrimaryColor.B,
		p2PrimaryColor.A,
	)
	renderer.FillRect(&sdl.Rect{
		X: 250,
		Y: 500,
		W: 250,
		H: 30,
	})
	gfx.StringColor(renderer, 260, 510, fmt.Sprintf("%s: %d", m.player2.GetName(), m.player2Points), m.player2.GetColorSecondary())

	// Draw time remaining
	white := sdl.Color{R: 255, B: 255, G: 255, A: 255}
	gfx.StringColor(renderer, 200, 575, fmt.Sprintf("Time remaining: %f", m.timeRemaining.Seconds()), white)

}
