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
}

func renderBracketColumn(renderer *sdl.Renderer, globalAssets *assets.Assets, players []player.Player, count int, columnIndex int) {
	height := 1030 / count
	for index, player := range players {
		renderName(
			renderer,
			globalAssets,
			player,
			int32(1000+columnIndex*153),
			int32(height*index),
			153,
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

	// assets.DrawText(
	// 	renderer,
	// 	globalAssets.Font36,
	// 	fmt.Sprintf("%s", player.Name),
	// 	X+10,
	// 	Y+10,
	// 	player.ColorSecondary,
	// )

	println(player.Name)
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

	// gfx.StringColor(renderer, X+10, int32(10+Y), fmt.Sprintf("%s", player.Name), player.ColorSecondary)
}
