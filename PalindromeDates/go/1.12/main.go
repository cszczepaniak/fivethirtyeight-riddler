package main

import (
	"fmt"
	"time"
)

func main() {
	n := 0
	t := time.Date(2020, time.February, 3, 0, 0, 0, 0, time.Local)
	for t.Year() < 2100 {
		if isPalindrome(t.Format(`01022006`)) {
			n++
			fmt.Printf("Date %s is a palindrome: %s\n", t.Format(`Jan 2 2006`), t.Format(`01022006`))
		}
		t = t.AddDate(0, 0, 1)
	}
	fmt.Printf("There are %d palindrome dates left this century (after Feb 02 2020)\n", n)
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
