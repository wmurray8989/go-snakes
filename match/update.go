package match

import (
	"time"

	"github.com/wmurray8989/go-snakes/round"
)

// Update runs the match
func (m *Match) Update() {
	if m.status != InProgress {
		return
	}

	m.activeRound.Update()

	m.timeRemaining = time.Minute*2 - time.Since(m.startTime)

	roundStatus := m.activeRound.GetStatus()
	if roundStatus != round.InProgress && m.lastRoundStatus == round.InProgress {
		m.roundEnd = true
	}

	if roundStatus != round.InProgress && m.roundEnd == false {
		// Increment winning player's points
		switch roundStatus {
		case round.WinnerSnake1:
			m.player1Points++
		case round.WinnerSnake2:
			m.player2Points++
		}

		// If match is not complete, start a new round
		if m.timeRemaining > 0 || m.player1Points == m.player2Points {
			m.activeRound = round.NewRound(m.player1, m.player2)
		} else {
			// Round is complete
			if m.player1Points > m.player2Points {
				println("Snake 1 Wins Match")
				m.status = WinnerSnake1
			} else {
				println("Snake 2 Wins Match")
				m.status = WinnerSnake2
			}
		}
	}

	m.lastRoundStatus = roundStatus
}
