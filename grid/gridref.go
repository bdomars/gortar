package grid

import (
	"fmt"
	"math"
	"strings"
)

type GridRef struct {
	Letter  byte
	Major   uint8
	Keypads []uint8
	Grid    *Grid
}

func (g GridRef) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%c%d", g.Letter, g.Major))
	for _, kp := range g.Keypads {
		b.WriteString(fmt.Sprintf("-%d", kp))
	}
	return b.String()
}

func (g GridRef) LetterValue() uint8 {
	return uint8(g.Letter-'A') + 1
}

func (g GridRef) Position() Position {
	base_size := g.Grid.BaseSize
	base_x := float64(g.LetterValue()) - 0.5
	base_y := float64(g.Major) - 0.5

	p := Position{
		x: base_x * base_size,
		y: base_y * base_size,
	}

	for n, kp := range g.Keypads {
		subcoord := kp_to_pos(kp)
		subcoord_scale := base_size / math.Pow(3, float64(n+1))
		subcoord = subcoord.Scale(subcoord_scale)
		p = p.Add(subcoord)
	}

	return p
}

func kp_to_pos(kp uint8) Position {
	kpint := int(kp)
	return Position{
		x: float64((kpint-1)%3 - 1),
		y: float64(1 - (kpint-1)/3),
	}
}

func (this GridRef) EqualTo(other GridRef) bool {

	if this.Letter != other.Letter {
		return false
	}

	if this.Major != other.Major {
		return false
	}

	if len(this.Keypads) != len(other.Keypads) {
		return false
	}

	for i := range this.Keypads {
		if this.Keypads[i] != other.Keypads[i] {
			return false
		}
	}

	return true
}
