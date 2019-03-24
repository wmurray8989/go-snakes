package round

import (
	"math/rand"
	"time"

	"github.com/wmurray8989/go-snakes/position"
)

// NewRound creates a Round
func NewRound(player1 Strategy, player2 Strategy) Round {
	var round = Round{}

	// random starting player
	startingPlayer := rand.Intn(1)
	if startingPlayer == 0 {
		round.playerTurn = false
	} else {
		round.playerTurn = true
	}

	// setup strategies
	round.player1 = player1
	round.player2 = player2

	// setup starting positions
	rand.Seed(time.Now().UTC().UnixNano())
	round.player1History = append(round.player1History, position.Position{X: rand.Intn(49), Y: rand.Intn(49)})
	round.player2History = append(round.player2History, position.Position{X: rand.Intn(49), Y: rand.Intn(49)})

	// setup colors
	round.p1Color.R = 255
	round.p1Color.A = 255

	round.p2Color.G = 255
	round.p2Color.A = 255

	round.gridColor.R = 100
	round.gridColor.G = 100
	round.gridColor.B = 100
	round.gridColor.A = 100

	return round
}
