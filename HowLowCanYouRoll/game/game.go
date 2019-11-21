package game

import (
	"math/big"
	"math/rand"
	"time"
)

func Play() *big.Float {
	var score string = `0.`
	rand.Seed(time.Now().UnixNano())
	for {
		roll := rand.Intn(10)
		updateScore(&score, roll)
		if lastDigit(score) == '0' {
			f, ok := new(big.Float).SetString(score)
			if ok {
				return f
			}
			return big.NewFloat(0)
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
