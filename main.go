package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/wmurray8989/go-snakes/player"
	"github.com/wmurray8989/go-snakes/snakes"
	"github.com/wmurray8989/go-snakes/tournament"
)

var winTitle = "Go Snakes"
var winWidth, winHeight int32 = 1920, 1080

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

	// renderer.SetLogicalSize(winWidth, winHeight)

	var activeTournament = tournament.NewTournament([32]player.Player{
		snakes.SpiralIn,
		snakes.SpiralOut,
		snakes.Random,
		snakes.Panicker,
		snakes.Seeker,
		snakes.Random,
		snakes.Sleeper,
		snakes.Random,
		snakes.SpiralIn,
		snakes.SpiralOut,
		snakes.Random,
		snakes.Panicker,
		snakes.Seeker,
		snakes.Random,
		snakes.Sleeper,
		snakes.Random,
		snakes.SpiralIn,
		snakes.SpiralOut,
		snakes.Random,
		snakes.Panicker,
		snakes.Seeker,
		snakes.Random,
		snakes.Sleeper,
		snakes.Random,
		snakes.SpiralIn,
		snakes.SpiralOut,
		snakes.Random,
		snakes.Panicker,
		snakes.Seeker,
		snakes.Random,
		snakes.Sleeper,
		snakes.Random,
	})

	fullscreen := false
	running := true
	lastTime := time.Time{}
	ticksPerSecond := 1000
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
				// case sdl.K_r:
				// 	if t.State == sdl.PRESSED {
				// 		activeMatch = match.NewMatch(snakes.Panic10, snakes.SpiralOut)
				// 		lastTime = time.Time{}
				// 	}
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

		if time.Since(lastTime) > (time.Second / time.Duration(ticksPerSecond)) {
			// Logic
			activeTournament.Update()

			// Render
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
			activeTournament.Render(renderer)
			renderer.Present()

			lastTime = time.Now().UTC()
		}

		renderer.SetViewport(&sdl.Rect{
			X: 0,
			Y: 0,
			W: winWidth,
			H: winHeight,
		})
	}

	return 0
}

func main() {
	os.Exit(run())
}
