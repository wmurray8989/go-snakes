package tournament

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/assets"
	"github.com/wmurray8989/go-snakes/player"
)

// renderBracketColumn renders one column of the bracket
func renderBracketColumn(renderer *sdl.Renderer, globalAssets *assets.Assets, players []player.Player, count int, columnIndex int) {
	height := 1056 / count
	for index, player := range players {
		renderBracketName(
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
