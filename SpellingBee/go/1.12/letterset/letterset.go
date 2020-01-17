package letterset

import "sort"

type LetterSet string

func New(s string) LetterSet {
	ls := make(map[rune]struct{})
	rs := make([]rune, 0)
	for _, r := range s {
		if _, ok := ls[r]; !ok {
			ls[r] = struct{}{}
			rs = append(rs, r)
		}
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i] < rs[j]
	})
	return LetterSet(rs)
}

func Length(s string) int {
	return len(New(s))
}
