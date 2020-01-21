package dice

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func roll() int {
	return rand.Intn(6) + 1
}

func RollN(n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = roll()
	}
	sort.Ints(res)
	return res
}
