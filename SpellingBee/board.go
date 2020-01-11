package main

import "errors"

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
