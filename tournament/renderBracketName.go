package tournament

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/assets"
	"github.com/wmurray8989/go-snakes/player"
)

func renderBracketName(renderer *sdl.Renderer, globalAssets *assets.Assets, player player.Player, X int32, Y int32, W int32, H int32) {
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
