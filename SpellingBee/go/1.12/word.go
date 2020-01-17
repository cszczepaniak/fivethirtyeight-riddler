package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/letterset"
)

type word struct {
	str     string
	letters letterset.LetterSet
}

func newWord(w string) (word, error) {
	if len([]rune(w)) < 4 {
		return word{}, errors.New(`words must be at least 4 letters long`)
	}
	return word{
		str:     w,
		letters: letterset.New(w),
	}, nil
}

func getWordList() ([]word, error) {
	ws, err := readWordsFromFile(`words.txt`)
	if err != nil {
		return nil, err
	}

	res := make([]word, 0, len(ws))
	for _, w := range ws {
		if len([]rune(w)) < 4 || strings.Contains(w, `s`) || letterset.NumUniqueLetters(w) > 7 {
			continue
		}
		word, err := newWord(w)
		if err != nil {
			return nil, err
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
