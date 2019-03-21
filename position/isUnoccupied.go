package position

func includes(elements []Position, p Position) bool {
	for _, element := range elements {
		if element.X == p.X && element.Y == p.Y {
			return true
		}
	}
	return false
}

func (p Position) IsUnoccupied(self []Position, opponent []Position) bool {
	if includes(self, p) || includes(opponent, p) {
		return false
	}
	return true
}
