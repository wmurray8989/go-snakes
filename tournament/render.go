package tournament

import (
	"fmt"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/player"
)

// Render renders the tournament
func (t *Tournament) Render(renderer *sdl.Renderer) {
	// render active match
	t.activeMatch.Render(renderer)

	// render bracket
	renderBracketColumn(renderer, t.series32[0:len(t.series32)], 32, 0)
	renderBracketColumn(renderer, t.series16[0:len(t.series16)], 16, 1)
	renderBracketColumn(renderer, t.series8[0:len(t.series8)], 8, 2)
	renderBracketColumn(renderer, t.series4[0:len(t.series4)], 4, 3)
	renderBracketColumn(renderer, t.series2[0:len(t.series2)], 2, 4)
}

func renderBracketColumn(renderer *sdl.Renderer, players []player.Player, count int, columnIndex int) {
	height := 1030 / count
	for index, player := range players {
		renderName(
			renderer,
			player,
			int32(1000+columnIndex*153),
			int32(height*index),
			153,
			int32(height),
		)
	}
}

func renderName(renderer *sdl.Renderer, player player.Player, X int32, Y int32, W int32, H int32) {
	renderer.SetDrawColor(
		player.ColorPrimary.R,
		player.ColorPrimary.G,
		player.ColorPrimary.B,
		player.ColorPrimary.A,
	)
	renderer.FillRect(&sdl.Rect{
		X: X,
		Y: Y,
		W: W,
		H: H,
	})
	gfx.StringColor(renderer, 1010, int32(10+Y), fmt.Sprintf("%s", player.Name), player.ColorSecondary)
}
