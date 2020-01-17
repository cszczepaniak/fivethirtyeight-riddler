package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/board"
)

type result struct {
	b     board.Board
	score int
}

func main() {
	start := time.Now()
	if !fileExists(`words.txt`) {
		err := downloadWords()
		if err != nil {
			panic(err)
		}
	}

	words, err := getWordList()
	if err != nil {
		panic(err)
	}

	var bestBoard board
	bestScore := 0

	alphabet, err := getAlphabetWithout([]rune{'s'})
	if err != nil {
		panic(err)
	}

	resChan := make(chan result)
	var wg sync.WaitGroup
	gps := divideAlphabet(alphabet, runtime.NumCPU())
	for _, g := range gps {
		go worker(resChan, g, words, &wg)
		wg.Add(1)
	}
	go func() {
		for {
			r := <-resChan
			if r.score > bestScore {
				bestScore = r.score
				bestBoard = r.b
			}
		}
	}()
	wg.Wait()
	fmt.Printf("best board: %s; with score %d\n", bestBoard, bestScore)
	d := time.Since(start)
	fmt.Printf("Elapsed time %s: \n", d)
}

func divideAlphabet(a []rune, n int) [][]rune {
	groupSize := len(a) / n
	leftover := len(a) % n
	res := make([][]rune, n)
	for i := 0; i < n; i++ {
		n := 0
		if leftover > 0 {
			n = groupSize + 1
			leftover--
		} else {
			n = groupSize
		}
		res[i] = a[0:n]
		a = a[n:]
	}
	return res
}

func worker(resChan chan result, letters []rune, words []word, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, r := range letters {
		fmt.Printf("NOW SERVING: %c\n", r)
		boards, err := allBoardsWithCenter(r)
		if err != nil {
			panic(err)
		}
		for _, b := range boards {
			score := 0
			for _, w := range words {
				if b.canMakeWord(w) {
					score += b.scoreWord(w)
				}
			}
			resChan <- result{
				b:     b,
				score: score,
			}
		}
	}
}
