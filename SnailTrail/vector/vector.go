package vector

import (
	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/utils"
)

// Vector2 represents a 2d vector
type Vector2 struct {
	X float64
	Y float64
}

func (v *Vector2) normalize() {
	comps := []float64{v.X, v.Y}
	mag := utils.Norm2(comps)
	v.X /= mag
	v.Y /= mag
}

// WithLength stretches or shrinks the vector to the specified length
func (v *Vector2) WithLength(k float64) {
	v.normalize()
	v.X *= k
	v.Y *= k
}
