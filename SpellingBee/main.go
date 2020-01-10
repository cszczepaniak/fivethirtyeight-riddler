package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type letterSet map[rune]struct{}

func newLetterSet(s string) letterSet {
	runes := []rune(s)
	ls := make(map[rune]struct{})
	for _, r := range runes {
		if _, ok := ls[r]; !ok {
			ls[r] = struct{}{}
		}
	}
	return ls
}

type board struct {
	middle rune
	others letterSet
}

func newBoard(middle rune, others []rune) (board, error) {
	if len(others) != 6 {
		return board{}, errors.New(`a board must have six outer letters`)
	}
	letters := make(map[rune]struct{}, len(others))
	for _, r := range others {
		if _, ok := letters[r]; !ok {
			letters[r] = struct{}{}
		} else {
			return board{}, errors.New(`a board cannot contain duplicate letters`)
		}
	}
	if _, ok := letters[middle]; ok {
		return board{}, errors.New(`a board cannot contain duplicate letters`)
	}
	return board{
		middle: middle,
		others: letters,
	}, nil
}

type word struct {
	str     string
	letters letterSet
}

func newWord(w string) word {
	return word{
		str:     w,
		letters: newLetterSet(w),
	}
}

func main() {
	if !fileExists(`words.txt`) {
		err := downloadWords()
		if err != nil {
			panic(err)
		}
	}

	words, err := filterWords()
	if err != nil {
		panic(err)
	}
}

func buildLetterSet(word string) letterSet {
	runes := []rune(word)
	letterSet := make(map[rune]struct{})
	for _, r := range runes {
		if _, ok := letterSet[r]; !ok {
			letterSet[r] = struct{}{}
		}
	}
	return letterSet
}

func filterWords() ([]string, error) {
	words, err := readWordsFromFile(`words.txt`)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0, len(words))
	for _, w := range words {
		if len([]rune(w)) < 4 || strings.Contains(w, `s`) {
			continue
		}
		res = append(res, w)
		if err != nil {
			return nil, err
		}
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
