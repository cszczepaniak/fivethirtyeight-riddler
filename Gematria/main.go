package main

import (
	"errors"
	"strings"
)

func main() {

}

func numberAsWord(n int) (string, error) {
	if n < 1 {
		return ``, errors.New(`number must be 1 or greater`)
	}
	str := firstTwoDigits(n)
	return str, nil
}

func firstTwoDigits(n int) string {
	n %= 100
	switch {
	case n < 10:
		return onesMap[n]
	case n < 20:
		return teensMap[n]
	}
	return tensMap[n/10] + onesMap[n%10]
}

func wordValue(w string) int {
	w = strings.ToLower(w)
	offset := 'a' - 1
	val := 0
	for _, r := range w {
		if r < 'a' || r > 'z' {
			continue
		}
		val += int(r - offset)
	}
	return val
}
