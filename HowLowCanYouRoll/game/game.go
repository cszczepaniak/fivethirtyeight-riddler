package game

import (
	"fmt"
	"math/rand"
	"time"
)

func Play() float64 {
	var score uint64 = 0
	rand.Seed(time.Now().UnixNano())
	for roll := rand.Intn(10); roll != 0; roll = rand.Intn(10) {
		updateScore(&score, roll)
	}
	fmt.Println(score)
	return 0
}

func lastDigit(n uint64) int {
	return int(n % uint64(10))
}

func updateScore(score *uint64, roll int) {
	if roll <= lastDigit(*score) || *score == 0 {
		*score = *score*10 + uint64(roll)
	}
}
