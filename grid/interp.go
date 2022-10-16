package grid

import (
	"gonum.org/v1/gonum/interp"
)

var spline interp.AkimaSpline
var mils_table []float64
var ranges_table []float64

func init() {
	mils_table := []float64{
		1579,
		1558,
		1538,
		1517,
		1496,
		1475,
		1453,
		1431,
		1409,
		1387,
		1364,
		1341,
		1317,
		1292,
		1267,
		1240,
		1212,
		1183,
		1152,
		1118,
		1081,
		1039,
		988,
		918,
		800,
	}

	ranges_table := make([]float64, 25)
	i := 0
	for n := 50; n < 1251; n += 50 {
		ranges_table[i] = float64(n)
		i++
	}

	spline.Fit(ranges_table, mils_table)

}

func GetMils(dist float64) float64 {
	return spline.Predict(dist)
}
