package match

import (
	"github.com/wmurray8989/go-snakes/round"
)

// Update runs the match
func (m *Match) Update() {
	m.activeRound.Update()

	switch roundStatus := m.activeRound.GetStatus(); roundStatus {
	case round.WinnerSnake1:
		m.player1Points++
		m.activeRound = round.NewRound(m.player1, m.player2)
	case round.WinnerSnake2:
		m.player2Points++
		m.activeRound = round.NewRound(m.player1, m.player2)
	}

}
