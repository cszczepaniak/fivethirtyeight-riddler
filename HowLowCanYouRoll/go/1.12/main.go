package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/cszczepaniak/fivethirtyeight-riddler/HowLowCanYouRoll/game"
)

func main() {
	nGames := flag.Int(`n`, 1000000, `number of games to simulate`)
	flag.Parse()

	defer printTime(time.Now())
	nWorkers := 10
	var sum float64 = 0
	results := make(chan float64)
	for i := 0; i < nWorkers; i++ {
		go worker(*nGames/nWorkers, results)
	}
	for i := 0; i < *nGames; i++ {
		sum += <-results
	}
	avg := sum / float64(*nGames)
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
