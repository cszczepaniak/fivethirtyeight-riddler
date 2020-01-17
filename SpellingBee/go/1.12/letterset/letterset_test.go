package letterset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		word string
		exp  LetterSet
	}{{
		word: `amalgam`,
		exp:  `aglm`,
	}, {
		word: `megaplex`,
		exp:  `aeglmpx`,
	}}
	for _, tc := range tests {
		ls := New(tc.word)
		assert.Equal(t, tc.exp, ls)
	}
}
func TestLength(t *testing.T) {
	tests := []struct {
		word string
		exp  int
	}{{
		word: `amalgam`,
		exp:  4,
	}, {
		word: `megaplex`,
		exp:  7,
	}}
	for _, tc := range tests {
		l := Length(tc.word)
		assert.Equal(t, tc.exp, l)
	}
}
