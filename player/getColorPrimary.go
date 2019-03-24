package player

import "github.com/veandco/go-sdl2/sdl"

// GetColorPrimary returns the player's primary color
func (p *Player) GetColorPrimary() sdl.Color {
	return p.colorPrimary
}
