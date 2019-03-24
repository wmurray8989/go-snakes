package position

// Position contains an X and Y coordinate
type Position struct {
	X int
	Y int
}

// Null is a null position
var Null = Position{
	X: -1,
	Y: -1,
}
