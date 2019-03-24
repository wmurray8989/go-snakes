package position

func (p Position) IsWithinBounds() bool {
	return !(p.X < 0 || p.X > 49 || p.Y < 0 || p.Y > 49)
}
