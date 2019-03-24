package match

import (
	"time"

	"github.com/wmurray8989/go-snakes/round"
)

// NewMatch creates a Match
func NewMatch(player1 round.Strategy, player2 round.Strategy) Match {
	var match = Match{}

	// record start time
	match.startTime = time.Now().UTC()

	// create first round
	match.activeRound = round.NewRound(player1, player2)

	return match
}
