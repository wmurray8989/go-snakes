package player

import "github.com/veandco/go-sdl2/sdl"

// NewPlayer creates a Player
func NewPlayer(strategy Strategy, color sdl.Color) Player {
	return Player{
		strategy: strategy,
		color:    color,
	}
}
