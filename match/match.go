package match

import (
	"time"

	"github.com/wmurray8989/go-snakes/player"
	"github.com/wmurray8989/go-snakes/round"
)

// Match is a timed series of rounds
type Match struct {
	status        Status
	startTime     time.Time
	timeRemaining time.Duration
	player1       player.Player
	player2       player.Player
	player1Points int
	player2Points int
	activeRound   round.Round
}

// Status is the current state of the match
type Status int

// InProgress is the state when the match has no winners
const InProgress Status = 0

// WinnerSnake1 is the state when snake 1 has won the match
const WinnerSnake1 Status = 1

// WinnerSnake2 is the state when snake 2 has won the match
const WinnerSnake2 Status = 2
