package main

import (
	"time"
)

func main() {
	n := 0
	t := time.Date(2020, time.February, 3, 0, 0, 0, 0, time.Local)
	for t.Year() < 2100 {
		if isPalindrome(t.Format(`01022006`)) {
			n++
		}
		t.AddDate(0, 0, 1)
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
