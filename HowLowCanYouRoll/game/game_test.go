package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_updateScore(t *testing.T) {
	tests := []struct {
		name   string
		score  int64
		roll   int64
		expect int64
	}{
		{
			name:   `easy test`,
			score:  9,
			roll:   1,
			expect: 91,
		},
	}
	for _, tc := range tests {
		score := updateScore(tc.score, tc.roll)
		assert.Equal(t, tc.expect, score, tc.name)
	}
}
