package tournament

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/assets"
)

// Render renders the tournament
func (t *Tournament) Render(renderer *sdl.Renderer, globalAssets *assets.Assets) {
	// render active match

	t.activeMatch.Render(renderer, globalAssets)

	// render bracket
	renderBracket(renderer, globalAssets, t)
}
