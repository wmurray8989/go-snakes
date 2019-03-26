package assets

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// DrawText draws text to the renderer
func DrawText(renderer *sdl.Renderer, font *ttf.Font, text string, X int32, Y int32, color sdl.Color) {
	surface, _ := font.RenderUTF8Solid(text, color)
	texture, _ := renderer.CreateTextureFromSurface(surface)
	renderer.Copy(texture, &sdl.Rect{
		X: 0,
		Y: 0,
		W: surface.W,
		H: surface.H,
	}, &sdl.Rect{
		X: X,
		Y: Y,
		W: surface.W,
		H: surface.H,
	})
}
