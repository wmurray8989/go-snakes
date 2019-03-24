package round

import "fmt"

// Update runs the round
func (p *Round) Update() {
	if p.status != InProgress {
		return
	}

	if p.playerTurn {
		err := p.player1.NextMove(p.player2)
		if err != nil {
			fmt.Println(err)
			p.status = WinnerSnake2
		}
	} else {
		err := p.player2.NextMove(p.player1)
		if err != nil {
			fmt.Println(err)
			p.status = WinnerSnake1
		}
	}

	p.playerTurn = !p.playerTurn
}
