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

	if !fileExists(`filtered.txt`) {
		err := filterWords()
		if err != nil {
			panic(err)
		}
	}
}

func filterWords() error {
	in, err := os.Open(`words.txt`)
	if err != nil {
		return err
	}
	defer in.Close()

	bytes, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}

	txt := string(bytes)
	words := strings.Split(txt, "\n")

	out, err := os.Create(`filtered.txt`)
	if err != nil {
		return err
	}
	defer out.Close()

	for _, w := range words {
		if len([]rune(w)) < 4 || strings.Contains(w, `s`) {
			continue
		}
		_, err := fmt.Fprintf(out, "%s\n", w)
		if err != nil {
			return err
		}
	}
	return nil
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
