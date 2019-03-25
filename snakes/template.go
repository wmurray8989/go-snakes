package snakes

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/direction"
	"github.com/wmurray8989/go-snakes/player"
	"github.com/wmurray8989/go-snakes/position"
)

// TemplateSnake moves up
var TemplateSnake = player.Player{
	Name:           "Template Snake",
	Author:         "",
	ColorPrimary:   sdl.Color{R: 255, G: 255, B: 255, A: 255},
	ColorSecondary: sdl.Color{R: 0, G: 0, B: 0, A: 255},
	Strategy: func(self []position.Position, opponent []position.Position) position.Position {
		currentPosition := self[len(self)-1]

		nextMove := currentPosition.DirectionAdd(direction.Up)

		return nextMove
	},
}
