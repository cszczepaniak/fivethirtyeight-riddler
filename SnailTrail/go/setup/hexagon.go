package setup

import (
	"math"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/point"
)

// Hexagon setup
type Hexagon int

// InitPoints returns the starting points for the snails
func (h Hexagon) InitPoints() []point.Point2 {
	return []point.Point2{
		{X: 0, Y: 5 * math.Sqrt(3)},
		{X: 5, Y: 10 * math.Sqrt(3)},
		{X: 15, Y: 10 * math.Sqrt(3)},
		{X: 20, Y: 5 * math.Sqrt(3)},
		{X: 15, Y: 0},
		{X: 5, Y: 0},
	}
}

// Centroid returns the centroid of the starting points for the snails
func (h Hexagon) Centroid() point.Point2 {
	return point.Point2{X: 10, Y: 5 * math.Sqrt(3)}
}
