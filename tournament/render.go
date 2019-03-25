package tournament

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Render renders the tournament
func (t *Tournament) Render(renderer *sdl.Renderer) {
	// render active match
	t.activeMatch.Render(renderer)

}
