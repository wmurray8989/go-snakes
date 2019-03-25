package tournament

import (
	"github.com/wmurray8989/go-snakes/match"
	"github.com/wmurray8989/go-snakes/player"
)

// NewTournament creates a Tournament
func NewTournament(players [32]player.Player) Tournament {
	return Tournament{
		series32:           players,
		activePlayer1Index: 0,
		activePlayer2Index: 1,
		activeMatch:        match.NewMatch(players[0], players[1]),
	}
}
