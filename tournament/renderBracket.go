package tournament

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/assets"
	"github.com/wmurray8989/go-snakes/player"
)

// renderBracket renders the bracket
func renderBracket(renderer *sdl.Renderer, globalAssets *assets.Assets, t *Tournament) {
	renderBracketColumn(renderer, globalAssets, t.series32[0:len(t.series32)], 32, 0)
	renderBracketColumn(renderer, globalAssets, t.series16[0:len(t.series16)], 16, 1)
	renderBracketColumn(renderer, globalAssets, t.series8[0:len(t.series8)], 8, 2)
	renderBracketColumn(renderer, globalAssets, t.series4[0:len(t.series4)], 4, 3)
	renderBracketColumn(renderer, globalAssets, t.series2[0:len(t.series2)], 2, 4)
	renderBracketColumn(renderer, globalAssets, []player.Player{t.champion}, 1, 5)

	// Draw bracket boarders
	renderer.SetDrawColor(
		200,
		200,
		200,
		255,
	)
	renderer.FillRect(&sdl.Rect{
		X: 1000,
		Y: 0,
		W: 20,
		H: 1080,
	})
	renderer.FillRect(&sdl.Rect{
		X: 1000,
		Y: 0,
		W: 920,
		H: 12,
	})
	renderer.FillRect(&sdl.Rect{
		X: 1000,
		Y: 1068,
		W: 920,
		H: 12,
	})
}
