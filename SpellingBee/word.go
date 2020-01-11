package main

import "errors"

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
