package main

import "strings"

func main() {

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
