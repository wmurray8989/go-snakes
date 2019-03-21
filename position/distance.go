package position

func integerAbs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func (p Position) DistanceTo(other Position) int {
	return integerAbs(p.X-other.X) + integerAbs(p.Y-other.Y)
}
