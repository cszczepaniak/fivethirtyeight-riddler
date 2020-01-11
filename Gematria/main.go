package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	bestNum, bestNumVal, max := 0, 0, 10000
	bestStr := ``
	for i := 1; i < max; i++ {
		w, err := numberAsWord(i)
		if err != nil {
			panic(err)
		}
		val := wordValue(w)
		if i < val {
			bestNum = i
			bestNumVal = val
			bestStr = w
		}
	}
	fmt.Printf("best number under %d: %d with an alphanumeric value of %d\n", max, bestNum, bestNumVal)
	fmt.Printf("and a string of %q\n", bestStr)
}

func joinWords(strs ...string) string {
	str := ``
	for _, s := range strs {
		str += strings.TrimSpace(s)
		str += ` `
	}
	return strings.TrimSpace(str)
}

func numberAsWord(n int) (string, error) {
	if n < 1 {
		return ``, errors.New(`number must be 1 or greater`)
	}
	str := ``
	i := 0
	for n > 0 {
		thisGroup := lowestThreeDigits(n)
		str = joinWords(thisGroup, groups[i], str)
		n /= 1000
		i++
	}
	return str, nil
}

func lowestThreeDigits(n int) string {
	tens := n % 100
	switch {
	case tens < 10:
		return onesMap[tens]
	case tens < 20:
		return teensMap[tens]
	}
	tensStr := joinWords(tensMap[tens/10], onesMap[tens%10])

	hundreds := (n % 1000) / 100
	if hundreds > 0 {
		return joinWords(onesMap[hundreds], `hundred`, tensStr)
	}
	return tensStr
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
