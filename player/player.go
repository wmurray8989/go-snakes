package player

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/position"
)

// Player is a participant in a match
type Player struct {
	Name           string
	Author         string
	Strategy       Strategy
	Moves          []position.Position
	ColorPrimary   sdl.Color
	ColorSecondary sdl.Color
}
