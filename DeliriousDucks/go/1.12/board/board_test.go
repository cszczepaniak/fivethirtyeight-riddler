package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	tests := []struct {
		desc string
		w    int
		h    int
		exp  Board
	}{{
		desc: `test 3x3`,
		w:    3,
		h:    3,
		exp: []Point{
			NewPoint(0, 0),
			NewPoint(1, 0),
			NewPoint(2, 0),
			NewPoint(0, 1),
			NewPoint(1, 1),
			NewPoint(2, 1),
			NewPoint(0, 2),
			NewPoint(1, 2),
			NewPoint(2, 2),
		},
	}, {
		desc: `test 2x4`,
		w:    2,
		h:    4,
		exp: []Point{
			NewPoint(0, 0),
			NewPoint(1, 0),
			NewPoint(0, 1),
			NewPoint(1, 1),
			NewPoint(0, 2),
			NewPoint(1, 2),
			NewPoint(0, 3),
			NewPoint(1, 3),
		},
	}}

	for _, tc := range tests {
		b := NewBoard(tc.w, tc.h)
		assert.Equal(t, tc.exp, b)
	}
}
