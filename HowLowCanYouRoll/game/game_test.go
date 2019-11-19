package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_updateScore(t *testing.T) {
	tests := []struct {
		name   string
		rolls  []int64
		expect int64
	}{
		{
			name:   `simple test`,
			rolls:  []int64{9, 1},
			expect: 91,
		},
		{
			name:   `test repeated digits`,
			rolls:  []int64{9, 9, 7},
			expect: 997,
		},
		{
			name:   `test higher digits`,
			rolls:  []int64{7, 9, 6},
			expect: 76,
		},
	}
	for _, tc := range tests {
		var score int64 = 0
		for _, r := range tc.rolls {
			score = updateScore(score, r)
		}
		assert.Equal(t, tc.expect, score, tc.name)
	}
}
