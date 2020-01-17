package game

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cszczepaniak/fivethirtyeight-riddler/DeliriousDucks/board"
	"github.com/cszczepaniak/fivethirtyeight-riddler/DeliriousDucks/duck"
)

func makeGame(w, h int, pos []board.Point) Game {
	b := board.NewBoard(w, h)
	ds := make([]duck.Duck, len(pos))
	for i, p := range pos {
		ds[i] = duck.New(p.X, p.Y, b)
	}
	return Game{
		Board: b,
		Ducks: ds,
	}
}

func TestIsOver(t *testing.T) {
	tests := []struct {
		desc string
		g    Game
		exp  bool
	}{{
		desc: `test when in same square`,
		g: makeGame(3, 3, []board.Point{
			board.NewPoint(1, 1),
			board.NewPoint(1, 1),
		}),
		exp: true,
	}, {
		desc: `test when in different square`,
		g: makeGame(3, 3, []board.Point{
			board.NewPoint(1, 2),
			board.NewPoint(1, 1),
		}),
		exp: false,
	}, {
		desc: `test three ducks game over`,
		g: makeGame(3, 3, []board.Point{
			board.NewPoint(1, 1),
			board.NewPoint(1, 1),
			board.NewPoint(1, 1),
		}),
		exp: true,
	}, {
		desc: `test three ducks not game over`,
		g: makeGame(3, 3, []board.Point{
			board.NewPoint(1, 1),
			board.NewPoint(1, 1),
			board.NewPoint(1, 0),
		}),
		exp: false,
	}}

	for _, tc := range tests {
		over := tc.g.isOver()
		assert.Equal(t, tc.exp, over)
	}
}
