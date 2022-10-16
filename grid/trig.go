package grid

import "math"

func GetAngleDegrees(to Position) float64 {
	angle_radians := math.Atan2(to.x, -to.y)
	angle := angle_radians * (180 / math.Pi)

	if angle < 0 {
		angle = 360 + angle
	}

	return angle
}

func GetDistance(to Position) float64 {
	x2 := math.Pow(to.x, 2)
	y2 := math.Pow(to.y, 2)
	return math.Sqrt(x2 + y2)
}
