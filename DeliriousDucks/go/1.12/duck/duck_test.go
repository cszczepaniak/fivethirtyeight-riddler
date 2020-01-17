package duck

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cszczepaniak/fivethirtyeight-riddler/DeliriousDucks/board"
)

func TestFindNeighboringPoints(t *testing.T) {
	tests := []struct {
		desc string
		x    int
		y    int
		b    board.Board
		exp  []board.Point
	}{{
		desc: `test when in middle square`,
		x:    1,
		y:    1,
		b:    board.NewBoard(3, 3),
		exp: []board.Point{
			board.NewPoint(1, 0),
			board.NewPoint(2, 1),
			board.NewPoint(1, 2),
			board.NewPoint(0, 1),
		},
	}, {
		desc: `test when in corner square`,
		x:    0,
		y:    0,
		b:    board.NewBoard(3, 3),
		exp: []board.Point{
			board.NewPoint(1, 0),
			board.NewPoint(0, 1),
		},
	}, {
		desc: `test when in corner square`,
		x:    0,
		y:    2,
		b:    board.NewBoard(3, 3),
		exp: []board.Point{
			board.NewPoint(0, 1),
			board.NewPoint(1, 2),
		},
	}, {
		desc: `test when in edge square`,
		x:    1,
		y:    0,
		b:    board.NewBoard(3, 3),
		exp: []board.Point{
			board.NewPoint(0, 0),
			board.NewPoint(2, 0),
			board.NewPoint(1, 1),
		},
	}, {
		desc: `test when in edge square`,
		x:    1,
		y:    2,
		b:    board.NewBoard(3, 3),
		exp: []board.Point{
			board.NewPoint(1, 1),
			board.NewPoint(2, 2),
			board.NewPoint(0, 2),
		},
	}}

	for _, tc := range tests {
		d := NewDuck(tc.x, tc.y, tc.b)
		pts := d.findNeighboringPoints()
		assert.Len(t, pts, len(tc.exp))
		for _, p := range tc.exp {
			assert.Contains(t, pts, p)
		}
	}
}
func TestMove(t *testing.T) {
	tests := []struct {
		desc    string
		x       int
		y       int
		b       board.Board
		possPts []board.Point
	}{{
		desc: `test when in middle square`,
		x:    1,
		y:    1,
		b:    board.NewBoard(3, 3),
		possPts: []board.Point{
			board.NewPoint(1, 0),
			board.NewPoint(2, 1),
			board.NewPoint(1, 2),
			board.NewPoint(0, 1),
		},
	}, {
		desc: `test when in corner square`,
		x:    0,
		y:    0,
		b:    board.NewBoard(3, 3),
		possPts: []board.Point{
			board.NewPoint(1, 0),
			board.NewPoint(0, 1),
		},
	}, {
		desc: `test when in corner square`,
		x:    0,
		y:    2,
		b:    board.NewBoard(3, 3),
		possPts: []board.Point{
			board.NewPoint(0, 1),
			board.NewPoint(1, 2),
		},
	}, {
		desc: `test when in edge square`,
		x:    1,
		y:    0,
		b:    board.NewBoard(3, 3),
		possPts: []board.Point{
			board.NewPoint(0, 0),
			board.NewPoint(2, 0),
			board.NewPoint(1, 1),
		},
	}, {
		desc: `test when in edge square`,
		x:    1,
		y:    2,
		b:    board.NewBoard(3, 3),
		possPts: []board.Point{
			board.NewPoint(1, 1),
			board.NewPoint(2, 2),
			board.NewPoint(0, 2),
		},
	}}

	for _, tc := range tests {
		d := NewDuck(tc.x, tc.y, tc.b)
		d.Move()
		assert.Contains(t, tc.possPts, d.Pos)
	}
}
