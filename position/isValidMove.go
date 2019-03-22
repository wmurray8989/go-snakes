package position

func (p Position) IsValidMove(self []Position, opponent []Position) bool {
	// Is within bounds
	if !p.IsWithinBounds() {
		return false
	}

	// Distance from last position is 1
	lastPosition := self[len(self)-1]
	if p.DistanceTo(lastPosition) != 1 {
		return false
	}

	// Is unoccupied
	if !p.IsUnoccupied(self, opponent) {
		return false
	}
	return true
}
