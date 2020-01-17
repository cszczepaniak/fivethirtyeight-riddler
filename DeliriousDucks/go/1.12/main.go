package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/cszczepaniak/fivethirtyeight-riddler/DeliriousDucks/game"
)

const (
	width      = 3
	height     = 3
	totalGames = 10000000
)

func main() {
	var totalMinutes uint64 = 0
	resChan := make(chan int)
	var wg sync.WaitGroup
	gps := splitEvenly(totalGames, runtime.NumCPU())
	for _, n := range gps {
		wg.Add(1)
		go worker(n, resChan, &wg)
	}
	go func() {
		for {
			totalMinutes += uint64(<-resChan)
		}
	}()
	wg.Wait()
	fmt.Printf("ducks met after and average of %0.3f minutes\n", float64(totalMinutes)/float64(totalGames))
}

func splitEvenly(n, k int) []int {
	gps := make([]int, k)
	// split n into k groups
	nomSize := n / k
	leftover := n % k
	for i := range gps {
		if leftover > 0 {
			gps[i] = nomSize + 1
			leftover--
		} else {
			gps[i] = nomSize
		}
	}
	return gps
}

func worker(nGames int, rc chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < nGames; i++ {
		g := game.New(game.Config{
			BoardWidth:  3,
			BoardHeight: 3,
			NumDucks:    2,
			StartX:      1,
			StartY:      1,
		})
		rc <- g.Play()
	}
}
