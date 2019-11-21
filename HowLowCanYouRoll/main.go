package main

import (
	"fmt"

	"github.com/cszczepaniak/fivethirtyeight-riddler/HowLowCanYouRoll/game"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(game.Play())
	}
}
