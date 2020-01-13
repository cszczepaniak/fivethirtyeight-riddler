package snail

import (
	"math"
	"testing"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/point"
	"github.com/stretchr/testify/assert"
)

func TestSnail_Step(t *testing.T) {
	tests := []struct {
		name         string
		start        point.Point2
		nextSnailPos point.Point2
		nSteps       int
		stepSize     float64
		expect       point.Point2
	}{
		{
			name:         "test one step, one dimensional",
			start:        point.Point2{X: 0, Y: 0},
			nextSnailPos: point.Point2{X: 1, Y: 0},
			nSteps:       1,
			stepSize:     0.5,
			expect:       point.Point2{X: 0.5, Y: 0},
		},
		{
			name:         "test multiple steps, one dimensional",
			start:        point.Point2{X: 0, Y: 0},
			nextSnailPos: point.Point2{X: 1, Y: 0},
			nSteps:       5,
			stepSize:     0.1,
			expect:       point.Point2{X: 0.5, Y: 0},
		},
		{
			name:         "test one step, two dimensional",
			start:        point.Point2{X: 0, Y: 0},
			nextSnailPos: point.Point2{X: 1, Y: 1},
			nSteps:       1,
			stepSize:     0.5,
			expect:       point.Point2{X: 0.5 / math.Sqrt(2), Y: 0.5 / math.Sqrt(2)},
		},
		{
			name:         "test overshoot",
			start:        point.Point2{X: 0, Y: 0},
			nextSnailPos: point.Point2{X: 1, Y: 0},
			nSteps:       2,
			stepSize:     2,
			expect:       point.Point2{X: 0, Y: 0},
		},
	}
	for _, tc := range tests {
		thisSnail := Snail{Pos: tc.start, next: &Snail{Pos: tc.nextSnailPos}}
		for i := 0; i < tc.nSteps; i++ {
			thisSnail.Step(tc.stepSize)
		}
		// assert.InEpsilon can't test against an expected value of zero. We have to explicitly
		// check for zero.
		if tc.expect.X == 0 {
			assert.Zero(t, thisSnail.Pos.X, tc.name)
		} else {
			assert.InEpsilon(t, tc.expect.X, thisSnail.Pos.X, 1e-8, tc.name)
		}
		if tc.expect.Y == 0 {
			assert.Zero(t, thisSnail.Pos.Y, tc.name)
		} else {
			assert.InEpsilon(t, tc.expect.Y, thisSnail.Pos.Y, 1e-8, tc.name)
		}
	}
}
