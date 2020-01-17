package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/board"
	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/letterset"
	"github.com/cszczepaniak/fivethirtyeight-riddler/SpellingBee/word"
)

func Combinations(superset []rune, n int) [][]rune {
	if len(superset) == n {
		return [][]rune{superset}
	}
	if n == 1 {
		res := make([][]rune, len(superset))
		for i, r := range superset {
			res[i] = []rune{r}
		}
		return res
	}
	res := make([][]rune, 0)
	for i, r := range superset {
		if i > len(superset)-n {
			break
		}
		others := superset[i+1:]
		combs := Combinations(others, n-1)
		for _, c := range combs {
			set := append(c, r)
			res = append(res, set)
		}
	}
	return res
}

func GetAlphabetWithout(without []rune) ([]rune, error) {
	runeMap := make(map[rune]struct{}, len(without))
	for _, r := range without {
		if r < 'a' || r > 'z' {
			return nil, errors.New(`must supply a rune to omit between 'a' and 'z'`)
		}
		if _, ok := runeMap[r]; ok {
			return nil, errors.New(`duplicate runes in input`)
		}
		runeMap[r] = struct{}{}
	}
	// we will always omit 's' and without, so capacity is 24
	res := make([]rune, 0, 24)
	offset := 'a'
	for i := 0; i < 26; i++ {
		r := rune(i) + offset
		if _, ok := runeMap[r]; ok {
			continue
		}
		res = append(res, r)
	}
	return res, nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DownloadWords() error {
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

func ReadWordsFromFile(file string) ([]string, error) {
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

func BoardsFromLetterset(ls letterset.LetterSet) ([]board.Board, error) {
	if len(ls) != 7 {
		return nil, errors.New(`letterset must represent a pangram`)
	}
	res := make([]board.Board, 0, 7)
	for _, r := range ls {
		res = append(res, board.Board{
			Middle:  r,
			Letters: ls,
		})
	}
	return res, nil
}

func GetPossibleBoards(words []word.Word) ([]board.Board, error) {
	bds := make([]board.Board, 0)
	for _, w := range words {
		if w.IsPangram() {
			b, err := BoardsFromLetterset(w.Letters)
			if err != nil {
				return nil, err
			}
			bds = append(bds, b...)
		}
	}
	return bds, nil
}

func AllBoardsWithCenter(middle rune) ([]board.Board, error) {
	a, err := GetAlphabetWithout([]rune{middle, 's'})
	if err != nil {
		return nil, err
	}
	outerCombs := Combinations(a, 6)
	res := make([]board.Board, len(outerCombs))
	for i, c := range outerCombs {
		b, err := board.New(middle, c)
		if err != nil {
			return nil, err
		}
		res[i] = b
	}
	return res, nil
}

func CalculatePointsMap(words []word.Word) map[letterset.LetterSet]int {
	pointsMap := make(map[letterset.LetterSet]int)
	for _, w := range words {
		if _, ok := pointsMap[w.Letters]; !ok {
			pointsMap[w.Letters] = w.Score()
		} else {
			pointsMap[w.Letters] += w.Score()
		}
	}
	return pointsMap
}

func BoardSubsets(b board.Board) []letterset.LetterSet {
	outer := make([]rune, 0, len(b.Letters)-1)
	for _, r := range b.Letters {
		if r == b.Middle {
			continue
		}
		outer = append(outer, r)
	}
	subsets := make([]letterset.LetterSet, 0)
	for i := 1; i < 8; i++ {
		cs := Combinations(outer, i)
		for _, c := range cs {
			subsets = append(subsets, letterset.FromRunes(c))
		}
	}
	return append(subsets, letterset.FromRunes([]rune{b.Middle}))
}
