package assets

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/ttf"
)

// NewAssets creates a Assets
func NewAssets() (*Assets, error) {
	var assets = &Assets{}

	var font *ttf.Font

	font, err := ttf.OpenFont("assets/fonts/junegull.ttf", 36)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		return assets, err
	}
	assets.Font36 = font

	font, err = ttf.OpenFont("assets/fonts/junegull.ttf", 30)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		return assets, err
	}
	assets.Font30 = font

	return assets, nil
}
