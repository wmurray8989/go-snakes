package match

import (
	"time"

	"github.com/wmurray8989/go-snakes/player"
	"github.com/wmurray8989/go-snakes/round"
)

// Match is a timed series of rounds
type Match struct {
	startTime     time.Time
	player1       player.Player
	player2       player.Player
	player1Points int
	player2Points int
	activeRound   round.Round
}
