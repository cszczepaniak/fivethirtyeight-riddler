package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordValue(t *testing.T) {
	tests := []struct {
		word string
		exp  int
	}{{
		word: `riddler`,
		exp:  70,
	}, {
		word: `one`,
		exp:  34,
	}, {
		word: `two`,
		exp:  58,
	}}
	for _, tc := range tests {
		v := wordValue(tc.word)
		assert.Equal(t, tc.exp, v)
	}
}
