package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_updateScore(t *testing.T) {
	tests := []struct {
		name      string
		rolls     []int
		expectStr string
	}{
		{
			name:      `simple test`,
			rolls:     []int{9, 1},
			expectStr: `0.91`,
		},
		{
			name:      `test repeated digits`,
			rolls:     []int{9, 9, 7},
			expectStr: `0.997`,
		},
		{
			name:      `test higher digits`,
			rolls:     []int{7, 9, 6},
			expectStr: `0.76`,
		},
		{
			name:      `test a lot of rolls`,
			rolls:     []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			expectStr: `0.9999999999999999999`,
		},
		{
			name:      `test zero`,
			rolls:     []int{0},
			expectStr: `0.0`,
		},
	}
	for _, tc := range tests {
		score := `0.`
		for _, r := range tc.rolls {
			updateScore(&score, r)
		}
		assert.Equal(t, tc.expectStr, score, tc.name)
	}
}
