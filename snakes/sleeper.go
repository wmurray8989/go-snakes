package snakes

import (
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/direction"
	"github.com/wmurray8989/go-snakes/player"
	"github.com/wmurray8989/go-snakes/position"
)

// Sleeper moves up and can randomly sleep
var Sleeper = player.Player{
	Name:           "Sleeper",
	Author:         "Null",
	ColorPrimary:   sdl.Color{R: 0, G: 0, B: 255, A: 255},
	ColorSecondary: sdl.Color{R: 0, G: 0, B: 0, A: 255},
	Strategy: func(self []position.Position, opponent []position.Position) position.Position {
		currentPosition := self[len(self)-1]

		rand.Seed(time.Now().UTC().UnixNano())
		randomDigit := rand.Intn(10)

		if randomDigit == 9 {
			time.Sleep(time.Minute)
		}

		nextMove := currentPosition.DirectionAdd(direction.Up)

		return nextMove
	},
}
