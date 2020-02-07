package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s   string
		exp bool
	}{{
		s:   `abc`,
		exp: false,
	}, {
		s:   `abcd`,
		exp: false,
	}, {
		s:   `racecar`,
		exp: true,
	}, {
		s:   `abba`,
		exp: true,
	}, {
		s:   `amanaplanacanalpanama`,
		exp: true,
	}, {
		s:   `02022020`,
		exp: true,
	}}
	for _, tc := range tests {
		got := isPalindrome(tc.s)
		assert.Equal(t, tc.exp, got)
	}
}
