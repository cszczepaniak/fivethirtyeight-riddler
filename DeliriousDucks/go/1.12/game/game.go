package game

import (
	"math/rand"
	"time"

	"github.com/cszczepaniak/fivethirtyeight-riddler/DeliriousDucks/board"
	"github.com/cszczepaniak/fivethirtyeight-riddler/DeliriousDucks/duck"
)

type Game struct {
	Board board.Board
	Ducks []duck.Duck

	r *rand.Rand
}

type Config struct {
	BoardWidth  int
	BoardHeight int
	NumDucks    int
	StartX      int
	StartY      int
}

func New(c Config) Game {
	b := board.NewBoard(c.BoardWidth, c.BoardHeight)
	ds := make([]duck.Duck, c.NumDucks)
	for i := range ds {
		ds[i] = duck.New(c.StartX, c.StartY, b)
	}
	return Game{
		Board: b,
		Ducks: ds,
		r:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *Game) Play() int {
	iter := 0
	for iter == 0 || !g.isOver() {
		for i, d := range g.Ducks {
			d.Move(g.r)
			g.Ducks[i] = d
		}
		iter++
	}
	return iter
}

func (g *Game) isOver() bool {
	posMap := make(map[board.Point]struct{}, len(g.Ducks))
	for _, d := range g.Ducks {
		if _, ok := posMap[d.Pos]; !ok {
			posMap[d.Pos] = struct{}{}
		}
	}
	return len(posMap) == 1
}
