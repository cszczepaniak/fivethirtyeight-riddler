package game

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_updateScore(t *testing.T) {
	tests := []struct {
		name      string
		rolls     []int64
		expectStr string
	}{
		{
			name:      `simple test`,
			rolls:     []int64{9, 1},
			expectStr: `91`,
		},
		{
			name:      `test repeated digits`,
			rolls:     []int64{9, 9, 7},
			expectStr: `997`,
		},
		{
			name:      `test higher digits`,
			rolls:     []int64{7, 9, 6},
			expectStr: `76`,
		},
		{
			name:      `test really big int`,
			rolls:     []int64{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			expectStr: `9999999999999999999`,
		},
	}
	for _, tc := range tests {
		rolls := make([]*big.Int, len(tc.rolls))
		for i, r := range tc.rolls {
			rolls[i] = big.NewInt(r)
		}
		expect := big.NewInt(0)
		if _, success := expect.SetString(tc.expectStr, 10); !success {
			panic(`conversion failed`)
		}

		score := big.NewInt(int64(0))
		for _, r := range tc.rolls {
			score = updateScore(score, r)
		}
		assert.Equal(t, expect, score, tc.name)
	}
}
