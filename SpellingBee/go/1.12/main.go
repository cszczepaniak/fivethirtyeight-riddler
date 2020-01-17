package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/board"
	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/utils"
	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/word"
)

type result struct {
	b     board.Board
	score int
}

func main() {
	start := time.Now()
	if !utils.FileExists(`words.txt`) {
		err := utils.DownloadWords()
		if err != nil {
			log.Fatal(err)
		}
	}

	wordStrs, err := utils.ReadWordsFromFile(`words.txt`)
	if err != nil {
		log.Fatal(err)
	}

	words, err := word.FilterWords(wordStrs)
	if err != nil {
		log.Fatal(err)
	}

	boards, err := utils.GetPossibleBoards(words)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of boards: %d\n", len(boards))
	return

	var bestBoard board.Board
	bestScore := 0

	alphabet, err := utils.GetAlphabetWithout([]rune{'s'})
	if err != nil {
		log.Fatal(err)
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

func worker(resChan chan result, letters []rune, words []word.Word, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, r := range letters {
		fmt.Printf("NOW SERVING: %c\n", r)
		boards, err := utils.AllBoardsWithCenter(r)
		if err != nil {
			log.Fatal(err)
		}
		for _, b := range boards {
			score := 0
			for _, w := range words {
				if b.CanMakeWord(w) {
					score += b.ScoreWord(w)
				}
			}
			resChan <- result{
				b:     b,
				score: score,
			}
		}
	}
}
