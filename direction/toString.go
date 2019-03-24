package direction

// ToString returns a string description of the direction.
func (d Direction) ToString() string {
	switch d {
	case Up:
		return "up"
	case Right:
		return "right"
	case Down:
		return "down"
	case Left:
		return "left"
	}
	return ""
}
