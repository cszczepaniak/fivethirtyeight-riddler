package main

import (
	"fmt"

	"github.com/cszczepaniak/fivethirtyeight-riddler/DeliriousDucks/game"
)

const (
	width  = 3
	height = 3
)

func main() {
	g := game.New(game.Config{
		BoardWidth:  3,
		BoardHeight: 3,
		NumDucks:    2,
		StartX:      1,
		StartY:      1,
	})
	i := g.Play()
	fmt.Printf("ducks met after %d minutes\n", i)
}
