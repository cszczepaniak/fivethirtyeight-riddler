package main

import (
	"github.com/cszczepaniak/fivethirtyeight-riddler/HowLowCanYouRoll/game"
)

func main() {
	game.Play()
	return

	// defer printTime(time.Now())

	// nGames := 100000000
	// nWorkers := 100
	// sum := big.NewFloat(0)
	// results := make(chan *big.Float)
	// for i := 0; i < nWorkers; i++ {
	// 	go worker(nGames/nWorkers, results)
	// }
	// for i := 0; i < nGames; i++ {
	// 	sum.Add(sum, <-results)
	// }
	// avg := new(big.Float).Quo(sum, big.NewFloat(float64(nGames)))
	// fmt.Println(avg)
}

// func worker(nGames int, results chan<- *big.Float) {
// 	for i := 0; i < nGames; i++ {
// 		results <- game.Play()
// 	}
// }

// func printTime(start time.Time) {
// 	elapsed := time.Since(start)
// 	fmt.Printf("execution took %0.5f sec\n", elapsed.Seconds())
// }
