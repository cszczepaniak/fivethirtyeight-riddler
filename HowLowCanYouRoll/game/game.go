package game

import "math/rand"

func Play() {
	var n int64 = 0
	for {
		roll := rand.Int63n(10)
		if roll == 0 {
			break
		}
		n = updateScore(n, roll)
	}
}

func updateScore(score, roll int64) int64 {
	if roll > score%10 && score != 0 {
		return score
	}
	return score*10 + roll
}
