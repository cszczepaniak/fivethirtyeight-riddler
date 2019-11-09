package main

import (
	"errors"
	"fmt"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/setup"
	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/snail"
)

func main() {
	const stepSize = 0.0001
	const maxSteps = 1000000
	var dist float64 = 0
	step := 0

	var s setup.Hexagon
	snails := snail.InitSnails(s)
	for snails[0].Pos.DistanceTo(s.Centroid()) > stepSize {
		for i := range snails {
			snails[i].Step(stepSize)
		}
		dist += stepSize
		step++
		if step > maxSteps {
			panic(errors.New(`too many steps taken`))
		}
	}
	fmt.Printf(`Total distance traveled: %0.2f meters\n`, dist)
}
