package position

import (
	"github.com/wmurray8989/go-snakes/direction"
)

func (p Position) DirectionAdd(d direction.Direction) Position {
	switch d {
	case direction.Up:
		p.Y = p.Y + 1
	case direction.Right:
		p.X = p.X + 1
	case direction.Down:
		p.Y = p.Y - 1
	case direction.Left:
		p.X = p.X - 1
	}
	return p
}
