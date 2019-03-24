package player

import "github.com/veandco/go-sdl2/sdl"

// NewPlayer creates a Player
func NewPlayer(name string, author string, strategy Strategy, colorPrimary sdl.Color, colorSecondary sdl.Color) Player {
	return Player{
		name:           name,
		author:         author,
		strategy:       strategy,
		colorPrimary:   colorPrimary,
		colorSecondary: colorSecondary,
	}
}
