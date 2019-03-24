package round

// Update runs the round
func (p *Round) Update() {
	if p.status != InProgress {
		return
	}

	// Handle panics
	defer func() {
		if r := recover(); r != nil {
			if p.playerTurn {
				p.status = WinnerSnake2
			} else {
				p.status = WinnerSnake1
			}
		}
	}()

	if p.playerTurn {
		p1Move := p.player1(p.player1History, p.player2History)
		if !(p1Move.IsValidMove(p.player1History, p.player2History)) {
			p.status = WinnerSnake2
		}
		p.player1History = append(p.player1History, p1Move)
	} else {
		p2Move := p.player2(p.player2History, p.player1History)
		if !(p2Move.IsValidMove(p.player2History, p.player1History)) {
			p.status = WinnerSnake1
		}
		p.player2History = append(p.player2History, p2Move)
	}

	p.playerTurn = !p.playerTurn
}
