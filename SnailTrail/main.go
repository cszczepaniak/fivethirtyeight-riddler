package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/setup"
	"github.com/cszczepaniak/fivethirtyeight-riddler/SnailTrail/snail"
)

func main() {
	var stepSize float64
	flag.Float64Var(&stepSize, `stepsize`, 0.01, `The step size the snails will take`)
	flag.Parse()

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
	fmt.Printf("Total distance traveled: %0.2f meters\n", dist)
}
