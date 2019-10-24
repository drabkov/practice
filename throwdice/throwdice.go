package throwdice

import (
	"math/rand"
	"time"
)

func ThrowDice(numberShots int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < numberShots; i++ {
		x := randInt()
		y := randInt()
		m[x+y] += 1
	}
	return m
}

func randInt() int {
	return 1 + rand.Intn(6)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
