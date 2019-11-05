package utils

import "math"

func Norm2(vec []float64) float64 {
	var accum float64 = 0
	for _, n := range vec {
		accum += math.Pow(n, 2)
	}
	return math.Sqrt(accum)
}
