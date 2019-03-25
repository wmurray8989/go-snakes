package tournament

import (
	"github.com/wmurray8989/go-snakes/match"
	"github.com/wmurray8989/go-snakes/player"
)

// Tournament is a collection of matches
type Tournament struct {
	series32           [32]player.Player
	series16           [16]player.Player
	series8            [8]player.Player
	series4            [4]player.Player
	series2            [2]player.Player
	champion           player.Player
	activeMatch        match.Match
	activePlayer1Index int
	activePlayer2Index int
}
