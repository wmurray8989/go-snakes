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

	// Draw Points
	white := sdl.Color{R: 255, B: 255, G: 255, A: 255}
	gfx.StringColor(renderer, 200, 550, fmt.Sprintf("Snake 1: %d", m.player1Points), white)
	gfx.StringColor(renderer, 200, 560, fmt.Sprintf("Snake 2: %d", m.player2Points), white)

	gfx.StringColor(renderer, 200, 575, fmt.Sprintf("Time remaining: %f", m.timeRemaining.Seconds()), white)

}
