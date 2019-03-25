package tournament

// Update runs the tournament
func (t *Tournament) Update() {
	// Update active match
	t.activeMatch.Update()

	// matchStatus := t.activeMatch.GetStatus()
	// if matchStatus != match.InProgress {
	// 	// Get winning snake
	// 	var winningSnake player.Player
	// 	switch matchStatus {
	// 	case match.WinnerSnake1:
	// 		t.activePlayer1Index
	// 	case match.WinnerSnake2:
	// 		t.player2Points++
	// 	}
	// }
}
