package main

import (
	"fmt"
	"time"

	"github.com/cszczepaniak/fivethirtyeight-riddler/HowLowCanYouRoll/game"
)

func main() {
	defer printTime(time.Now())

	nGames := 100000000
	nWorkers := 100
	var sum float64 = 0
	results := make(chan float64)
	for i := 0; i < nWorkers; i++ {
		go worker(nGames/nWorkers, results)
	}
	for i := 0; i < nGames; i++ {
		sum += <-results
	}
	avg := sum / float64(nGames) //new(big.Float).Quo(sum, big.NewFloat(float64(nGames)))
	fmt.Println(avg)
}

func worker(nGames int, results chan<- float64) {
	for i := 0; i < nGames; i++ {
		results <- game.Play()
	}
}

func printTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("execution took %0.5f sec\n", elapsed.Seconds())
}
