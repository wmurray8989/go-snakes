package position

var directions = [4]Position{
	{X: 0, Y: 1},
	{X: 0, Y: -1},
	{X: 1, Y: 0},
	{X: -1, Y: 0},
}

func (p Position) GetValidMoves(self []Position, opponent []Position) []Position {
	availableMoves := []Position{}

	for _, direction := range directions {
		move := Position{
			X: p.X + direction.X,
			Y: p.Y + direction.Y,
		}
		if move.IsUnoccupied(self, opponent) {
			availableMoves = append(availableMoves, move)
		}
	}

	return availableMoves
}
