package duck

import (
	"math/rand"

	"github.com/cszczepaniak/fivethirtyeight-riddler/DeliriousDucks/board"
)

type Duck struct {
	Pos   board.Point
	Board board.Board
}

func New(x, y int, b board.Board) Duck {
	return Duck{
		Pos:   board.NewPoint(x, y),
		Board: b,
	}
}

func (d *Duck) Move() {
	pts := d.findNeighboringPoints()
	idx := rand.Intn(len(pts))
	p := pts[idx]
	d.Pos.X = p.X
	d.Pos.Y = p.Y
}

func (d *Duck) findNeighboringPoints() []board.Point {
	// can have at most 4 neighboring points
	pts := make([]board.Point, 0, 4)
	for _, p := range d.Board {
		if p.Y == d.Pos.Y {
			if p.X == d.Pos.X+1 || p.X == d.Pos.X-1 {
				pts = append(pts, p)
			}
		}
		if p.X == d.Pos.X {
			if p.Y == d.Pos.Y+1 || p.Y == d.Pos.Y-1 {
				pts = append(pts, p)
			}
		}
	}
	return pts
}
