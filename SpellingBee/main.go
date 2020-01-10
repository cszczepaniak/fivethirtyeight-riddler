package main

import (
	"errors"
	"fmt"
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
	middle  rune
	letters letterSet
}

func newBoard(middle rune, others []rune) (board, error) {
	if len(others) != 6 {
		return board{}, errors.New(`a board must have six outer letters`)
	}
	letters := make(map[rune]struct{}, len(others)+1)
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
	letters[middle] = struct{}{}
	return board{
		middle:  middle,
		letters: letters,
	}, nil
}

func (b board) String() string {
	str := `[` + string(b.middle) + `: `
	for c := range b.letters {
		if c == b.middle {
			continue
		}
		str += string(c)
	}
	return str + `]`
}

func (b *board) canMakeWord(w word) bool {
	// the word must contain the middle letter
	if _, ok := w.letters[b.middle]; !ok {
		return false
	}
	for l := range w.letters {
		if _, ok := b.letters[l]; !ok {
			return false
		}
	}
	return true
}

func (b *board) scoreWord(w word) int {
	if !b.canMakeWord(w) {
		return 0
	}
	runes := []rune(w.str)
	if len(runes) == 4 {
		return 1
	}
	score := len(runes)
	isBonus := true
	for l := range b.letters {
		if _, ok := w.letters[l]; !ok {
			isBonus = false
			break
		}
	}
	if isBonus {
		score += 7
	}
	return score
}

type word struct {
	str     string
	letters letterSet
}

func newWord(w string) (word, error) {
	if len([]rune(w)) < 4 {
		return word{}, errors.New(`words must be at least 4 letters long`)
	}
	return word{
		str:     w,
		letters: newLetterSet(w),
	}, nil
}

func main() {
	if !fileExists(`words.txt`) {
		err := downloadWords()
		if err != nil {
			panic(err)
		}
	}

	b, err := newBoard('g', []rune{'a', 'p', 'x', 'm', 'e', 'l'})
	if err != nil {
		panic(err)
	}

	w, err := newWord(`amalgam`)
	if err != nil {
		panic(err)
	}
	score := b.scoreWord(w)
	fmt.Printf("score on board %s for word %s: %d\n", b, w.str, score)

	// words, err := filterWords()
	// if err != nil {
	// 	panic(err)
	// }
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
