package game

import (
	"math/big"
	"math/rand"
)

func Play() {
	n := big.NewInt(0)
	for {
		roll := rand.Int63n(10)
		if roll == 0 {
			break
		}
		n = updateScore(n, roll)
	}
}

func updateScore(score *big.Int, roll int64) *big.Int {
	r := big.NewInt(roll)
	rem := new(big.Int).Mod(score, big.NewInt(int64(10)))
	if r.Cmp(rem) > 0 && score.Cmp(big.NewInt(0)) != 0 {
		return score
	}
	score.Add(new(big.Int).Mul(score, big.NewInt(10)), r)
	return score
}
