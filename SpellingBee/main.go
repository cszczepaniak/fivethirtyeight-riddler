package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func getWordList() ([]word, error) {
	ws, err := readWordsFromFile(`words.txt`)
	if err != nil {
		return nil, err
	}

	res := make([]word, 0, len(ws))
	for _, w := range ws {
		if len([]rune(w)) < 4 || strings.Contains(w, `s`) {
			continue
		}
		word, err := newWord(w)
		if err != nil {
			panic(err)
		}
		res = append(res, word)
	}
	return res, nil
}

func readWordsFromFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	txt := string(bytes)
	return strings.Split(txt, "\n"), nil
}

func downloadWords() error {
	resp, err := http.Get(`https://norvig.com/ngrams/enable1.txt`)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	f, err := os.Create(`words.txt`)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
