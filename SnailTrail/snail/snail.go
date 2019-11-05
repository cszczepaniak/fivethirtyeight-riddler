package snail

import (
	"math"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/point"
	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/vector"
)

var hexPoints = [...]point.Point2{
	{X: 0, Y: 5 * math.Sqrt(3)},
	{X: 5, Y: 10 * math.Sqrt(3)},
	{X: 15, Y: 10 * math.Sqrt(3)},
	{X: 20, Y: 5 * math.Sqrt(3)},
	{X: 15, Y: 0},
	{X: 5, Y: 0},
}

// Snail represents one of the snails
type Snail struct {
	Pos point.Point2
	Dir vector.Vector2

	// Link to the next snail
	next *Snail
}

// InitSnails creates the six snails
func InitSnails() []Snail {
	snails := make([]Snail, 6)
	for i, p := range hexPoints {
		snails[i].Pos = p
		snails[i].next = &snails[(i+1)%6]
	}
	return snails
}
