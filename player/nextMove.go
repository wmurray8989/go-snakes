package player

import (
	"errors"
)

const sideLength = 50
const cellSize = 10

// NextMove returns the player's next move
func (p *Player) NextMove(opponent Player) (err error) {
	// Check if no moves are available
	currentPosition := p.moves[len(p.moves)-1]
	validMoves := currentPosition.GetValidMoves(p.moves, opponent.moves)
	if len(validMoves) == 0 {
		return errors.New("No Valid Moves")
	}

	// Handle panics
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Panicked")
		}
	}()

	move := p.strategy(p.moves, opponent.moves)
	if !(move.IsValidMove(p.moves, opponent.moves)) {
		err = errors.New("Invalid Move")
	}

	p.moves = append(p.moves, move)
	return err
}
