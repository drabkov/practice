package throwdice

import (
	"math/rand"
	"time"
)

func RandInt() int {
	return 1 + rand.Intn(6)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
