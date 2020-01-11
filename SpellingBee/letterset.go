package main

type letterSet map[rune]struct{}

func newLetterSet(s string) letterSet {
	ls := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := ls[r]; !ok {
			ls[r] = struct{}{}
		}
	}
	return ls
}
