package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstTwoDigits(t *testing.T) {
	tests := []struct {
		n   int
		exp string
	}{{
		n:   123,
		exp: `twentythree`,
	}, {
		n:   13,
		exp: `thirteen`,
	}, {
		n:   99,
		exp: `ninetynine`,
	}, {
		n:   1,
		exp: `one`,
	}}
	for _, tc := range tests {
		w := firstTwoDigits(tc.n)
		assert.Equal(t, tc.exp, w)
	}
}

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
