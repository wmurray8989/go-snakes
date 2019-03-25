package tournament

import (
	"github.com/wmurray8989/go-snakes/match"
	"github.com/wmurray8989/go-snakes/player"
)

// NewTournament creates a Tournament
func NewTournament(players ...player.Player) Tournament {
	return Tournament{
		series32:    players,
		activeMatch: match.NewMatch(players[0], players[1]),
	}
}
