package player

import "github.com/veandco/go-sdl2/sdl"

// NewPlayer creates a Player
func NewPlayer(strategy Strategy, colorPrimary sdl.Color, colorSecondary sdl.Color) Player {
	return Player{
		strategy:       strategy,
		colorPrimary:   colorPrimary,
		colorSecondary: colorSecondary,
	}
}
