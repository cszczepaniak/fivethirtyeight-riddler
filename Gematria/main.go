package main

import (
	"errors"
	"strings"
)

func main() {

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
		str = joinWords(str, thisGroup, groups[i])
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
