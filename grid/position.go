package grid

type Position struct {
	x float64
	y float64
}

func NewPosition(x float64, y float64) Position {
	return Position{x: x, y: y}
}

func (p *Position) Scale(factor float64) {
	p.x = p.x * factor
	p.y = p.y * factor
}

func (p *Position) Add(other Position) {
	p.x = p.x + other.x
	p.y = p.y + other.y
}

func (p *Position) Sub(other Position) {
	p.x = p.x - other.x
	p.y = p.y - other.y
}
