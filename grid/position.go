package grid

type Position struct {
	x float64
	y float64
}

func NewPosition(x float64, y float64) Position {
	return Position{x: x, y: y}
}

func (p Position) Scale(factor float64) Position {
	return Position{
		x: p.x * factor,
		y: p.y * factor,
	}
}

func (p Position) Add(other Position) Position {
	return Position{
		x: p.x + other.x,
		y: p.y + other.y,
	}
}

func (p Position) Sub(other Position) Position {
	return Position{
		x: p.x - other.x,
		y: p.y - other.y,
	}
}
