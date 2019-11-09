package point

import (
	"math"
	"testing"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/vector"
	"github.com/stretchr/testify/assert"
)

func TestPoint_DistanceTo(t *testing.T) {
	tests := []struct {
		name   string
		pt     Point2
		expect float64
	}{
		{
			name:   "test easy distance",
			pt:     Point2{X: 3, Y: 4},
			expect: 5,
		},
		{
			name:   "test negative point",
			pt:     Point2{X: 3, Y: -4},
			expect: 5,
		},
		{
			name:   "test irrational",
			pt:     Point2{X: 1, Y: 1},
			expect: math.Sqrt(2),
		},
	}
	refPt := Point2{X: 0, Y: 0}
	for _, tc := range tests {
		dist := refPt.DistanceTo(tc.pt)
		assert.InEpsilon(t, tc.expect, dist, 1e-8, tc.name)
	}
}

func TestPoint2_VectorFrom(t *testing.T) {
	tests := []struct {
		name   string
		p1     Point2
		p2     Point2
		expect vector.Vector2
	}{
		{
			name:   "test easy case",
			p1:     Point2{X: 3, Y: 4},
			p2:     Point2{X: 0, Y: 0},
			expect: vector.Vector2{X: 3, Y: 4},
		},
		{
			name:   "test case not from origin",
			p1:     Point2{X: 4, Y: 5},
			p2:     Point2{X: 1, Y: 1},
			expect: vector.Vector2{X: 3, Y: 4},
		},
		{
			name:   "test negative components",
			p1:     Point2{X: 1, Y: 1},
			p2:     Point2{X: 4, Y: 5},
			expect: vector.Vector2{X: -3, Y: -4},
		},
	}
	for _, tc := range tests {
		vec := tc.p1.VectorFrom(tc.p2)
		assert.InEpsilon(t, tc.expect.X, vec.X, 1e-8, tc.name)
		assert.InEpsilon(t, tc.expect.Y, vec.Y, 1e-8, tc.name)
	}
}
