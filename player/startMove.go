package player

import (
	"math/rand"
	"time"

	"github.com/wmurray8989/go-snakes/position"
)

// StartMove resets the player's moves and generates a random starting position
func (p *Player) StartMove() {
	rand.Seed(time.Now().UTC().UnixNano())
	p.Moves = []position.Position{
		position.Position{
			X: rand.Intn(49),
			Y: rand.Intn(49),
		},
	}
}
