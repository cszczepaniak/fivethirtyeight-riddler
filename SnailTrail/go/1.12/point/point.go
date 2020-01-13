package point

import (
	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/utils"
	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/vector"
)

// Point2 represents a cartesian point
type Point2 struct {
	X float64
	Y float64
}

// DistanceTo calculates the distance between this point and another
func (p1 Point2) DistanceTo(p2 Point2) float64 {
	coords := make([]float64, 2)
	coords[0] = p1.X - p2.X
	coords[1] = p1.Y - p2.Y
	return utils.Norm2(coords)
}

// VectorFrom creates a vector from p2 to p1
func (p1 Point2) VectorFrom(p2 Point2) vector.Vector2 {
	return vector.Vector2{X: p1.X - p2.X, Y: p1.Y - p2.Y}
}
