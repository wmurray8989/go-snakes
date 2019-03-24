package position

import (
	"github.com/wmurray8989/go-snakes/direction"
)

func (p Position) DirectionTo(other Position) direction.Direction {
	xDistance := other.X - p.X
	yDistance := other.Y - p.Y

	if xDistance == 0 && yDistance == 0 {
		return direction.Null
	}

	if integerAbs(xDistance) > integerAbs(yDistance) {
		// horizontal
		if xDistance > 0 {
			return direction.Right
		} else {
			return direction.Left
		}
	} else {
		// vertical
		if yDistance > 0 {
			return direction.Up
		} else {
			return direction.Down
		}
	}
}
