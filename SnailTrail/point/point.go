package point

import (
	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/utils"
)

// Point represents a cartesian point
type Point struct {
	X float64
	Y float64
}

// DistanceTo calculates the distance between this point and another
func (p Point) DistanceTo(pt Point) float64 {
	coords := make([]float64, 2)
	coords[0] = p.X - pt.X
	coords[1] = p.Y - pt.Y
	return utils.Norm2(coords)
}
