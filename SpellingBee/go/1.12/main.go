package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/board"
	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/letterset"
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

	pointsMap := utils.CalculatePointsMap(words)

	boards, err := utils.GetPossibleBoards(words)
	if err != nil {
		log.Fatal(err)
	}

	var bestBoard board.Board
	bestScore := 0

	resChan := make(chan result)
	var wg sync.WaitGroup
	gps := divideBoards(boards, runtime.NumCPU())
	for _, g := range gps {
		go worker(resChan, g, words, pointsMap, &wg)
		wg.Add(1)
	}
	totalBds := len(boards)
	sizeOfOneTenth := totalBds / 10
	doneSoFar := 0
	go func() {
		for {
			r := <-resChan
			if r.score > bestScore {
				bestScore = r.score
				bestBoard = r.b
			}
			doneSoFar++
			if doneSoFar%sizeOfOneTenth == 0 {
				fmt.Printf("%d of %d boards done!\n", doneSoFar, totalBds)
			}
		}
	}()
	wg.Wait()
	fmt.Printf("best board: %s; with score %d\n", bestBoard, bestScore)
	d := time.Since(start)
	fmt.Printf("Elapsed time: %s \n", d)
}

func divideBoards(bs []board.Board, n int) [][]board.Board {
	gpSize := len(bs) / n
	leftover := len(bs) % n
	res := make([][]board.Board, n)
	for i := 0; i < n; i++ {
		num := 0
		if leftover > 0 {
			num = gpSize + 1
			leftover--
		} else {
			num = gpSize
		}
		res[i] = bs[0:num]
		bs = bs[num:]
	}
	return res
}

func worker(resChan chan result, boards []board.Board, words []word.Word, pointsMap map[letterset.LetterSet]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, b := range boards {
		score := 0
		subsets := utils.BoardSubsets(b)
		for _, ls := range subsets {
			if s, ok := pointsMap[ls]; ok {
				score += s
			}
		}
		resChan <- result{
			b:     b,
			score: score,
		}
	}
}
