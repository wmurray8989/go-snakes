package tournament

import (
	"github.com/wmurray8989/go-snakes/match"
)

// Update runs the tournament
func (t *Tournament) Update() {
	// Update active match
	t.activeMatch.Update()

	matchStatus := t.activeMatch.GetStatus()
	if matchStatus != match.InProgress {
		// Move winner to next round
		switch t.activeSeriesIndex {
		case 0:
			switch matchStatus {
			case match.WinnerSnake1:
				t.series16[t.activePlayer1Index/2] = t.series32[t.activePlayer1Index]
			case match.WinnerSnake2:
				t.series16[t.activePlayer1Index/2] = t.series32[t.activePlayer2Index]
			}
			if t.activePlayer1Index == 30 {
				// next series
				t.activeSeriesIndex++
				t.activePlayer1Index = 0
				t.activePlayer2Index = 1
				t.activeMatch = match.NewMatch(t.series16[t.activePlayer1Index], t.series16[t.activePlayer2Index])
			} else {
				t.activePlayer1Index = t.activePlayer1Index + 2
				t.activePlayer2Index = t.activePlayer2Index + 2
				t.activeMatch = match.NewMatch(t.series32[t.activePlayer1Index], t.series32[t.activePlayer2Index])
			}
		case 1:
			switch matchStatus {
			case match.WinnerSnake1:
				t.series8[t.activePlayer1Index/2] = t.series16[t.activePlayer1Index]
			case match.WinnerSnake2:
				t.series8[t.activePlayer1Index/2] = t.series16[t.activePlayer2Index]
			}
			if t.activePlayer1Index == 14 {
				// next series
				t.activeSeriesIndex++
				t.activePlayer1Index = 0
				t.activePlayer2Index = 1
				t.activeMatch = match.NewMatch(t.series8[t.activePlayer1Index], t.series8[t.activePlayer2Index])
			} else {
				t.activePlayer1Index = t.activePlayer1Index + 2
				t.activePlayer2Index = t.activePlayer2Index + 2
				t.activeMatch = match.NewMatch(t.series16[t.activePlayer1Index], t.series16[t.activePlayer2Index])
			}
		case 2:
			switch matchStatus {
			case match.WinnerSnake1:
				t.series4[t.activePlayer1Index/2] = t.series8[t.activePlayer1Index]
			case match.WinnerSnake2:
				t.series4[t.activePlayer1Index/2] = t.series8[t.activePlayer2Index]
			}
			if t.activePlayer1Index == 6 {
				// next series
				t.activeSeriesIndex++
				t.activePlayer1Index = 0
				t.activePlayer2Index = 1
				t.activeMatch = match.NewMatch(t.series4[t.activePlayer1Index], t.series4[t.activePlayer2Index])
			} else {
				t.activePlayer1Index = t.activePlayer1Index + 2
				t.activePlayer2Index = t.activePlayer2Index + 2
				t.activeMatch = match.NewMatch(t.series8[t.activePlayer1Index], t.series8[t.activePlayer2Index])
			}
		case 3:
			switch matchStatus {
			case match.WinnerSnake1:
				t.series2[t.activePlayer1Index/2] = t.series4[t.activePlayer1Index]
			case match.WinnerSnake2:
				t.series2[t.activePlayer1Index/2] = t.series4[t.activePlayer2Index]
			}
			if t.activePlayer1Index == 2 {
				// next series
				t.activeSeriesIndex++
				t.activePlayer1Index = 0
				t.activePlayer2Index = 1
				t.activeMatch = match.NewMatch(t.series2[t.activePlayer1Index], t.series2[t.activePlayer2Index])
			} else {
				t.activePlayer1Index = t.activePlayer1Index + 2
				t.activePlayer2Index = t.activePlayer2Index + 2
				t.activeMatch = match.NewMatch(t.series4[t.activePlayer1Index], t.series4[t.activePlayer2Index])
			}
		case 4:
			switch matchStatus {
			case match.WinnerSnake1:
				t.champion = t.series2[t.activePlayer1Index]
			case match.WinnerSnake2:
				t.champion = t.series2[t.activePlayer2Index]
			}
		}
	}
}
