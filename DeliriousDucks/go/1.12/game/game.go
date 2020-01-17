package game

import (
	"github.com/cszczepaniak/fivethirtyeight-riddler/DeliriousDucks/board"
	"github.com/cszczepaniak/fivethirtyeight-riddler/DeliriousDucks/duck"
)

type Game struct {
	Board board.Board
	Ducks []duck.Duck
}

type GameConfig struct {
	BoardWidth  int
	BoardHeight int
	NumDucks    int
	StartX      int
	StartY      int
}

func New(c GameConfig) Game {
	b := board.NewBoard(c.BoardWidth, c.BoardHeight)
	ds := make([]duck.Duck, c.NumDucks)
	for i := range ds {
		ds[i] = duck.New(c.StartX, c.StartY, b)
	}
	return Game{
		Board: b,
		Ducks: ds,
	}
}

func (g *Game) Play() int {
	iter := 0
	for iter == 0 || !g.isOver() {
		for _, d := range g.Ducks {
			d.Move()
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
