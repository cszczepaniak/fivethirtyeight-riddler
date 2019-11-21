package game

import (
	"math/rand"
)

func Play() string {
	var score string = `0.`
	for {
		roll := rand.Intn(10)
		updateScore(&score, roll)
		if lastDigit(score) == '0' {
			return score
		}
	}
}

func lastDigit(s string) rune {
	runes := []rune(s)
	return runes[len(runes)-1]
}

func updateScore(score *string, roll int) {
	scoreRunes := []rune(*score)
	lastDigit := scoreRunes[len(scoreRunes)-1]
	rollRune := rune(roll + 48)
	if lastDigit == '.' {
		scoreRunes = append(scoreRunes, rollRune)
	} else if rollRune <= lastDigit {
		scoreRunes = append(scoreRunes, rollRune)
	}
	*score = string(scoreRunes)
}
