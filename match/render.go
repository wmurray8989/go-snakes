package match

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/assets"
)

// Render renders the match
func (m *Match) Render(renderer *sdl.Renderer, globalAssets *assets.Assets) {

	// render active round
	m.activeRound.Render(renderer, globalAssets)

	// Draw Points player 1
	renderer.SetDrawColor(
		m.player1.ColorPrimary.R,
		m.player1.ColorPrimary.G,
		m.player1.ColorPrimary.B,
		m.player1.ColorPrimary.A,
	)
	renderer.FillRect(&sdl.Rect{
		X: 0,
		Y: 1000,
		W: 500,
		H: 30,
	})
	assets.DrawText(
		renderer,
		globalAssets.Font30,
		fmt.Sprintf("%s", m.player1.Name),
		10,
		1000,
		m.player1.ColorSecondary,
	)
	assets.DrawText(
		renderer,
		globalAssets.Font30,
		fmt.Sprintf("%d", m.player1Points),
		350,
		1000,
		m.player1.ColorSecondary,
	)

	// Draw Points player 2
	renderer.SetDrawColor(
		m.player2.ColorPrimary.R,
		m.player2.ColorPrimary.G,
		m.player2.ColorPrimary.B,
		m.player2.ColorPrimary.A,
	)
	renderer.FillRect(&sdl.Rect{
		X: 500,
		Y: 1000,
		W: 500,
		H: 30,
	})
	assets.DrawText(
		renderer,
		globalAssets.Font30,
		fmt.Sprintf("%s", m.player2.Name),
		510,
		1000,
		m.player2.ColorSecondary,
	)
	assets.DrawText(
		renderer,
		globalAssets.Font30,
		fmt.Sprintf("%d", m.player2Points),
		850,
		1000,
		m.player2.ColorSecondary,
	)

	// Draw time remaining
	renderer.SetDrawColor(
		200,
		200,
		200,
		255,
	)
	renderer.FillRect(&sdl.Rect{
		X: 0,
		Y: 1030,
		W: 1920,
		H: 50,
	})

	assets.DrawText(
		renderer,
		globalAssets.Font50,
		fmt.Sprintf("Time remaining: %.1f", m.timeRemaining.Seconds()),
		250,
		1030,
		sdl.Color{
			R: 0,
			B: 0,
			G: 0,
			A: 255,
		},
	)
}
