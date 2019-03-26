package tournament

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/assets"
	"github.com/wmurray8989/go-snakes/player"
)

// Render renders the tournament
func (t *Tournament) Render(renderer *sdl.Renderer, globalAssets *assets.Assets) {
	// render active match
	t.activeMatch.Render(renderer, globalAssets)

	// render bracket
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

func renderBracketColumn(renderer *sdl.Renderer, globalAssets *assets.Assets, players []player.Player, count int, columnIndex int) {
	height := 1056 / count
	for index, player := range players {
		renderName(
			renderer,
			globalAssets,
			player,
			int32(1020+columnIndex*150),
			int32(height*index+12),
			150,
			int32(height),
		)
	}
}

func renderName(renderer *sdl.Renderer, globalAssets *assets.Assets, player player.Player, X int32, Y int32, W int32, H int32) {
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
	renderer.SetDrawColor(
		100,
		100,
		100,
		100,
	)
	renderer.DrawRect(&sdl.Rect{
		X: X,
		Y: Y,
		W: W,
		H: H,
	})

	if player.Name != "" {
		assets.DrawText(
			renderer,
			globalAssets.Font30,
			fmt.Sprintf("%s", player.Name),
			X+2,
			Y,
			player.ColorSecondary,
		)
	}
}
