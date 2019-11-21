package game

import (
	"math/rand"
)

func Play() {
	var score string = `0.`
	for {
		roll := rand.Intn(10)
		if roll == 0 {
			break
		}
		updateScore(&score, roll)
	}
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
