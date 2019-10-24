package throwdice

import (
	"fmt"
	"math/rand"
	"time"
)

func ThrowDice(numberShots int) {

	m := make(map[int]int)

	for i := 0; i < numberShots; i++ {
		x := randInt()
		y := randInt()
		m[x+y] += 1
	}

	for i, v := range m {
		fmt.Printf("%d - %2.1f %%\n", i, float64(v)/float64(numberShots)*100)
	}
}

func randInt() int {
	return 1 + rand.Intn(6)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
