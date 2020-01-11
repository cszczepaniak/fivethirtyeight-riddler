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

	words, err := getWordList()
	if err != nil {
		panic(err)
	}
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
		res = append(res, word{
			str:     w,
			letters: newLetterSet(w),
		})
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
