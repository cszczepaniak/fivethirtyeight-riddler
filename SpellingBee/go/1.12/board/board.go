package board

import (
	"errors"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/letterset"
	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/word"
)

type Board struct {
	Middle  rune
	Letters letterset.LetterSet
}

func New(middle rune, others []rune) (Board, error) {
	if len(others) != 6 {
		return Board{}, errors.New(`a board must have six outer letters`)
	}
	letters := make(map[rune]struct{}, len(others)+1)
	for _, r := range others {
		if _, ok := letters[r]; !ok {
			letters[r] = struct{}{}
		} else {
			return Board{}, errors.New(`a board cannot contain duplicate letters`)
		}
	}
	if _, ok := letters[middle]; ok {
		return Board{}, errors.New(`a board cannot contain duplicate letters`)
	}
	letters[middle] = struct{}{}
	return Board{
		Middle:  middle,
		Letters: letters,
	}, nil
}

func (b Board) String() string {
	str := `[` + string(b.Middle) + `: `
	for c := range b.Letters {
		if c == b.Middle {
			continue
		}
		str += string(c)
	}
	return str + `]`
}

func (b *Board) CanMakeWord(w word.Word) bool {
	// the word must contain the middle letter
	if _, ok := w.Letters[b.Middle]; !ok {
		return false
	}
	for l := range w.Letters {
		if _, ok := b.Letters[l]; !ok {
			return false
		}
	}
	return true
}

func (b *Board) ScoreWord(w word.Word) int {
	if !b.CanMakeWord(w) {
		return 0
	}
	runes := []rune(w.Str)
	if len(runes) == 4 {
		return 1
	}
	score := len(runes)
	isBonus := true
	for l := range b.Letters {
		if _, ok := w.Letters[l]; !ok {
			isBonus = false
			break
		}
	}
	if isBonus {
		score += 7
	}
	return score
}
