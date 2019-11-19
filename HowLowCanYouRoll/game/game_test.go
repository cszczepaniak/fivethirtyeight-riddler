package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_updateScore(t *testing.T) {
	tests := []struct {
		name   string
		score  int64
		rolls  []int64
		expect int64
	}{
		{
			name:   `easy test`,
			score:  9,
			rolls:  []int64{9, 1},
			expect: 91,
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
