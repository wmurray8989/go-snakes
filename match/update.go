package match

// Update runs the match
func (p *Match) Update() {
	if p.status != matchRunning {
		return
	}

	p1Move := p.player1(p.player1History, p.player2History)
	if !(p1Move.IsValidMove(p.player1History, p.player2History)) {
		p.status = snake2Wins
	}
	p.player1History = append(p.player1History, p1Move)

	p2Move := p.player2(p.player2History, p.player1History)
	if !(p2Move.IsValidMove(p.player2History, p.player1History)) {
		p.status = snake1Wins
	}
	p.player2History = append(p.player2History, p2Move)
}
