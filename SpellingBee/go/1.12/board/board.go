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
	all := append(others, middle)
	return Board{
		Middle:  middle,
		Letters: letterset.FromRunes(all),
	}, nil
}

func (b Board) String() string {
	str := `[` + string(b.Middle) + `: `
	for _, c := range b.Letters {
		if c == b.Middle {
			continue
		}
		str += string(c)
	}
	return str + `]`
}

func (b *Board) CanMakeWord(w word.Word) bool {
	// the word must contain the middle letter
	if !w.Letters.Contains(b.Middle) {
		return false
	}
	for _, l := range w.Letters {
		if !b.Letters.Contains(l) {
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
	for _, l := range b.Letters {
		if !w.Letters.Contains(l) {
			isBonus = false
			break
		}
	}
	if isBonus {
		score += 7
	}
	return score
}
