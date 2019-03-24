package direction

// TurnLeft returns a new direction rotated 90 degrees anti-clockwise.
func (d Direction) TurnLeft() Direction {
	if (d == 0) {
		return Direction(3)
	}
	return Direction(d - 1)
}
