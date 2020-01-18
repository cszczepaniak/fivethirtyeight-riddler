package utils

import (
	"testing"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/board"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoardSubsets(t *testing.T) {
	type boardConfig struct {
		middle rune
		others []rune
	}
	tests := []struct {
		b    boardConfig
		nExp int
	}{{
		b: boardConfig{
			middle: 'a',
			others: []rune{'b', 'c', 'd', 'e', 'f', 'g'},
		},
		nExp: 64,
	}}
	for _, tc := range tests {
		b, err := board.New(tc.b.middle, tc.b.others)
		require.NoError(t, err)
		combs := BoardSubsets(b)
		assert.Equal(t, tc.nExp, len(combs))
	}
}
