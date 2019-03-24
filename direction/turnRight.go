package direction

// TurnRight returns a new direction rotated 90 degrees clockwise.
func (d Direction) TurnRight() Direction {
	if (d == 3) {
		return Direction(0)
	}
	return Direction( d + 1)
}
