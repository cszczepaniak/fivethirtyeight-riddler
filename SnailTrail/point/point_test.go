package point

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint_DistanceTo(t *testing.T) {
	tests := []struct {
		name   string
		pt     Point
		expect float64
	}{
		{
			name:   "test easy distance",
			pt:     Point{X: 3, Y: 4},
			expect: 5,
		},
		{
			name:   "test negative point",
			pt:     Point{X: 3, Y: -4},
			expect: 5,
		},
		{
			name:   "test irrational",
			pt:     Point{X: 1, Y: 1},
			expect: math.Sqrt(2),
		},
	}
	refPt := Point{X: 0, Y: 0}
	for _, tc := range tests {
		dist := refPt.DistanceTo(tc.pt)
		assert.InEpsilon(t, tc.expect, dist, 1e-8, tc.name)
	}
}
