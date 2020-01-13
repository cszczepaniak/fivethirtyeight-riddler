package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_updateScore(t *testing.T) {
	tests := []struct {
		name   string
		rolls  []int
		expect uint64
	}{
		{
			name:   `simple test`,
			rolls:  []int{9, 1},
			expect: 91,
		},
		{
			name:   `test repeated digits`,
			rolls:  []int{9, 9, 7},
			expect: 997,
		},
		{
			name:   `test higher digits`,
			rolls:  []int{7, 9, 6},
			expect: 76,
		},
		{
			name:   `test zero`,
			rolls:  []int{0},
			expect: 0,
		},
	}
	for _, tc := range tests {
		var score uint64 = 0
		for _, r := range tc.rolls {
			updateScore(&score, r)
		}
		assert.Equal(t, tc.expect, score, tc.name)
	}
}
