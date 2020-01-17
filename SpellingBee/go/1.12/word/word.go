package word

import (
	"errors"
	"strings"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/letterset"
)

type Word struct {
	Str     string
	Letters letterset.LetterSet
}

func New(w string) (Word, error) {
	if len([]rune(w)) < 4 {
		return Word{}, errors.New(`words must be at least 4 letters long`)
	}
	return Word{
		Str:     w,
		Letters: letterset.New(w),
	}, nil
}

func (w Word) IsPangram() bool {
	return len(w.Letters) == 7
}

// FilterWords removes words which won't be used according to the rules and convers them into Word structs
func FilterWords(words []string) ([]Word, error) {
	res := make([]Word, 0, len(words))
	for _, w := range words {
		if len([]rune(w)) < 4 || strings.Contains(w, `s`) || letterset.NumUniqueLetters(w) > 7 {
			continue
		}
		word, err := New(w)
		if err != nil {
			return nil, err
		}
		res = append(res, word)
	}
	return res, nil
}
