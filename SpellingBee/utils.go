package main

import "errors"

func combinations(superset []rune, n int) [][]rune {
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
		combs := combinations(others, n-1)
		for _, c := range combs {
			set := append(c, r)
			res = append(res, set)
		}
	}
	return res
}

func getAlphabetWithout(without []rune) ([]rune, error) {
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
