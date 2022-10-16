package grid

type Grid struct {
	BaseSize float64
}

func NewGrid(size float64) *Grid {
	return &Grid{BaseSize: size}
}

func (g *Grid) GetPosition(ref GridRef) Position {
	return Position{}
}
