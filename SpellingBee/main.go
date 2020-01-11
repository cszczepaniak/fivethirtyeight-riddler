package main

import (
	"fmt"
)

func main() {
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

	for _, r := range alphabet {
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
			if score > bestScore {
				bestScore = score
				bestBoard = b
			}
		}
	}
	fmt.Printf("best board: %s; with score %d\n", bestBoard, bestScore)
}
