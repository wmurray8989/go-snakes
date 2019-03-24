package player

import "github.com/veandco/go-sdl2/sdl"

// GetColorSecondary returns the player's secondary color
func (p *Player) GetColorSecondary() sdl.Color {
	return p.colorSecondary
}
