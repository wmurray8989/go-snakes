package tournament

import (
	"github.com/wmurray8989/go-snakes/match"
	"github.com/wmurray8989/go-snakes/player"
)

// Tournament is a collection of matches
type Tournament struct {
	series32    []player.Player
	series16    []player.Player
	series8     []player.Player
	series4     []player.Player
	series2     []player.Player
	champion    player.Player
	activeMatch match.Match
}
