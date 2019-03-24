package match

import (
	"time"

	"github.com/wmurray8989/go-snakes/round"
)

// Match is a timed series of rounds
type Match struct {
	startTime     time.Time
	player1       round.Strategy
	player2       round.Strategy
	player1Points int
	player2Points int
	activeRound   round.Round
}
