package setup

import "github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/point"

// Setup is an interface to determine the setup of the problem
type Setup interface {
	InitPoints() []point.Point2
	Centroid() point.Point2
}
