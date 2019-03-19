package main

import (
	"fmt"
	"os"
	"strconv"

	"bitbucket.org/wmurray8989/go-snakes/simulation"
	"bitbucket.org/wmurray8989/go-snakes/snakes"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle = "Go Snakes"
var winWidth, winHeight int32 = 1920, 1080

var color = map[string]sdl.Color{
	"red":   sdl.Color{R: 255, G: 0, B: 0, A: 255},
	"green": sdl.Color{R: 0, G: 255, B: 0, A: 255},
	"blue":  sdl.Color{R: 0, G: 0, B: 255, A: 255},
	"black": sdl.Color{R: 0, G: 0, B: 0},
}

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize SDL: %s\n", err)
		return 1
	}
	defer sdl.Quit()

	if window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 2
	}
	defer window.Destroy()

	if renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 3 // don't use os.Exit(3); otherwise, previous deferred calls will never run
	}
	defer renderer.Destroy()

	var simulation = simulation.NewSimulation()

	fullscreen := false
	running := true
	var lastTime = sdl.GetTicks()
	for running {

		// Events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				switch t.Keysym.Sym {
				case sdl.K_q:
					running = false
				case sdl.K_ESCAPE:
					running = false
				case sdl.K_f:
					if t.State == sdl.PRESSED {
						if fullscreen {
							window.SetFullscreen(0)
						} else {
							window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
						}
						fullscreen = !fullscreen
					}
				}
			}
		}

		// Logic
		result := simulation.Update(snakes.StraightLine, snakes.StraightLine)
		if result != 0 {
			running = false
			break
		}

		// Render
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		simulation.Render(renderer)

		// FPS
		var nowTime = sdl.GetTicks()
		if nowTime != lastTime {
			gfx.StringColor(renderer, 16, 16, "FPS: "+strconv.FormatUint(uint64(1000/(nowTime-lastTime)), 10), color["green"])
		}
		lastTime = nowTime

		renderer.SetViewport(&sdl.Rect{
			X: 0,
			Y: 0,
			W: winWidth,
			H: winHeight,
		})

		renderer.Present()

		sdl.Delay(100)
	}

	return 0
}

func main() {
	os.Exit(run())
}
