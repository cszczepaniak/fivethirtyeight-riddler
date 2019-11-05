package vector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector2_normalize(t *testing.T) {
	tests := []struct {
		name   string
		vec    Vector2
		expect Vector2
	}{
		{
			name:   "test easy distance",
			vec:    Vector2{X: 3, Y: 4},
			expect: Vector2{X: 0.6, Y: 0.8},
		},
		{
			name:   "test negative point",
			vec:    Vector2{X: 3, Y: -4},
			expect: Vector2{X: 0.6, Y: -0.8},
		},
		{
			name:   "test irrational",
			vec:    Vector2{X: 1, Y: 1},
			expect: Vector2{X: 1 / math.Sqrt(2), Y: 1 / math.Sqrt(2)},
		},
	}
	for _, tc := range tests {
		tc.vec.normalize()
		assert.InEpsilon(t, tc.expect.X, tc.vec.X, 1e-8, tc.name)
	}
}
