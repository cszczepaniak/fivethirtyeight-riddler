package letterset

type LetterSet map[rune]struct{}

func New(s string) LetterSet {
	ls := make(map[rune]struct{})
	for _, r := range s {
		if _, ok := ls[r]; !ok {
			ls[r] = struct{}{}
		}
	}
	return ls
}

func NumUniqueLetters(s string) int {
	return len(New(s))
}
