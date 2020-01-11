package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberAsWord(t *testing.T) {
	tests := []struct {
		n   int
		exp string
	}{{
		n:   1000,
		exp: `one thousand`,
	}, {
		n:   1234567,
		exp: `one million two hundred thirty four thousand five hundred sixty seven`,
	}, {
		n:   12345,
		exp: `twelve thousand three hundred fourty five`,
	}, {
		n:   987654,
		exp: `nine hundred eighty seven thousand six hundred fifty four`,
	}, {
		n:   123,
		exp: `one hundred twenty three`,
	}, {
		n:   999,
		exp: `nine hundred ninety nine`,
	}, {
		n:   387,
		exp: `three hundred eighty seven`,
	}, {
		n:   13,
		exp: `thirteen`,
	}, {
		n:   99,
		exp: `ninety nine`,
	}, {
		n:   1,
		exp: `one`,
	}}
	for _, tc := range tests {
		w, err := numberAsWord(tc.n)
		assert.NoError(t, err)
		assert.Equal(t, tc.exp, w)
	}
}
func TestLowestThreeDigits(t *testing.T) {
	tests := []struct {
		n   int
		exp string
	}{{
		n:   123,
		exp: `one hundred twenty three`,
	}, {
		n:   999,
		exp: `nine hundred ninety nine`,
	}, {
		n:   387,
		exp: `three hundred eighty seven`,
	}, {
		n:   13,
		exp: `thirteen`,
	}, {
		n:   99,
		exp: `ninety nine`,
	}, {
		n:   1,
		exp: `one`,
	}}
	for _, tc := range tests {
		w := lowestThreeDigits(tc.n)
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
