package round

import (
	"math/rand"

	"github.com/wmurray8989/go-snakes/player"
)

// NewRound creates a Round
func NewRound(player1 player.Player, player2 player.Player) Round {
	var round = Round{}

	// random starting player
	startingPlayer := rand.Intn(1)
	if startingPlayer == 0 {
		round.playerTurn = false
	} else {
		round.playerTurn = true
	}

	// setup players
	round.player1 = player1
	round.player2 = player2

	// setup starting positions
	round.player1.StartMove()
	round.player2.StartMove()

	// setup colors
	round.gridColor.R = 100
	round.gridColor.G = 100
	round.gridColor.B = 100
	round.gridColor.A = 100

	return round
}
