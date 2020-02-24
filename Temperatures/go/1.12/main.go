package main

import (
	"fmt"
	"math"
)

/*
Note: you clearly don't need to write a program to solve this, but I did anyway...
*/

const (
	start = -99
	end   = 100
)

func main() {
	for c := start; c < end; c++ {
		f := celciusToFahrenheit(c)
		if f < -99 || f > 99 {
			continue
		}
		if areMirrors(c, f) {
			fmt.Printf("%dC = %dF\n", c, f)
		}
	}
}

func celciusToFahrenheit(c int) int {
	cFlt := float64(c)
	return int(math.Round(cFlt*9/5 + 32))
}

func areMirrors(n1, n2 int) bool {
	return (n1%10 == n2/10) && (n2%10 == n1/10)
}
