package game

import (
	"math/rand"
	"time"
)

func Play() float64 {
	var score uint64 = 0
	rand.Seed(time.Now().UnixNano())
	for roll := rand.Intn(10); roll != 0; roll = rand.Intn(10) {
		updateScore(&score, roll)
	}
	return scoreToFloat(score)
}

func lastDigit(n uint64) int {
	return int(n % uint64(10))
}

func updateScore(score *uint64, roll int) {
	if roll <= lastDigit(*score) || *score == 0 {
		*score = *score*10 + uint64(roll)
	}
}

func scoreToFloat(score uint64) float64 {
	denom := 1
	quo := score
	for quo > 0 {
		quo /= 10
		denom *= 10
	}
	return float64(score) / float64(denom)
}
